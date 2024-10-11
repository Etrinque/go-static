package lexer

import (
	"bytes"
	"go-static/token"
)

/*
package Lexer is responsiple for the lexical analysis of MD files.  It implements the ast/mdnode elements. Lexer is implemented by the Parser pkg. It crawls through the document char by char and tracks it position and relative chars
*/

type Lexer struct {
	Input        []byte
	Position     int
	ReadPosition int
	Char         byte
}

func (l *Lexer) New(input []byte) *Lexer {
	lex := &Lexer{
		Input:        input,
		Position:     0,
		ReadPosition: 0,
	}
	lex.readChar()
	return lex
}

func (lex *Lexer) readChar() {
	if lex.ReadPosition >= len(lex.Input) {
		lex.Char = 0
	} else {
		lex.Char = lex.Input[lex.ReadPosition]
	}
	lex.Position = lex.ReadPosition
	lex.ReadPosition += 1
}

func (lex *Lexer) nextChar(char byte) {}

func (lex *Lexer) peekChar() {}

// checkToken iterates over slice of tags (tokens) to find and build a propper tag for passing into parser for node building
func (lex *Lexer) checkToken() (bool, string) {
	ch := lex.Char
	for _, s := range token.TokenList {
		if bytes.ContainsAny([]byte(s), string(ch)) {
			return true, s
		}
	}
	return false, ""
}

func (lex *Lexer) isAlpha(char byte) bool {
	if char >= 'a' && char <= 'z' || char >= 'A' && char <= 'Z' {
		return true
	}

	return false
}

func (lex *Lexer) isDigit(char byte) bool {
	if char >= '0' && char <= '9' {
		return true
	}

	return false
}
