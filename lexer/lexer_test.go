package lexer

import (
    "testing"

    "github.com/Jihyun3478/saegim-lang/token"
)

func TestNextToken(t *testing.T) {
    input := `
변수 나이 = 25;
변수 합 = 10 + 20;
변수 결과 = 5 > 3;
(5 + 2) * 3;
`

    tests := []struct {
        expectedType    token.TokenType
        expectedLiteral string
    }{
        {token.VARIABLE, "변수"},
        {token.IDENT, "나이"},
        {token.ASSIGN, "="},
        {token.INT, "25"},
        {token.SEMICOLON, ";"},
        
        {token.VARIABLE, "변수"},
        {token.IDENT, "합"},
        {token.ASSIGN, "="},
        {token.INT, "10"},
        {token.PLUS, "+"},
        {token.INT, "20"},
        {token.SEMICOLON, ";"},
        
        {token.VARIABLE, "변수"},
        {token.IDENT, "결과"},
        {token.ASSIGN, "="},
        {token.INT, "5"},
        {token.GT, ">"},
        {token.INT, "3"},
        {token.SEMICOLON, ";"},
        
        {token.LPAREN, "("},
        {token.INT, "5"},
        {token.PLUS, "+"},
        {token.INT, "2"},
        {token.RPAREN, ")"},
        {token.ASTERISK, "*"},
        {token.INT, "3"},
        {token.SEMICOLON, ";"},

        {token.EOF, ""},
    }

    lexer := New(input)

    for index, tt := range tests {
        tok := lexer.NextToken()
        if tok.Type != tt.expectedType {
            t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q", 
                index, tt.expectedType, tok.Type)
        }
        if tok.Literal != tt.expectedLiteral {
            t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", 
                index, tt.expectedLiteral, tok.Literal)
        }
    }
}
