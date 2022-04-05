package main

import (
	"fmt"
	"os"
)

func createFile(p string) *os.File { // p = path
	fmt.Println("Creating")
	f, _ := os.Create(p)
	return f
}

func writeFile(f *os.File, text string) {
	fmt.Println("Writing")
	fmt.Fprintln(f, text)
}

func closeFile(f *os.File) {
	fmt.Println("Closing")
	f.Close()
}

func main() {
	f := createFile("C:\\Users\\naty-\\OneDrive\\Documentos\\2-JULIETA\\3ro UCC\\ARQUITECTURA DE SOFTWARE 1\\arq-software\\XX.txt")
	defer closeFile(f)
	writeFile(f, "I am Iron Man")
}
