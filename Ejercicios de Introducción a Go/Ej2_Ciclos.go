package main

import (
	"fmt"
)

func main() {
	for pos, char := range "mc1992♀♫" {
		fmt.Printf("Character %#U starts at byte position %d\n", char, pos)
	}
}
