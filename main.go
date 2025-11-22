package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/Jihyun3478/saegim-lang/evaluator"
	"github.com/Jihyun3478/saegim-lang/lexer"
	"github.com/Jihyun3478/saegim-lang/object"
	"github.com/Jihyun3478/saegim-lang/parser"
	"github.com/Jihyun3478/saegim-lang/repl"
	"github.com/Jihyun3478/saegim-lang/token"
)

func main() {
	if len(os.Args) >= 2 {
		runFile(os.Args[1])
		return
	}

	repl.Start(os.Stdin, os.Stdout)
}

func runFile(filename string) {
	ext := filepath.Ext(filename)

	var keywords token.KeywordSet
	var builtins map[string]*object.Builtin

	switch ext {
	case ".sg":
		keywords = token.StandardKeywords
		builtins = evaluator.StandardBuiltins
	case ".hbg":
		keywords = token.ChungcheongKeywords
		builtins = evaluator.ChungcheongBuiltins
	default:
		fmt.Printf("지원하지 않는 확장자: %s\n", ext)
		fmt.Println("사용 가능: .sg (표준어), .hbg (충청도)")
		os.Exit(1)
	}

	content, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("파일 읽기 오류: %s\n", err)
		os.Exit(1)
	}

	evaluator.SetBuiltins(builtins)

	l := lexer.New(string(content), keywords)
	p := parser.New(l)
	program := p.ParseProgram()

	if len(p.Errors()) > 0 {
		fmt.Println("파서 오류:")
		for _, msg := range p.Errors() {
			fmt.Println("\t" + msg)
		}
		os.Exit(1)
	}

	env := object.NewEnvironment()
	result := evaluator.Eval(program, env)

	if result != nil && result.Type() == object.ERROR_OBJ {
		fmt.Println("실행 오류:", result.Inspect())
		os.Exit(1)
	}
}
