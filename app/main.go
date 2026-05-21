package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"slices"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("$ ")

		line, err := reader.ReadString('\n')

		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading input:", err)
			os.Exit(1)
		}
		cmdLine := strings.TrimSpace(line)

		words := strings.Split(cmdLine, " ")

		cmd := words[0]
		args := words[1:]

		switch cmd {

		case "exit":
			handleExit()
		case "echo":
			handleEcho(args)
		case "type":
			handleType(args)
		default:
			fmt.Println(cmd + ": not found")
		}

	}
}

func handleExit() {
	os.Exit(0)
}

func handleEcho(args []string) {
	msg := strings.Join(args, " ")

	fmt.Println(msg)
}

func handleType(args []string) {
	builtins := []string{"exit", "echo", "type"}

	for _, arg := range args {

		if slices.Contains(builtins, arg) {
			fmt.Printf("%s is a shell builtin\n", arg)
			continue
		}

		path, err := exec.LookPath(arg)
		if err != nil {
			fmt.Printf("%s not found\n", arg)
			continue
		}

		fmt.Printf("%s is %s\n", arg, path)

	}
}
