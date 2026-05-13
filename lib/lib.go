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
	found := false
	for _, name := range Command {
		if name == input {
			found = true
			fmt.Printf("%s is a shell builtin\n", input)
			break
		}
	}

	pathDirs := strings.SplitSeq(os.Getenv("PATH"), ":")

	exists := false
	executable := false

	for dir := range pathDirs {
		res := strings.Contains(dir, input)

		if res {
			exists = true
		}

		fileInfo, err := os.Stat(dir)

		if err == nil {
			executable = isExecAny(fileInfo.Mode())
		}

		if exists && executable {
			fmt.Printf("%s is %s\n", input, dir)
			return
		}
	}

	if !found {
		fmt.Printf("%s: not found\n", input)
	}
}

func NotFoundHandler(input string) {
	fmt.Printf("%s: command not found\n", input)
}
