package parser

import (
    "testing"
    
    "github.com/Jihyun3478/saegim-lang/ast"
    "github.com/Jihyun3478/saegim-lang/lexer"
)

func TestVariableStatements(t *testing.T) {
    input := `
	변수 생년월일 = 20010101;
	변수 나이 = 25;
	`
    
    l := lexer.New(input)
    p := New(l)
    
    program := p.ParseProgram()
    checkParserErrors(t, p)
    
    if len(program.Statements) != 2 {
        t.Fatalf("program.Statements does not contain 2 statements. got=%d",
            len(program.Statements))
    }
    
    tests := []struct {
        expectedIdentifier string
    }{
        {"생년월일"},
        {"나이"},
    }
    
    for i, tt := range tests {
        stmt := program.Statements[i]
        if !testVariableStatement(t, stmt, tt.expectedIdentifier) {
            return
        }
    }
}

func testVariableStatement(t *testing.T, s ast.Statement, name string) bool {
    if s.TokenLiteral() != "변수" {
        t.Errorf("s.TokenLiteral not '변수'. got=%q", s.TokenLiteral())
        return false
    }
    
    varStmt, ok := s.(*ast.VariableStatement)
    if !ok {
        t.Errorf("s not *ast.VariableStatement. got=%T", s)
        return false
    }
    
    if varStmt.Name.Value != name {
        t.Errorf("varStmt.Name.Value not '%s'. got=%s", name, varStmt.Name.Value)
        return false
    }
    
    return true
}

func checkParserErrors(t *testing.T, p *Parser) {
    errors := p.Errors()
    if len(errors) == 0 {
        return
    }
    
    t.Errorf("parser has %d errors", len(errors))
    for _, msg := range errors {
        t.Errorf("parser error: %q", msg)
    }
    t.FailNow()
}
