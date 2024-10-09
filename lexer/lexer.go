package lexer

import (
	"go-static/ast"
	"strings"
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

func New(input []byte) *Lexer {
	lex := &Lexer{
		Input: input,
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

func (lex *Lexer) buildTag(char byte) []byte {
	var tag []byte

	tag = append(tag, char)

	return tag
}

func (lex *Lexer) buildBlock(char byte) ([]byte, error) {
	var builder strings.Builder
	var block []byte

	block = append(block, char)

	_, err := builder.Write(block)
	if err != nil {
		return nil, err
	}

	return []byte(builder.String()), nil
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

func (lex *Lexer) CheckTagMap(char string, tagmap ast.MDTagMap) bool {
	for _, v := range tagmap {
		if strings.Contains(v.Open, char) {
			if char == v.Open {
				return true
			}
		} else {
			return false
		}
	}
	return false
}
