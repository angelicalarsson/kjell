package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	commands := make(map[string]Command)

	commands["exit"] = &ExitCommand{}
	commands["echo"] = &EchoCommand{}

	commands["type"] = &TypeCommand{
		builtins: []string{"echo", "exit", "type", "pwd", "cd"},
	}
	commands["pwd"] = &PwdCommand{}
	commands["cd"] = &CdCommand{}

	for {
		fmt.Print("$ ")

		input, err := reader.ReadString('\n')
		if err != nil {
			break
		}

		input = strings.TrimSpace(input)
		if input == "" {
			continue
		}

		parts, err := ParseInput(input)

		if err != nil {
			fmt.Print(err)
		}

		cmd := parts[0]
		args := parts[1:]

		if command, exists := commands[cmd]; exists {
			err = command.Execute(args)
		} else {
			ext := &ExternalCommand{
				executable: cmd,
			}
			err = ext.Execute(args)
		}

		if err != nil {
			fmt.Println(err)
		}

	}

}
