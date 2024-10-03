package main

import (
	"fmt"
	"runtime"
)

const DIST string = "./dist"
const CONTENT string = "./content"
const TEMPLATE string = "./template.html"
const STATIC string = "./static"
const SITE_GEN_ERROR string = "error gen item"

func main() {
	fmt.Println("Static Site gen")

	// fmt.Println("Copying Files To Distribution Directory")
	// copyFiles(STATIC, DIST)

	// fmt.Println("Generating Sites From Content Directory")
	// genContent(CONTENT, TEMPLATE, DIST)

	// component := hello("Eric")
	// component.Render(context.Background(), os.Stdout)
	PrintMemUsage()

}

func PrintMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Print("\n==================\n")
	fmt.Printf("Alloc = %v MiB\n", bToMb(m.Alloc))
	fmt.Printf("TotalAlloc = %v MiB\n", bToMb(m.TotalAlloc))
	fmt.Printf("Sys = %v MiB\n", bToMb(m.Sys))
	fmt.Printf("NumGC = %v\n", m.NumGC)
	fmt.Print("==================\n")
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
