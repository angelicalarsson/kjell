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

		parts, err := parseInput(input)

		if err != nil {
			fmt.Printf("Error with parseInput: %v", err)
		}

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

func parseInput(input string) ([]string, error) {
	inSingleQuote := false
	var parts []string
	var builder strings.Builder

	for _, r := range input {

		if r == '\'' {
			inSingleQuote = !inSingleQuote
			continue
		}

		if r == ' ' && !inSingleQuote {
			if builder.Len() == 0 {
				continue
			}
			parts = append(parts, builder.String())
			builder.Reset()
			continue
		}

		builder.WriteRune(r)
	}
	parts = append(parts, builder.String())

	return parts, nil
}
