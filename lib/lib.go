package lib

import (
	"errors"
	"fmt"
	"iter"
	"os"
	"os/exec"
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

func getSliceOfPathDirs() iter.Seq[string] {
	return strings.SplitSeq(os.Getenv("PATH"), ":")
}

func isShellBuiltIn(input string) bool {
	for _, name := range Command {
		return name == input
	}

	return false
}

func getFileData(input string) (os.FileInfo, string, error) {
	pathDirs := getSliceOfPathDirs()

	for dir := range pathDirs {
		filePath := dir + "/" + input
		fileInfo, err := os.Stat(filePath)
		if err == nil {
			return fileInfo, filePath, nil
		}
	}

	return nil, "", errors.New("failed to get file data")
}

func TypeHandler(input string) {
	if isShellBuiltIn(input) {
		fmt.Printf("%s is a shell builtin\n", input)
		return
	}

	fileInfo, filePath, err := getFileData(input)

	if err == nil && isExecAny(fileInfo.Mode()) {
		fmt.Printf("%s is %s\n", input, filePath)
		return
	}

	fmt.Printf("%s: not found\n", input)
}

func DefaultHandler(input string, args []string) {
	fileInfo, filePath, err := getFileData(input)

	if err == nil && isExecAny(fileInfo.Mode()) {
		cmd := exec.Command(filePath, args...)

		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Stdin = os.Stdin

		if err := cmd.Run(); err != nil {
			fmt.Printf("error occured")
		}
		return
	}

	fmt.Printf("%s: command not found\n", input)
}
