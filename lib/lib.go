package lib

import (
	"fmt"
	"os"
	"strings"
)

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

func isExecAny(mode os.FileMode) bool {
	return mode&0111 != 0
}

func TypeHandler(input string) {
	for _, name := range Command {
		if name == input {
			fmt.Printf("%s is a shell builtin\n", input)
			return
		}
	}

	pathDirs := strings.SplitSeq(os.Getenv("PATH"), ":")

	for dir := range pathDirs {
		filePath := dir + "/" + input

		fileInfo, err := os.Stat(filePath)

		if err == nil {
			if isExecAny(fileInfo.Mode()) {
				fmt.Printf("%s is %s\n", input, filePath)
				return
			}
		}
	}

	fmt.Printf("%s: not found\n", input)
}

func NotFoundHandler(input string) {
	fmt.Printf("%s: command not found\n", input)
}
