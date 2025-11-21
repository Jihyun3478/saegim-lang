package object

import "fmt"

type ObjectType string

const (
	INTEGER_OBJ = "INTEGER"
	NULL_OBJ = "NULL"
	ERROR_OBJ = "ERROR"
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
