package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("$ ")

		command, err := reader.ReadString('\n')

		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading input:", err)
			os.Exit(1)
		}

		command = strings.TrimSpace(command)

		builtin := map[string]bool{
			"exit": true,
			"echo": true,
			"type": true,
		}

		if command == "exit" {
			break
		} else if strings.HasPrefix(command, "echo ") {
			fmt.Println(command[5:])
		} else if strings.HasPrefix(command, "type ") {
			s := strings.TrimPrefix(command, "type ")
			if _, ok := builtin[s]; ok {
				fmt.Printf("%s is a shell builtin\n", s)
			} else {
				fmt.Println(s + ": not found")

			}
		} else {
			fmt.Println(command + ": command not found")
		}

	}
}
