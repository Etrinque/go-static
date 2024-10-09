package lexer

import (
	"go-static/ast"
	"strings"
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

func (lex *Lexer) checkTagMap(char []byte, tagmap ast.MDTagMap) bool {
	var chars = string(char)
	for _, v := range tagmap {
		if strings.Contains(v.Open, chars) {
			if chars == v.Open {
				return true
			}
		} else {
			return false
		}
	}
	return false
}
