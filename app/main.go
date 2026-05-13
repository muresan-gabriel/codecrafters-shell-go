package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/codecrafters-io/shell-starter-go/lib"
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
		command := splitInput[0]
		args := splitInput[1:]

		switch command {
		case lib.Command[lib.Exit]:
			return
		case lib.Command[lib.Echo]:
			lib.EchoHandler(args)
		case lib.Command[lib.Type]:
			lib.TypeHandler(args[0])
		default:
			lib.DefaultHandler(input, args)
		}
	}
}
