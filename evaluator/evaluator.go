package evaluator

import (
	"github.com/Jihyun3478/saegim-lang/ast"
	"github.com/Jihyun3478/saegim-lang/object"
)

var (
	NULL = &object.Null{}
)

func Eval(node ast.Node, environment *object.Environment) object.Object {
	switch node := node.(type) {
	case *ast.Program:
		return evalProgram(node, environment)
	
	case *ast.VariableStatement:
		value := Eval(node.Value, environment)
		if isError(value) {
			return value
		}
		environment.Set(node.Name.Value, value)
		return NULL

	case *ast.ExpressionStatement:
		return Eval(node.Expression, environment)
	
	case *ast.Identifier:
		return evalIdentifier(node, environment)
	
	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}
	}

	return NULL
}

func evalProgram(program *ast.Program, environment *object.Environment) object.Object {
	var result object.Object

	for _, statement := range program.Statements {
		result = Eval(statement, environment)

		if isError(result) {
			return result
		}
	}

	return result
}

func evalIdentifier(node *ast.Identifier, environment *object.Environment) object.Object {
	value, ok := environment.Get(node.Value)
	if !ok {
		return newError("identifier not found: " + node.Value)
	}
	return value
}

func newError(message string) *object.Error {
	return &object.Error{Message: message}
}

func isError(obj object.Object) bool {
	if obj != nil {
		return obj.Type() == object.ERROR_OBJ
	}
	return false
}
