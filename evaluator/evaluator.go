package evaluator

import (
	"fmt"

	"github.com/Jihyun3478/saegim-lang/ast"
	"github.com/Jihyun3478/saegim-lang/object"
)

var (
	NULL = &object.Null{}
	TRUE = &object.Boolean{Value: true}
	FALSE = &object.Boolean{Value: false}
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

	case *ast.BlockStatement:
		return evalBlockStatement(node, environment)

	case *ast.IfExpression:
		return evalIfExpression(node, environment)
	
	case *ast.Identifier:
		return evalIdentifier(node, environment)
	
	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}

	case *ast.Boolean:
		return nativeBoolToBooleanObject(node.Value)

	case *ast.InfixExpression:
		left := Eval(node.Left, environment)
		if isError(left) {
			return  left
		}

		right := Eval(node.Right, environment)
		if isError(right) {
			return right
		}

		return evalInfixExpression(node.Operator, left, right)
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

func evalInfixExpression(operator string, left, right object.Object) object.Object {
	if left.Type() == object.INTEGER_OBJ && right.Type() == object.INTEGER_OBJ {
		return evalIntegerInfixExpression(operator, left, right)
	}

	if left.Type() != right.Type() {
		return newError(fmt.Sprintf("type mismatch: %s %s %s", left.Type(), operator, right.Type()))
	}

	return newError(fmt.Sprintf("unknown operator: %s %s %s", left.Type(), operator, right.Type()))
}

func evalIntegerInfixExpression(operator string, left, right object.Object) object.Object {
	leftValue := left.(*object.Integer).Value
	rightValue := right.(*object.Integer).Value

	switch operator {
	case "+":
		return &object.Integer{Value: leftValue + rightValue}
	case "-":
		return &object.Integer{Value: leftValue - rightValue}
	case "*":
		return &object.Integer{Value: leftValue * rightValue}
	case "/":
		return &object.Integer{Value: leftValue / rightValue}
	case "<":
		return nativeBoolToBooleanObject(leftValue < rightValue)
	case ">":
		return nativeBoolToBooleanObject(leftValue > rightValue)
	case "==":
		return nativeBoolToBooleanObject(leftValue == rightValue)
	case "!=":
		return nativeBoolToBooleanObject(leftValue != rightValue)
	default:
		return newError(fmt.Sprintf("unknown operator: %s %s %s", left.Type(), operator, right.Type()))
	}
}

func evalBlockStatement(block *ast.BlockStatement, environment *object.Environment) object.Object {
	var result object.Object

	for _, statement := range block.Statements {
		result = Eval(statement, environment)

		if result != nil && isError(result) {
			return result
		}
	}

	return result
}

func evalIfExpression(ifExpression *ast.IfExpression, environment *object.Environment) object.Object {
	condition := Eval(ifExpression.Condition, environment)
	if isError(condition) {
		return  condition
	}

	if isTruthy(condition) {
		return Eval(ifExpression.Consequence, environment)
	} else if ifExpression.Alternative != nil {
		return Eval(ifExpression.Alternative, environment)
	} else {
		return NULL
	}
}

func nativeBoolToBooleanObject(input bool) *object.Boolean {
	if input {
		return TRUE
	}
	return FALSE
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

func isTruthy(obj object.Object) bool {
	switch obj {
	case NULL:
		return false
	case TRUE:
		return true
	case FALSE:
		return false
	default:
		return true
	}
}
