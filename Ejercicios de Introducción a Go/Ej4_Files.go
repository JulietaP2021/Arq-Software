// File Write
package main

import (
	"fmt"
	"os"
)

func main() {
	f, _ := os.Create("C:\\Users\\naty-\\OneDrive\\Documentos\\2-JULIETA\\3ro UCC\\ARQUITECTURA DE SOFTWARE 1\\arq-software\\X.txt")
	fmt.Fprintln(f, "Avengers, assemble!")
	f.Close()
}
