package lexer

import (
	"fmt"
	"testing"
)

func TestReadChar(t *testing.T) {
	var l *Lexer
	var tt string = "testing 1 2 3 4 {} <> () [] \r \n \t"
	// tb := fmt.Sprintf("%b", tt)
	var lex = l.New([]byte(tt))

	for i := range tt {
		// pos := lex.Position
		// rd := lex.ReadPosition
		ch := lex.Char

		// fmt.Println(pos)
		// fmt.Println(rd)
		fmt.Printf("expect: %v, got=%s  \n", string(tt[i]), string(ch))

		lex.readChar()
	}

	t.Logf("passed string: %s", tt)
}
