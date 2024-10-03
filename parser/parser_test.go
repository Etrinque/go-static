package parser

import (
	"fmt"
	"testing"
)

func TestReadFile(t *testing.T) {
	file, _ := ReadFile("C:/Users/SOLVD/github_repos/go-static/content/markdown-cheat-sheet.md")
	// p := parser.NewParser(file)
	fmt.Println(string(file))
}
