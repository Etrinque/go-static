package parser

import (
	"fmt"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

var err = godotenv.Load("../.env")

var TEST_FILE string = os.Getenv("TEST_FILE")

func TestReadFile(t *testing.T) {
	fmt.Printf("this is the test file loc: %s", TEST_FILE)
	file, _ := ReadFile(TEST_FILE)
	// p := parser.NewParser(file)
	fmt.Println(string(file))
}
