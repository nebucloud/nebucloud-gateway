package wrapperspb

import (
	"github.com/wundergraph/graphql-go-tools/pkg/ast"
	"github.com/wundergraph/graphql-go-tools/pkg/astparser"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

// Expose Google defined ptypes as this package types
type (
	DoubleValue = wrapperspb.DoubleValue
	FloatValue  = wrapperspb.FloatValue
	BoolValue   = wrapperspb.BoolValue
	Int32Value  = wrapperspb.Int32Value
	Int64Value  = wrapperspb.Int64Value
	UInt32Value = wrapperspb.UInt32Value
	UInt64Value = wrapperspb.UInt64Value
	StringValue = wrapperspb.StringValue
	BytesValue  = wrapperspb.BytesValue
)

var (
	gql__type_DoubleValue  *ast.ObjectTypeDefinition
	gql__type_FloatValue   *ast.ObjectTypeDefinition
	gql__type_Int64Value   *ast.ObjectTypeDefinition
	gql__type_Uint64Value  *ast.ObjectTypeDefinition
	gql__type_Int32Value   *ast.ObjectTypeDefinition
	gql__type_Uint32Value  *ast.ObjectTypeDefinition
	gql__type_BoolValue    *ast.ObjectTypeDefinition
	gql__type_StringValue  *ast.ObjectTypeDefinition
	gql__input_DoubleValue *ast.InputObjectTypeDefinition
	gql__input_FloatValue  *ast.InputObjectTypeDefinition
	gql__input_Int64Value  *ast.InputObjectTypeDefinition
	gql__input_Uint64Value *ast.InputObjectTypeDefinition
	gql__input_Int32Value  *ast.InputObjectTypeDefinition
	gql__input_Uint32Value *ast.InputObjectTypeDefinition
	gql__input_BoolValue   *ast.InputObjectTypeDefinition
	gql__input_StringValue *ast.InputObjectTypeDefinition
)

func parseTypeDefinition(name, typeName string) *ast.ObjectTypeDefinition {
	inputString := `
	"""
	Represents a ` + name + ` wrapper
	"""
	type Google_type_Wrappers_` + name + ` {
		value: ` + typeName + `!
	}
	`
	doc, report := astparser.ParseGraphqlDocumentString(inputString)
	if report.HasErrors() {
		panic("Failed to parse Google_type_Wrappers_" + name + ": " + report.Error())
	}
	if len(doc.ObjectTypeDefinitions) > 0 {
		return &doc.ObjectTypeDefinitions[0]
	}
	panic("Failed to find Google_type_Wrappers_" + name + " in parsed document")
}

func parseInputTypeDefinition(name, typeName string) *ast.InputObjectTypeDefinition {
	inputString := `
	"""
	Represents a ` + name + ` input wrapper
	"""
	input Google_input_Wrappers_` + name + ` {
		value: ` + typeName + `!
	}
	`
	doc, report := astparser.ParseGraphqlDocumentString(inputString)
	if report.HasErrors() {
		panic("Failed to parse Google_input_Wrappers_" + name + ": " + report.Error())
	}
	if len(doc.InputObjectTypeDefinitions) > 0 {
		return &doc.InputObjectTypeDefinitions[0]
	}
	panic("Failed to find Google_input_Wrappers_" + name + " in parsed document")
}

func Gql__type_DoubleValue() *ast.ObjectTypeDefinition {
	if gql__type_DoubleValue == nil {
		gql__type_DoubleValue = parseTypeDefinition("DoubleValue", "Float")
	}
	return gql__type_DoubleValue
}

func Gql__type_FloatValue() *ast.ObjectTypeDefinition {
	if gql__type_FloatValue == nil {
		gql__type_FloatValue = parseTypeDefinition("FloatValue", "Float")
	}
	return gql__type_FloatValue
}

func Gql__type_Int64Value() *ast.ObjectTypeDefinition {
	if gql__type_Int64Value == nil {
		gql__type_Int64Value = parseTypeDefinition("Int64Value", "Int")
	}
	return gql__type_Int64Value
}

func Gql__type_Uint64Value() *ast.ObjectTypeDefinition {
	if gql__type_Uint64Value == nil {
		gql__type_Uint64Value = parseTypeDefinition("Uint64Value", "Int")
	}
	return gql__type_Uint64Value
}

func Gql__type_Int32Value() *ast.ObjectTypeDefinition {
	if gql__type_Int32Value == nil {
		gql__type_Int32Value = parseTypeDefinition("Int32Value", "Int")
	}
	return gql__type_Int32Value
}

func Gql__type_Uint32Value() *ast.ObjectTypeDefinition {
	if gql__type_Uint32Value == nil {
		gql__type_Uint32Value = parseTypeDefinition("Uint32Value", "Int")
	}
	return gql__type_Uint32Value
}

func Gql__type_BoolValue() *ast.ObjectTypeDefinition {
	if gql__type_BoolValue == nil {
		gql__type_BoolValue = parseTypeDefinition("BoolValue", "Boolean")
	}
	return gql__type_BoolValue
}

func Gql__type_StringValue() *ast.ObjectTypeDefinition {
	if gql__type_StringValue == nil {
		gql__type_StringValue = parseTypeDefinition("StringValue", "String")
	}
	return gql__type_StringValue
}

func Gql__input_DoubleValue() *ast.InputObjectTypeDefinition {
	if gql__input_DoubleValue == nil {
		gql__input_DoubleValue = parseInputTypeDefinition("DoubleValue", "Float")
	}
	return gql__input_DoubleValue
}

func Gql__input_FloatValue() *ast.InputObjectTypeDefinition {
	if gql__input_FloatValue == nil {
		gql__input_FloatValue = parseInputTypeDefinition("FloatValue", "Float")
	}
	return gql__input_FloatValue
}

func Gql__input_Int64Value() *ast.InputObjectTypeDefinition {
	if gql__input_Int64Value == nil {
		gql__input_Int64Value = parseInputTypeDefinition("Int64Value", "Int")
	}
	return gql__input_Int64Value
}

func Gql__input_Uint64Value() *ast.InputObjectTypeDefinition {
	if gql__input_Uint64Value == nil {
		gql__input_Uint64Value = parseInputTypeDefinition("Uint64Value", "Int")
	}
	return gql__input_Uint64Value
}

func Gql__input_Int32Value() *ast.InputObjectTypeDefinition {
	if gql__input_Int32Value == nil {
		gql__input_Int32Value = parseInputTypeDefinition("Int32Value", "Int")
	}
	return gql__input_Int32Value
}

func Gql__input_Uint32Value() *ast.InputObjectTypeDefinition {
	if gql__input_Uint32Value == nil {
		gql__input_Uint32Value = parseInputTypeDefinition("Uint32Value", "Int")
	}
	return gql__input_Uint32Value
}

func Gql__input_BoolValue() *ast.InputObjectTypeDefinition {
	if gql__input_BoolValue == nil {
		gql__input_BoolValue = parseInputTypeDefinition("BoolValue", "Boolean")
	}
	return gql__input_BoolValue
}

func Gql__input_StringValue() *ast.InputObjectTypeDefinition {
	if gql__input_StringValue == nil {
		gql__input_StringValue = parseInputTypeDefinition("StringValue", "String")
	}
	return gql__input_StringValue
}

// AddWrapperTypesToSchema adds all Wrapper type definitions to the given schema
func AddWrapperTypesToSchema(schema *ast.Document) {
	schema.ObjectTypeDefinitions = append(schema.ObjectTypeDefinitions,
		*Gql__type_DoubleValue(),
		*Gql__type_FloatValue(),
		*Gql__type_Int64Value(),
		*Gql__type_Uint64Value(),
		*Gql__type_Int32Value(),
		*Gql__type_Uint32Value(),
		*Gql__type_BoolValue(),
		*Gql__type_StringValue(),
	)
	schema.InputObjectTypeDefinitions = append(schema.InputObjectTypeDefinitions,
		*Gql__input_DoubleValue(),
		*Gql__input_FloatValue(),
		*Gql__input_Int64Value(),
		*Gql__input_Uint64Value(),
		*Gql__input_Int32Value(),
		*Gql__input_Uint32Value(),
		*Gql__input_BoolValue(),
		*Gql__input_StringValue(),
	)
}
