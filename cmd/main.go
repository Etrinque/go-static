package main

import (
	"fmt"
)

const DIST string = "./dist"
const CONTENT string = "./content"
const TEMPLATE string = "./template.html"
const STATIC string = "./static"
const SITE_GEN_ERROR string = "error gen item"

func main() {
	// var err = errors.New(SITE_GEN_ERROR)
	// fmt.Println("Static Site gen")
	// fmt.Println(err)
	// fmt.Println(DIST)

	fmt.Println("Copying Files To Distribution Directory")
	copyFiles(STATIC, DIST)

	fmt.Println("Generating Sites From Content Directory")
	genContent(CONTENT, TEMPLATE, DIST)

}
