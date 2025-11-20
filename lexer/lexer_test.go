package lexer

import (
	"testing"

	"github.com/Jihyun3478/saegim-lang/token"
)

func TestNextToken(t *testing.T) {
	input := `변수 나이 = 25;`

	tests := []struct {
		expectedType token.TokenType
		expectedLiteral string
	}{
		{token.VARIABLE, "변수"},
		{token.IDENT, "나이"},
		{token.ASSIGN, "="},
		{token.INT, "25"},
		{token.SEMICOLON, ";"},

		{token.EOF, ""},
	}

	lexer := New(input)

	for index, tt := range tests {
		tok := lexer.NextToken()
		if tok.Type  != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q", index, tt.expectedType, tok.Type)
		}
		if tok.Literal  != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", index, tt.expectedLiteral, tok.Literal)
		}
	}
}
