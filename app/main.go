package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	for true {
		fmt.Print("$ ")
		scanner := bufio.NewScanner(os.Stdin)
		if !scanner.Scan() {
			return
		}
		input := scanner.Text()
		fmt.Printf("%s: command not found\n", input)
	}
}
