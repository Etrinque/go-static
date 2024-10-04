package parser

import (
	"fmt"
	"testing"
)

func TestReadFile(t *testing.T) {
	file, _ := ReadFile("")
	// p := parser.NewParser(file)
	fmt.Println(string(file))
}
