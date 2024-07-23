package emptypb

import (
	"github.com/wundergraph/graphql-go-tools/pkg/ast"
	"github.com/wundergraph/graphql-go-tools/pkg/astparser"
	"google.golang.org/protobuf/types/known/emptypb"
)

// Expose Google defined ptypes as this package types
type Empty = emptypb.Empty

var (
	gql__type_Empty  *ast.ObjectTypeDefinition
	gql__input_Empty *ast.InputObjectTypeDefinition
)

func Gql__type_Empty() *ast.ObjectTypeDefinition {
	if gql__type_Empty == nil {
		inputString := `
		"""
		Represents an empty type
		"""
		type Google_type_Empty {
			_: Boolean
		}
		`
		doc, report := astparser.ParseGraphqlDocumentString(inputString)
		if report.HasErrors() {
			panic("Failed to parse Google_type_Empty: " + report.Error())
		}
		if len(doc.ObjectTypeDefinitions) > 0 {
			gql__type_Empty = &doc.ObjectTypeDefinitions[0]
		} else {
			panic("Failed to find Google_type_Empty in parsed document")
		}
	}
	return gql__type_Empty
}

func Gql__input_Empty() *ast.InputObjectTypeDefinition {
	if gql__input_Empty == nil {
		inputString := `
		"""
		Represents an empty input type
		"""
		input Google_input_Empty {
			_: Boolean
		}
		`
		doc, report := astparser.ParseGraphqlDocumentString(inputString)
		if report.HasErrors() {
			panic("Failed to parse Google_input_Empty: " + report.Error())
		}
		if len(doc.InputObjectTypeDefinitions) > 0 {
			gql__input_Empty = &doc.InputObjectTypeDefinitions[0]
		} else {
			panic("Failed to find Google_input_Empty in parsed document")
		}
	}
	return gql__input_Empty
}

// AddEmptyTypeToSchema adds the Empty type definitions to the given schema
func AddEmptyTypeToSchema(schema *ast.Document) {
	schema.ObjectTypeDefinitions = append(schema.ObjectTypeDefinitions, *Gql__type_Empty())
	schema.InputObjectTypeDefinitions = append(schema.InputObjectTypeDefinitions, *Gql__input_Empty())
}
