package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type Shell struct {
	in       *bufio.Reader
	out      io.Writer
	commands map[string]Command
}

func (s *Shell) run() {
	for {
		fmt.Print("$ ")

		input, err := s.in.ReadString('\n')
		if err != nil {
			break
		}

		input = strings.TrimSpace(input)
		if input == "" {
			continue
		}

		parts, err := ParseQuotations(input)
		if err != nil {
			fmt.Fprintln(s.out, "Error parsing quotes:", err)
			continue
		}

		args, target, err := ParseRedirection(parts)
		if err != nil {
			fmt.Fprintln(s.out, "Error parsing redirection:", err)
			continue
		}

		cmdOut := s.out
		var outFile *os.File

		if target != "" {
			// O_CREATE: Create it if it doesn't exist
			// O_WRONLY: Open for writing only
			// O_TRUNC: Truncate (empty) the file if it already exists (standard '>' behavior)
			// 0644: Standard permissions (read/write for owner, read for others)
			outFile, err = os.OpenFile(target, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
			if err != nil {
				fmt.Fprintln(s.out, "Error opening redirect file:", err)
				continue
			}

			cmdOut = outFile
		}

		if command, exists := s.commands[args[0]]; exists {
			err = command.Execute(args[1:], cmdOut)
		} else {
			ext := &ExternalCommand{
				path: args[0],
			}
			err = ext.Execute(args[1:], cmdOut)
		}

		if err != nil {
			fmt.Fprintln(s.out, err)
		}

		if outFile != nil {
			outFile.Close()
		}
	}
}

func NewShell(in *os.File, out *os.File) *Shell {
	return &Shell{
		in:  bufio.NewReader(in),
		out: out,
		commands: map[string]Command{
			"exit": &ExitCommand{},
			"echo": &EchoCommand{},
			"pwd":  &PwdCommand{},
			"cd":   &CdCommand{},
			"type": &TypeCommand{
				builtins: []string{"echo", "exit", "type", "pwd", "cd"},
			},
		},
	}
}

func main() {
	kjell := NewShell(os.Stdin, os.Stdout)
	kjell.run()
}
