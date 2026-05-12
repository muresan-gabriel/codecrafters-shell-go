package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Print("$ ")
	scanner := bufio.NewScanner(os.Stdin)
	input := scanner.Scan()
	fmt.Printf("%s: command not found", input)
}
