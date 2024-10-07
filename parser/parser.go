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
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	buf, err := os.ReadFile(file.Name())
	if err != nil {
		return nil, err
	}
	return buf, nil
}

func buildNode() *ast.MDNode {}

func buildTree() {}

// This need some research on file meta parsing, Notably os.FileInfo
// func GetFileMeta(file *os.File) *ast.MDMeta {
// 	return &ast.MDMeta{Name: file.Name()}
// }
