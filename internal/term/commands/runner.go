package commands

import (
	"bytes"
	"fmt"
	"os/exec"
)

// TerminalCommand holds the command to be executed and its arguments
type TerminalCommand struct {
	Command string   // The command to run
	Args    []string // The arguments for the command
}

// Run executes the command and returns the output or an error
func (tc *TerminalCommand) Run() (string, string, error) {
	cmd := exec.Command(tc.Command, tc.Args...)
	var out bytes.Buffer
	var errOut bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &errOut

	// Run the command
	err := cmd.Run()
	if err != nil {
		return "", errOut.String(), fmt.Errorf("command execution error: %v, stderr: %s", err, errOut.String())
	}

	// Return the command output
	return out.String(), errOut.String(), nil
}
