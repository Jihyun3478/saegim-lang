package evaluator

import (
    "testing"
    
    "github.com/Jihyun3478/saegim-lang/lexer"
    "github.com/Jihyun3478/saegim-lang/parser"
    "github.com/Jihyun3478/saegim-lang/object"
)

func TestEvalIntegerExpression(t *testing.T) {
    tests := []struct {
        input    string
        expected int64
    }{
        {"5", 5},
        {"10", 10},
        {"5 + 5", 10},
        {"5 - 3", 2},
        {"2 * 3", 6},
        {"10 / 2", 5},
        {"5 + 2 * 3", 11},
        {"(5 + 2) * 3", 21},
    }
    
    for _, tt := range tests {
        evaluated := testEval(tt.input)
        testIntegerObject(t, evaluated, tt.expected)
    }
}

func TestVariableStatements(t *testing.T) {
    input := `
    변수 나이 = 25;
    나이;
    `
    
    evaluated := testEval(input)
    testIntegerObject(t, evaluated, 25)
}

func TestEvalBooleanExpression(t *testing.T) {
    tests := []struct {
        input    string
        expected bool
    }{
        {"5 > 3", true},
        {"5 < 3", false},
        {"5 == 5", true},
        {"5 != 5", false},
    }
    
    for _, tt := range tests {
        evaluated := testEval(tt.input)
        testBooleanObject(t, evaluated, tt.expected)
    }
}

func TestIfElseExpressions(t *testing.T) {
    tests := []struct {
        input    string
        expected interface{}
    }{
        {"만약 (참) { 10 }", 10},
        {"만약 (거짓) { 10 }", nil},
        {"만약 (1) { 10 }", 10},
        {"만약 (1 < 2) { 10 }", 10},
        {"만약 (1 > 2) { 10 }", nil},
        {"만약 (1 > 2) { 10 } 아니면 { 20 }", 20},
        {"만약 (1 < 2) { 10 } 아니면 { 20 }", 10},
    }
    
    for _, tt := range tests {
        evaluated := testEval(tt.input)
        integer, ok := tt.expected.(int)
        if ok {
            testIntegerObject(t, evaluated, int64(integer))
        } else {
            testNullObject(t, evaluated)
        }
    }
}

func TestFunctionObject(t *testing.T) {
    input := "함수(나이) { 나이 + 1; };"
    
    evaluated := testEval(input)
    
    fn, ok := evaluated.(*object.Function)
    if !ok {
        t.Fatalf("object is not Function. got=%T (%+v)", evaluated, evaluated)
    }
    
    if len(fn.Parameters) != 1 {
        t.Fatalf("function has wrong parameters. Parameters=%+v",
            fn.Parameters)
    }
    
    if fn.Parameters[0].String() != "나이" {
        t.Fatalf("parameter is not '나이'. got=%q", fn.Parameters[0])
    }
}

func TestFunctionApplication(t *testing.T) {
    tests := []struct {
        input    string
        expected int64
    }{
        {"변수 항등 = 함수(나이) { 나이; }; 항등(5);", 5},
        {"변수 항등 = 함수(나이) { 반환 나이; }; 항등(5);", 5},
        {"변수 더블 = 함수(나이) { 나이 * 2; }; 더블(5);", 10},
        {"변수 더하기 = 함수(첫째, 둘째) { 첫째 + 둘째; }; 더하기(5, 5);", 10},
        {"변수 더하기 = 함수(첫째, 둘째) { 첫째 + 둘째; }; 더하기(5 + 5, 더하기(5, 5));", 20},
        {"함수(나이) { 나이; }(5)", 5},
    }
    
    for _, tt := range tests {
        testIntegerObject(t, testEval(tt.input), tt.expected)
    }
}

func TestReturnStatements(t *testing.T) {
    tests := []struct {
        input    string
        expected int64
    }{
        {"반환 10;", 10},
        {"반환 10; 9;", 10},
        {"반환 2 * 5; 9;", 10},
        {"9; 반환 2 * 5; 9;", 10},
    }
    
    for _, tt := range tests {
        evaluated := testEval(tt.input)
        testIntegerObject(t, evaluated, tt.expected)
    }
}

func TestStringLiteral(t *testing.T) {
    input := `"안녕하세요"`
    
    evaluated := testEval(input)
    str, ok := evaluated.(*object.String)
    if !ok {
        t.Fatalf("object is not String. got=%T (%+v)", evaluated, evaluated)
    }
    
    if str.Value != "안녕하세요" {
        t.Errorf("String has wrong value. got=%q", str.Value)
    }
}

func TestBuiltinFunctions(t *testing.T) {
    tests := []struct {
        input    string
        expected interface{}
    }{
        {`출력(25)`, nil},
        {`출력("안녕")`, nil},
    }
    
    for _, tt := range tests {
        testEval(tt.input)
    }
}

func testEval(input string) object.Object {
    l := lexer.New(input)
    p := parser.New(l)
    program := p.ParseProgram()
    env := object.NewEnvironment()
    
    return Eval(program, env)
}

func testIntegerObject(t *testing.T, obj object.Object, expected int64) bool {
    result, ok := obj.(*object.Integer)
    if !ok {
        t.Errorf("object is not Integer. got=%T (%+v)", obj, obj)
        return false
    }
    
    if result.Value != expected {
        t.Errorf("object has wrong value. got=%d, want=%d",
            result.Value, expected)
        return false
    }
    
    return true
}

func testBooleanObject(t *testing.T, obj object.Object, expected bool) bool {
    result, ok := obj.(*object.Boolean)
    if !ok {
        t.Errorf("object is not Boolean. got=%T (%+v)", obj, obj)
        return false
    }
    
    if result.Value != expected {
        t.Errorf("object has wrong value. got=%t, want=%t",
            result.Value, expected)
        return false
    }
    
    return true
}

func testNullObject(t *testing.T, obj object.Object) bool {
    if obj != NULL {
        t.Errorf("object is not NULL. got=%T (%+v)", obj, obj)
        return false
    }
    return true
}
