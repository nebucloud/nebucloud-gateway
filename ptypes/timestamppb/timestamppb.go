package timestamppb

import (
	"github.com/wundergraph/graphql-go-tools/pkg/ast"
	"github.com/wundergraph/graphql-go-tools/pkg/astparser"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// Expose Google defined ptypes as this package types
type Timestamp = timestamppb.Timestamp

var (
	gql__type_Timestamp  *ast.ObjectTypeDefinition
	gql__input_Timestamp *ast.InputObjectTypeDefinition
)

func Gql__type_Timestamp() *ast.ObjectTypeDefinition {
	if gql__type_Timestamp == nil {
		inputString := `
		"""
		Represents a timestamp with seconds and nanos
		"""
		type Google_type_Timestamp {
			seconds: Int!
			nanos: Int!
		}
		`
		doc, report := astparser.ParseGraphqlDocumentString(inputString)
		if report.HasErrors() {
			panic("Failed to parse Google_type_Timestamp: " + report.Error())
		}
		if len(doc.ObjectTypeDefinitions) > 0 {
			gql__type_Timestamp = &doc.ObjectTypeDefinitions[0]
		} else {
			panic("Failed to find Google_type_Timestamp in parsed document")
		}
	}
	return gql__type_Timestamp
}

func Gql__input_Timestamp() *ast.InputObjectTypeDefinition {
	if gql__input_Timestamp == nil {
		inputString := `
		"""
		Represents a timestamp input with seconds and nanos
		"""
		input Google_input_Timestamp {
			seconds: Int!
			nanos: Int!
		}
		`
		doc, report := astparser.ParseGraphqlDocumentString(inputString)
		if report.HasErrors() {
			panic("Failed to parse Google_input_Timestamp: " + report.Error())
		}
		if len(doc.InputObjectTypeDefinitions) > 0 {
			gql__input_Timestamp = &doc.InputObjectTypeDefinitions[0]
		} else {
			panic("Failed to find Google_input_Timestamp in parsed document")
		}
	}
	return gql__input_Timestamp
}

// AddTimestampTypeToSchema adds the Timestamp type definitions to the given schema
func AddTimestampTypeToSchema(schema *ast.Document) {
	schema.ObjectTypeDefinitions = append(schema.ObjectTypeDefinitions, *Gql__type_Timestamp())
	schema.InputObjectTypeDefinitions = append(schema.InputObjectTypeDefinitions, *Gql__input_Timestamp())
}
