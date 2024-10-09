package lexer

import (
	"fmt"
	"go-static/ast"
	"testing"
)

func TestCheckTagmap(t *testing.T) {
	var lex *Lexer
	var tt = []string{"**", "*", "***", "a", "# ", "### "}

	var testMap = ast.MDTagMap.GenMap(ast.MDTagMap{})

	// map[string]struct{ open, close string }{
	// 	"NONE":            {open: "", close: ""},
	// 	"HEADING_1":       {open: "# ", close: ""},
	// 	"HEADING_3":       {open: "### ", close: ""},
	// 	"BOLD":            {open: "**", close: "**"},
	// 	"ITALIC":          {open: "*", close: "*"},
	// 	"BOLD_AND_ITALIC": {open: "***", close: "***"},
	// }

	for _, s := range tt {
		for _, v := range testMap {
			ok := lex.CheckTagMap(s, v)
			if !ok {
				fmt.Println("cannot locate tag")
			}

		}
	}
}
