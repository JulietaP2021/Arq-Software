// File Read
package main

import (
	"fmt"
	"os"
)

func main() {
	f, _ := os.ReadFile("C:\\Users\\naty-\\OneDrive\\Documentos\\2-JULIETA\\3ro UCC\\ARQUITECTURA DE SOFTWARE 1\\arq-software\\X.txt")
	fmt.Print(string(f))
}
