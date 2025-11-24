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

var builtins = StandardBuiltins

func SetBuiltins(buildinMap map[string]*object.Builtin) {
	builtins = buildinMap
}

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

	case *ast.ReturnStatement:
		value := Eval(node.ReturnValue, environment)
		if isError(value) {
			return value
		}
		return &object.ReturnValue{Value: value}
	
	case *ast.Identifier:
		return evalIdentifier(node, environment)
	
	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}

	case *ast.Boolean:
		return nativeBoolToBooleanObject(node.Value)

	case *ast.FunctionalLiteral:
		parameters := node.Parameters
		body := node.Body
		return &object.Function{Parameters: parameters, Env: environment, Body: body}

	case *ast.CallExpression:
		function := Eval(node.Function, environment)
		if isError(function) {
			return function
		}

		arguments := evalExpressions(node.Arguments, environment)
		if len(arguments) == 1 && isError(arguments[0]) {
			return arguments[0]
		}

		return applyFunction(function, arguments)

	case *ast.StringLiteral:
		return &object.String{Value: node.Value}
	
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

		switch result := result.(type) {
		case *object.ReturnValue:
			return result.Value
		case *object.Error:
			return result
		}
	}

	return result
}

func evalIdentifier(node *ast.Identifier, environment *object.Environment) object.Object {
	if value, ok := environment.Get(node.Value); ok {
		return value
	}

	if builtin, ok := builtins[node.Value]; ok {
		return builtin
	}

	return newError("identifier not found: " + node.Value)
}

func evalInfixExpression(operator string, left, right object.Object) object.Object {
	if left.Type() == object.INTEGER_OBJ && right.Type() == object.INTEGER_OBJ {
		return evalIntegerInfixExpression(operator, left, right)
	}

	if left.Type() == object.STRING_OBJ && right.Type() == object.STRING_OBJ {
		return evalStringInfixExpression(operator, left, right)
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
	case "<=":
		return nativeBoolToBooleanObject(leftValue <= rightValue)
	case ">=":
		return nativeBoolToBooleanObject(leftValue >= rightValue)
	case "==":
		return nativeBoolToBooleanObject(leftValue == rightValue)
	case "!=":
		return nativeBoolToBooleanObject(leftValue != rightValue)
	default:
		return newError(fmt.Sprintf("unknown operator: %s %s %s", left.Type(), operator, right.Type()))
	}
}

func evalStringInfixExpression(operator string, left, right object.Object) object.Object {
	if operator != "+" {
		return newError(fmt.Sprintf("unknown operator: %s %s %s", left.Type(), operator, right.Type()))
	}
	
	leftValue := left.(*object.String).Value
	rightValue := right.(*object.String).Value
	return &object.String{Value: leftValue + rightValue}
}

func evalBlockStatement(block *ast.BlockStatement, environment *object.Environment) object.Object {
	var result object.Object

	for _, statement := range block.Statements {
		result = Eval(statement, environment)

		if result != nil {
			resultType := result.Type()
			if resultType == object.RETURN_VALUE_OBJ || resultType == object.ERROR_OBJ {
				return result
			}
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

func evalExpressions(expressions []ast.Expression, environment *object.Environment) []object.Object {
	var result []object.Object

	for _, expression := range expressions {
		evaluated := Eval(expression, environment)
		if isError(evaluated) {
			return []object.Object{evaluated}
		}
		result = append(result, evaluated)
	}

	return result
}

func applyFunction(function object.Object, arguments []object.Object) object.Object {
	switch function := function.(type) {

	case *object.Function:
		extendedEnv := extendFunctionEnv(function, arguments)
		evaluated := Eval(function.Body, extendedEnv)
		return unwrapReturnValue(evaluated)
	
	case *object.Builtin:
		return function.Fn(arguments...)

	default:
		return newError(fmt.Sprintf("not a function: %s", function.Type()))
	}
}

func extendFunctionEnv(fn *object.Function, arguments []object.Object) *object.Environment {
	environment := object.NewEnclosedEnvironment(fn.Env)

	for parameterIndex, parameter := range fn.Parameters {
		environment.Set(parameter.Value, arguments[parameterIndex])
	}

	return environment
}

func unwrapReturnValue(obj object.Object) object.Object {
	if returnValue, ok := obj.(*object.ReturnValue); ok {
		return returnValue.Value
	}
	return obj
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
