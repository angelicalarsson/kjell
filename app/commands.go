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
			fmt.Printf("%s not found\n", arg)
			continue
		}

		fmt.Printf("%s is %s\n", arg, path)
	}

	return nil
}

type ExternalCommand struct{}

func (c *ExternalCommand) Execute(args []string) error {
	targetDir := args[0]

	_, err := exec.LookPath(args[0])

	if err != nil {
		return fmt.Errorf("%s: command not found\n", targetDir)
	}

	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()

	return nil
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
	targetDir := args[0]

	targetDir = strings.ReplaceAll(targetDir, "~", os.Getenv("HOME"))

	err := os.Chdir(targetDir)
	if err != nil {
		fmt.Printf("cd: %s: No such file or directory\n", targetDir)
	}

	return nil
}
