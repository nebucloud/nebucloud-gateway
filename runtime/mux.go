package runtime

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/wundergraph/graphql-go-tools/pkg/ast"
	"github.com/wundergraph/graphql-go-tools/pkg/astnormalization"
	"github.com/wundergraph/graphql-go-tools/pkg/astparser"
	"github.com/wundergraph/graphql-go-tools/pkg/asttransform"
	"github.com/wundergraph/graphql-go-tools/pkg/operationreport"
	"google.golang.org/grpc"
)

type MiddlewareFunc func(ctx context.Context, w http.ResponseWriter, r *http.Request) (context.Context, error)

type GraphqlHandler interface {
	CreateConnection(context.Context) (*grpc.ClientConn, func(), error)
	GetMutations(*grpc.ClientConn) map[string]*ast.FieldDefinition
	GetQueries(*grpc.ClientConn) map[string]*ast.FieldDefinition
}

type ServeMux struct {
	middlewares  []MiddlewareFunc
	Schema       *ast.Document
	ErrorHandler GraphqlErrorHandler

	handlers []GraphqlHandler
}

func NewServeMux(ms ...MiddlewareFunc) *ServeMux {
	return &ServeMux{
		middlewares: ms,
		handlers:    make([]GraphqlHandler, 0),
	}
}

func (s *ServeMux) AddHandler(h GraphqlHandler) error {
	if err := s.validateHandler(h); err != nil {
		return err
	}
	s.handlers = append(s.handlers, h)
	return nil
}

func buildSchema(queries, mutations map[string]*ast.FieldDefinition) (*ast.Document, error) {
	document := ast.NewDocument()

	if len(queries) > 0 {
		queryType := ast.ObjectTypeDefinition{
			Description: ast.Description{
				IsDefined:     true,
				IsBlockString: false,
				Content:       document.Input.AppendInputBytes([]byte("The query root of the schema.")),
			},
			Name: document.Input.AppendInputString("Query"),
		}
		queryTypeIndex := document.AddObjectTypeDefinition(queryType)
		for name, field := range queries {
			fieldDef := document.AddFieldDefinition(ast.FieldDefinition{
				Name: document.Input.AppendInputString(name),
				Type: field.Type,
			})
			document.ObjectTypeDefinitions[queryTypeIndex].FieldsDefinition.Refs = append(document.ObjectTypeDefinitions[queryTypeIndex].FieldsDefinition.Refs, fieldDef)
		}
	}

	if len(mutations) > 0 {
		mutationType := ast.ObjectTypeDefinition{
			Description: ast.Description{
				IsDefined:     true,
				IsBlockString: false,
				Content:       document.Input.AppendInputBytes([]byte("The mutation root of the schema.")),
			},
			Name: document.Input.AppendInputString("Mutation"),
		}
		mutationTypeIndex := document.AddObjectTypeDefinition(mutationType)
		for name, field := range mutations {
			fieldDef := document.AddFieldDefinition(ast.FieldDefinition{
				Name: document.Input.AppendInputString(name),
				Type: field.Type,
			})
			document.ObjectTypeDefinitions[mutationTypeIndex].FieldsDefinition.Refs = append(document.ObjectTypeDefinitions[mutationTypeIndex].FieldsDefinition.Refs, fieldDef)
		}
	}

	return document, nil
}

func (s *ServeMux) validateHandler(h GraphqlHandler) error {
	queries := h.GetQueries(nil)
	mutations := h.GetMutations(nil)

	if len(queries) == 0 && len(mutations) == 0 {
		return nil
	}

	document, err := buildSchema(queries, mutations)
	if err != nil {
		return fmt.Errorf("schema validation error: %s", err)
	}

	report := &operationreport.Report{}
	asttransform.MergeDefinitionWithBaseSchema(document)
	if report.HasErrors() {
		return fmt.Errorf("schema validation error: %s", report.Error())
	}

	return nil
}

func (s *ServeMux) Use(ms ...MiddlewareFunc) *ServeMux {
	s.middlewares = append(s.middlewares, ms...)
	return s
}

func (s *ServeMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var params struct {
		Query         string                 `json:"query"`
		OperationName string                 `json:"operationName"`
		Variables     map[string]interface{} `json:"variables"`
	}

	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Parse the query
	operation, parseErrors := parseOperation(s.Schema)
	if len(parseErrors) > 0 {
		respondWithErrors(w, parseErrors)
		return
	}

	// Execute the query
	result, errors := ExecuteGraphQL(r.Context(), s.Schema, operation, params.Variables)

	if len(errors) > 0 {
		if s.ErrorHandler != nil {
			s.ErrorHandler(errors)
		} else {
			defaultGraphqlErrorHandler(errors)
		}
		respondWithErrors(w, errors)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.Write(result)
	}
}

func respondWithErrors(w http.ResponseWriter, errors []GraphqlError) {
	response := struct {
		Errors []GraphqlError `json:"errors"`
	}{
		Errors: errors,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func parseOperation(schema *ast.Document) (*ast.OperationDefinition, []GraphqlError) {
	report := &operationreport.Report{}
	queryDocument := ast.NewDocument()
	queryParser := astparser.NewParser()
	queryParser.Parse(queryDocument, report)

	if report.HasErrors() {
		return nil, ConvertToGraphQLErrors(report.ExternalErrors)
	}

	normalizer := astnormalization.NewNormalizer(false, false)
	normalizer.NormalizeOperation(queryDocument, schema, report)

	if report.HasErrors() {
		return nil, ConvertToGraphQLErrors(report.ExternalErrors)
	}

	if len(queryDocument.OperationDefinitions) == 0 {
		return nil, []GraphqlError{{Message: "No operation found in the query"}}
	}

	// Get the first operation definition
	operation := &queryDocument.OperationDefinitions[0]

	return operation, nil
}
