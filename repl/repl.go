package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/Jihyun3478/saegim-lang/evaluator"
	"github.com/Jihyun3478/saegim-lang/lexer"
	"github.com/Jihyun3478/saegim-lang/object"
	"github.com/Jihyun3478/saegim-lang/parser"
	"github.com/Jihyun3478/saegim-lang/token"
)

const PROMPT = ">> "

const SAEGIM_LOGO = `
 ███████╗ █████╗ ███████╗ ██████╗ ██╗███╗   ███╗
 ██╔════╝██╔══██╗██╔════╝██╔════╝ ██║████╗ ████║
 ███████╗███████║█████╗  ██║  ███╗██║██╔████╔██║
 ╚════██║██╔══██║██╔══╝  ██║   ██║██║██║╚██╔╝██║
 ███████║██║  ██║███████╗╚██████╔╝██║██║ ╚═╝ ██║
 ╚══════╝╚═╝  ╚═╝╚══════╝ ╚═════╝ ╚═╝╚═╝     ╚═╝
`

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	fmt.Fprintln(out, SAEGIM_LOGO)
	fmt.Fprintln(out, "새김 언어 인터프리터")
	fmt.Fprintln(out)
	fmt.Fprintln(out, "모드를 선택하세요")
	fmt.Fprintln(out, "1. 표준어 버전")
	fmt.Fprintln(out, "2. 충청도 방언 버전")
	fmt.Fprint(out, "\n선택: ")

	scanner.Scan()
	mode := scanner.Text()

	var keywords token.KeywordSet
	var builtins map[string]*object.Builtin

	if mode == "1" {
		keywords = token.StandardKeywords
		builtins = evaluator.StandardBuiltins
		fmt.Fprintln(out, "\n표준어 모드")
	} else if mode == "2" {
		keywords = token.ChungcheongKeywords
		builtins = evaluator.ChungcheongBuiltins
		fmt.Fprintln(out, "\n충청도 모드")
	} else {
		fmt.Fprintln(out, "잘못된 선택입니다.")
		return
	}

	evaluator.SetBuiltins(builtins)

	fmt.Fprintln(out)
	fmt.Fprintln(out, "코드를 입력하세요 (종료 - exit)")
	fmt.Fprintln(out, "-----------------------------")

	env := object.NewEnvironment()

	for {
		fmt.Fprint(out, PROMPT)

		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()

		if line == "exit" || line == "종료" {
			fmt.Fprintln(out, "안녕히 가세요.")
			return
		}

		if line == "" {
			continue
		}

		l := lexer.New(line, keywords)
		p := parser.New(l)
		program := p.ParseProgram()

		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		evaluated := evaluator.Eval(program, env)
		if evaluated != nil {
			if evaluated.Type() == object.ERROR_OBJ {
				io.WriteString(out, "오류: ")
				io.WriteString(out, evaluated.Inspect())
				io.WriteString(out, "\n")
			} else if evaluated.Type() != object.NULL_OBJ {
				io.WriteString(out, evaluated.Inspect())
				io.WriteString(out, "\n")
			}
		}
	}
}

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, "파서 오류:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
