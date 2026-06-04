package main

import (
	"fmt"
	"os"
	"os/exec"
	"slices"
	"strings"
)

type Command interface {
	Execute(args []string) error
}

type ExitCommand struct{}

func (c *ExitCommand) Execute(args []string) error {
	os.Exit(0)

	return nil
}

type EchoCommand struct{}

func (c *EchoCommand) Execute(args []string) error {
	if len(args) == 0 {
		return nil
	}

	msg := strings.Join(args, " ")
	fmt.Println(msg)

	return nil
}

type TypeCommand struct {
	builtins []string
}

func (c *TypeCommand) Execute(args []string) error {
	for _, arg := range args {

		if slices.Contains(c.builtins, arg) {
			fmt.Printf("%s is a shell builtin\n", arg)
			continue
		}

		path, err := exec.LookPath(arg)
		if err != nil {
			return fmt.Errorf("%s not found", arg)
		}

		fmt.Printf("%s is %s\n", arg, path)
	}

	return nil
}

type ExternalCommand struct {
	executable string
}

func (c *ExternalCommand) Execute(args []string) error {

	if len(c.executable) == 0 {
		return fmt.Errorf("\n")
	}

	_, err := exec.LookPath(c.executable)

	if err != nil {
		return fmt.Errorf("%s: command not found", c.executable)
	}

	var cmd = exec.Command(c.executable, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

type PwdCommand struct{}

func (c *PwdCommand) Execute(args []string) error {
	path, err := os.Getwd()
	if err != nil {
		return err
	}
	fmt.Println(path)

	return nil
}

type CdCommand struct{}

func (c *CdCommand) Execute(args []string) error {
	var dir string

	if len(args) == 0 {
		dir = os.Getenv("HOME")
	} else {
		dir = args[0]
		dir = strings.ReplaceAll(dir, "~", os.Getenv("HOME"))
	}

	err := os.Chdir(dir)
	if err != nil {
		return fmt.Errorf("cd: %s: No such file or directory", dir)
	}

	return nil
}
