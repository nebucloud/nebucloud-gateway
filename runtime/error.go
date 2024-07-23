package runtime

import (
	"context"
	"encoding/json"
	"regexp"
	"strings"

	"github.com/wundergraph/graphql-go-tools/pkg/ast"
	"github.com/wundergraph/graphql-go-tools/pkg/execution"
	"github.com/wundergraph/graphql-go-tools/pkg/graphqlerrors"
	"github.com/wundergraph/graphql-go-tools/pkg/operationreport"
)

var (
	// gRPC common status error describes following error message.
	// Find submatch of status text and pluck it.
	// See: https://github.com/grpc/grpc-go/blob/master/internal/status/status.go#L146
	grpcBackendErrorMatcher = regexp.MustCompile(`rpc error: code = ([^\s]+).*desc = (.+)`)
)

// GraphqlError represents a GraphQL error
type GraphqlError struct {
	Message    string                   `json:"message"`
	Locations  []graphqlerrors.Location `json:"locations,omitempty"`
	Path       []interface{}            `json:"path,omitempty"`
	Extensions map[string]interface{}   `json:"extensions,omitempty"`
}

// GraphqlErrorHandler is a function type for custom error handling
type GraphqlErrorHandler func(errs []GraphqlError)

// defaultGraphqlErrorHandler adds error extension code from gRPC error message
func defaultGraphqlErrorHandler(errs []GraphqlError) {
	for i := range errs {
		m := grpcBackendErrorMatcher.FindStringSubmatch(errs[i].Message)
		if m == nil {
			continue
		}
		errs[i].Message = strings.TrimSpace(m[2])
		if errs[i].Extensions == nil {
			errs[i].Extensions = make(map[string]interface{})
		}
		errs[i].Extensions["code"] = strings.ToUpper(m[1])
	}
}

// ConvertToGraphQLErrors converts operationreport.ExternalErrors to GraphqlErrors
func ConvertToGraphQLErrors(externalErrors []operationreport.ExternalError) []GraphqlError {
	var graphqlErrors []GraphqlError
	for _, err := range externalErrors {
		graphqlError := GraphqlError{
			Message:   err.Message,
			Locations: err.Locations,
			Path:      convertPath(err.Path),
		}
		graphqlErrors = append(graphqlErrors, graphqlError)
	}
	return graphqlErrors
}

// convertPath converts ast.PathItem to []interface{}
func convertPath(path []ast.PathItem) []interface{} {
	var result []interface{}
	for _, item := range path {
		switch item.Kind {
		case ast.ArrayIndex:
			result = append(result, item.ArrayIndex)
		case ast.FieldName:
			result = append(result, string(item.FieldName))
		}
	}
	return result
}

// Custom MarshalJSON for GraphqlError to handle Path marshaling
func (e *GraphqlError) MarshalJSON() ([]byte, error) {
	type Alias GraphqlError
	return json.Marshal(&struct {
		Path json.RawMessage `json:"path,omitempty"`
		*Alias
	}{
		Path:  e.marshalPath(),
		Alias: (*Alias)(e),
	})
}

// Helper function to marshal Path
func (e *GraphqlError) marshalPath() json.RawMessage {
	if len(e.Path) == 0 {
		return nil
	}
	pathJSON, err := json.Marshal(e.Path)
	if err != nil {
		return nil
	}
	return pathJSON
}

// ExecuteGraphQL executes a GraphQL operation
func ExecuteGraphQL(ctx context.Context, schema *ast.Document, operation *ast.OperationDefinition, variables map[string]interface{}) ([]byte, []GraphqlError) {
	executor := execution.NewExecutor(nil)

	rootNode := &CustomRootNode{
		schema:        schema,
		operationType: operation.OperationType,
	}

	var buf strings.Builder
	report := &operationreport.Report{}
	executionContext := execution.Context{
		Context:   ctx,
		Variables: convertVariables(variables),
	}

	err := executor.Execute(executionContext, rootNode, &buf)

	if err != nil {
		return nil, []GraphqlError{{Message: err.Error()}}
	}

	if report.HasErrors() {
		return nil, ConvertToGraphQLErrors(report.ExternalErrors)
	}

	return []byte(buf.String()), nil
}
