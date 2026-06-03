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

	builtinNames := []string{"echo", "exit", "type", "pwd", "cd"}
	commands["type"] = &TypeCommand{
		builtins: builtinNames,
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

		parts := strings.Fields(input)
		cmdName := parts[0]
		args := parts[1:]

		if command, exists := commands[cmdName]; exists {
			command.Execute(args)
		} else {
			externalCmd := &ExternalCommand{}
			externalCmd.Execute(parts)
		}

	}

}
