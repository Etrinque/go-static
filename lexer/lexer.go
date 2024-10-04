package lexer

import (
	"go-static/ast"
)

/*
package Lexer is responsiple for the lexical analysis of MD files.  It implements the ast/mdnode elements. Lexer is implemented by the Parser pkg. It crawls through the document char by char and tracks it position and relative chars
*/

type Lexer struct {
	input        []byte
	position     int
	readPosition int
	char         byte
}

func New(input []byte) *Lexer {
	lex := &Lexer{
		input: input,
	}
	lex.readChar()
	return lex
}

func (lex *Lexer) readChar() {
	if lex.readPosition >= len(lex.input) {
		lex.char = 0
	} else {
		lex.char = lex.input[lex.readPosition]
	}
	lex.position = lex.readPosition
	lex.readPosition += 1
}

func(lex *Lexer) nextChar() {}


// buildStr walks through input and builds tags as they come based on the tagMap
func(lex *Lexer) buildStr() string {
	return ""
}

func(lex *Lexer) checkTagMap(tagmap ast.MDTagMap) bool {
	var chars = "!["
	for k, _ := range(tagmap) {
		if chars != k {
			return false
		} else {
			switch chars {
			case "![":
				return true
			default:
				return true
			}
		}
	}

	return false
}
