package runtime

import (
	"github.com/wundergraph/graphql-go-tools/pkg/ast"
	"github.com/wundergraph/graphql-go-tools/pkg/execution"
)

// CustomRootNode implements the execution.RootNode interface
type CustomRootNode struct {
	schema        *ast.Document
	operationType ast.OperationType
	fields        []execution.Field
}

func (n *CustomRootNode) Kind() execution.NodeKind {
	return execution.ObjectKind
}

func (n *CustomRootNode) OperationType() ast.OperationType {
	return n.operationType
}

func (n *CustomRootNode) Schema() *ast.Document {
	return n.schema
}

func (n *CustomRootNode) Fields() []execution.Field {
	return n.fields
}

func (n *CustomRootNode) HasResolversRecursively() bool {
	for _, field := range n.fields {
		if field.HasResolversRecursively() {
			return true
		}
	}
	return false
}

// NewCustomRootNode creates a new CustomRootNode
func NewCustomRootNode(schema *ast.Document, opType ast.OperationType, fields []execution.Field) *CustomRootNode {
	return &CustomRootNode{
		schema:        schema,
		operationType: opType,
		fields:        fields,
	}
}
