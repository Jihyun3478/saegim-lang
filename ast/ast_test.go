package ast

import (
	"testing"
	
	"github.com/Jihyun3478/saegim-lang/token"
)

func TestString(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			&VariableStatement{
				Token: token.Token{Type: token.VARIABLE, Literal: "변수"},
				Name: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "myVar"},
					Value: "myVar",
				},
				Value: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "anotherVar"},
					Value: "anotherVar",
				},
			},
		},
	}

	if program.String() != "변수 myVar = anotherVar;" {
		t.Errorf("program.String() wrong. got=%q", program.String())
	}
}
