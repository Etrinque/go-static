package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("Static Site gen")
	var siteGenError = errors.New("error gen item")
	fmt.Print(siteGenError)
}
