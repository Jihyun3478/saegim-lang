package object

import (
	"fmt"
	"strings"
	"bytes"
	
	"github.com/Jihyun3478/saegim-lang/ast"
)

type ObjectType string

const (
	INTEGER_OBJ = "INTEGER"
	BOOLEAN_OBJ = "BOOLEAN"
	NULL_OBJ = "NULL"
	ERROR_OBJ = "ERROR"
	RETURN_VALUE_OBJ = "RETURN_VALUE"
	FUNCTION_OBJ = "FUNCTION"
	STRING_OBJ = "STRING"
	BUILTIN_OBJ = "BUILTIN"
)

type Object interface {
	Type() ObjectType
	Inspect() string
}

type Integer struct {
	Value int64
}

func (integer *Integer) Type() ObjectType {
	return INTEGER_OBJ
}

func (integer *Integer) Inspect() string {
	return fmt.Sprintf("%d", integer.Value)
}

type Boolean struct {
	Value bool
}

func (boolean *Boolean) Type() ObjectType {
	return BOOLEAN_OBJ
}

func (boolean *Boolean) Inspect() string {
	return fmt.Sprintf("%t", boolean.Value)
}

type Null struct{}

func (null *Null) Type() ObjectType {
	return NULL_OBJ
}

func (null *Null) Inspect() string {
	return "null"
}

type Error struct {
	Message string
}

func (error *Error) Type() ObjectType {
	return ERROR_OBJ
}

func (error *Error) Inspect() string {
	return "ERROR: " + error.Message
}

type ReturnValue struct {
	Value Object
}

func (returnValue *ReturnValue) Type() ObjectType {
	return RETURN_VALUE_OBJ
}

func (returnValue *ReturnValue) Inspect() string {
	return returnValue.Value.Inspect()
}

type Function struct {
	Parameters []*ast.Identifier
	Body *ast.BlockStatement
	Env *Environment
}

func (function *Function) Type() ObjectType {
	return FUNCTION_OBJ
}

func (function *Function) Inspect() string {
	var out bytes.Buffer

	parameters := []string{}
	for _, p := range function.Parameters {
		parameters = append(parameters, p.String())
	}

	out.WriteString("함수")
	out.WriteString("(")
	out.WriteString(strings.Join(parameters, ", "))
	out.WriteString(") {\n")
	out.WriteString(function.Body.String())
	out.WriteString("\n}")

	return out.String()
}

type String struct {
	Value string
}

func (string *String) Type() ObjectType {
	return STRING_OBJ
}

func (string *String) Inspect() string {
	return string.Value
}

type BuiltinFunction func(arguments ...Object) Object

type Builtin struct {
	Fn BuiltinFunction
}

func (builtin *Builtin) Type() ObjectType {
	return BUILTIN_OBJ
}

func (builtin *Builtin) Inspect() string {
	return "builtin function"
}
