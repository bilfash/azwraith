package cliintrepeter

import (
	"os/exec"
)

type (
	CliInterpreter interface {
		Execute(command string, args ...string) (string, error)
	}
	cliInterpreter struct {
	}
)

func NewCliInterpreter() CliInterpreter {
	return &cliInterpreter{
	}
}

func (ci *cliInterpreter) Execute(command string, args ...string) (string, error) {
	cmd := exec.Command(command, args...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return err.Error(), err
	}
	return string(out), nil
}
