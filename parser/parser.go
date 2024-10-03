package parser

import (
	"go-static/ast"
	"go-static/lexer"
	"os"
)

// package parser is responsible for Creating and implementing the parsing program.
// It defines the

type Parser struct {
	Lex      *lexer.Lexer
	currChar byte
	peekChar byte
}

type Parsed struct {
	Nodes []ast.MDNode
}

func NewParser(input []byte) *Parser {
	return &Parser{
		Lex: lexer.New(input),
	}
}

func ReadFile(path string) ([]byte, error) {
	buf, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return buf, nil
}
