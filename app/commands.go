package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"slices"
	"strings"
)

type Command interface {
	Execute(args []string, out io.Writer) error
}

type ExitCommand struct{}

func (c *ExitCommand) Execute(args []string, _ io.Writer) error {
	os.Exit(0)

	return nil
}

type EchoCommand struct{}

func (c *EchoCommand) Execute(args []string, out io.Writer) error {
	if len(args) == 0 {
		fmt.Fprintln(out)
		return nil
	}
	msg := strings.Join(args, " ")
	_, err := fmt.Fprintln(out, msg)

	return err
}

type TypeCommand struct {
	builtins []string
}

func (c *TypeCommand) Execute(args []string, out io.Writer) error {
	for _, arg := range args {

		if slices.Contains(c.builtins, arg) {
			fmt.Fprintf(out, "%s is a shell builtin\n", arg)
			continue
		}

		path, err := exec.LookPath(arg)
		if err != nil {
			return fmt.Errorf("%s not found", arg)
		}

		fmt.Fprintf(out, "%s is %s\n", arg, path)
	}

	return nil
}

type ExternalCommand struct {
	path string
}

func (c *ExternalCommand) Execute(args []string, out io.Writer) error {

	if len(c.path) == 0 {
		return fmt.Errorf("\n")
	}

	if _, err := exec.LookPath(c.path); err != nil {
		return fmt.Errorf("%s: command not found", c.path)
	}

	cmd := exec.Command(c.path, args...)
	cmd.Stdout = out
	cmd.Stderr = os.Stderr

	err := cmd.Run()

	if err != nil {
		// Suppress status 1
		if _, isExitError := err.(*exec.ExitError); isExitError {
			return nil
		}
	}

	return err
}

type PwdCommand struct{}

func (c *PwdCommand) Execute(args []string, out io.Writer) error {
	path, err := os.Getwd()
	if err != nil {
		return err
	}
	fmt.Fprintln(out, path)

	return nil
}

type CdCommand struct{}

func (c *CdCommand) Execute(args []string, out io.Writer) error {
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
