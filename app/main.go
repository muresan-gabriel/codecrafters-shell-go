package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	for true {
		fmt.Print("$ ")
		scanner := bufio.NewScanner(os.Stdin)
		if !scanner.Scan() {
			return
		}
		input := scanner.Text()
		splitInput := strings.Split(input, " ")
		switch splitInput[0] {
		case "exit":
			return
		case "echo":
			echoedString := ""
			for _, value := range splitInput[1:] {
				echoedString += value + " "
			}
			fmt.Println(echoedString)
		default:
			fmt.Printf("%s: command not found\n", input)
		}
	}
}
