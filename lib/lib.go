package lib

import "fmt"

type Commands int

const (
	Exit Commands = iota
	Echo
	Type
)

var Command = map[Commands]string{
	Exit: "exit",
	Echo: "echo",
	Type: "type",
}

func EchoHandler(input []string) {
	echoedString := ""
	for _, value := range input {
		echoedString += value + " "
	}
	fmt.Println(echoedString)
}

func TypeHandler(input string) {
	found := false
	for _, name := range Command {
		if name == input {
			found = true
			fmt.Printf("%s is a shell builtin\n", input)
			break
		}
	}

	if !found {
		fmt.Printf("%s: not found\n", input)
	}
}

func NotFoundHandler(input string) {
	fmt.Printf("%s: command not found\n", input)
}
