package main

import (
    "fmt"
	
	"github.com/Jihyun3478/saegim-lang/token"
	"github.com/Jihyun3478/saegim-lang/lexer"
	"github.com/Jihyun3478/saegim-lang/parser"
	"github.com/Jihyun3478/saegim-lang/object"
	"github.com/Jihyun3478/saegim-lang/evaluator"
)

func main() {
    fmt.Println("새김 언어 인터프리터")
    fmt.Println("1. 표준어")
    fmt.Println("2. 충청도")
    
    var mode int
    fmt.Scan(&mode)
    
    var keywords token.KeywordSet
    var builtins map[string]*object.Builtin
    
    if mode == 1 {
        keywords = token.StandardKeywords
        builtins = evaluator.StandardBuiltins
        fmt.Println("표준어 모드")
    } else {
        keywords = token.ChungcheongKeywords
        builtins = evaluator.ChungcheongBuiltins
        fmt.Println("충청도 모드")
    }
    
    evaluator.SetBuiltins(builtins)
    
    input := `변수 나이 = 25; 출력(나이);`
    
    l := lexer.New(input, keywords)
    p := parser.New(l)
    program := p.ParseProgram()
    env := object.NewEnvironment()
    
    evaluator.Eval(program, env)
}
