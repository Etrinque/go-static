package lexer

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
