package main

import (
	"fmt"
)

func main() {
	fmt.Print("$ ")
	input, _ := fmt.Scanln()
	fmt.Printf("%s: command not found", input)
}
