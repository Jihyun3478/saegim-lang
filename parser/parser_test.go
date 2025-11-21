package parser

import (
    "fmt"
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

func TestIntegerExpression(t *testing.T) {
    input := "5;"
    
    l := lexer.New(input)
    p := New(l)
    program := p.ParseProgram()
    checkParserErrors(t, p)
    
    if len(program.Statements) != 1 {
        t.Fatalf("program has not enough statements. got=%d",
            len(program.Statements))
    }
    
    stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
    if !ok {
        t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T",
            program.Statements[0])
    }
    
    literal, ok := stmt.Expression.(*ast.IntegerLiteral)
    if !ok {
        t.Fatalf("exp not *ast.IntegerLiteral. got=%T", stmt.Expression)
    }
    
    if literal.Value != 5 {
        t.Errorf("literal.Value not %d. got=%d", 5, literal.Value)
    }
}

func TestInfixExpressions(t *testing.T) {
    infixTests := []struct {
        input      string
        leftValue  int64
        operator   string
        rightValue int64
    }{
        {"5 + 5;", 5, "+", 5},
        {"5 - 5;", 5, "-", 5},
        {"5 * 5;", 5, "*", 5},
        {"5 / 5;", 5, "/", 5},
        {"5 > 5;", 5, ">", 5},
        {"5 < 5;", 5, "<", 5},
        {"5 == 5;", 5, "==", 5},
        {"5 != 5;", 5, "!=", 5},
    }
    
    for _, tt := range infixTests {
        l := lexer.New(tt.input)
        p := New(l)
        program := p.ParseProgram()
        checkParserErrors(t, p)
        
        if len(program.Statements) != 1 {
            t.Fatalf("program.Statements does not contain %d statements. got=%d\n",
                1, len(program.Statements))
        }
        
        stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
        if !ok {
            t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T",
                program.Statements[0])
        }
        
        exp, ok := stmt.Expression.(*ast.InfixExpression)
        if !ok {
            t.Fatalf("exp is not ast.InfixExpression. got=%T", stmt.Expression)
        }
        
        if !testIntegerLiteral(t, exp.Left, tt.leftValue) {
            return
        }
        
        if exp.Operator != tt.operator {
            t.Fatalf("exp.Operator is not '%s'. got=%s",
                tt.operator, exp.Operator)
        }
        
        if !testIntegerLiteral(t, exp.Right, tt.rightValue) {
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

func testIntegerLiteral(t *testing.T, il ast.Expression, value int64) bool {
    integ, ok := il.(*ast.IntegerLiteral)
    if !ok {
        t.Errorf("il not *ast.IntegerLiteral. got=%T", il)
        return false
    }
    
    if integ.Value != value {
        t.Errorf("integ.Value not %d. got=%d", value, integ.Value)
        return false
    }
    
    if integ.TokenLiteral() != fmt.Sprintf("%d", value) {
        t.Errorf("integ.TokenLiteral not %d. got=%s", value,
            integ.TokenLiteral())
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
