package evaluator

import (
	"fmt"

	"github.com/Jihyun3478/saegim-lang/object"
)

var ChungcheongBuiltins = map[string]*object.Builtin{
	"출력해유": &object.Builtin{
		Fn: func(arguments ...object.Object) object.Object {
			for _, argument := range arguments {
				fmt.Println(argument.Inspect())
			}
			return NULL
		},
	},
}
