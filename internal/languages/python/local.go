package python

import (
	"TermCraft/internal/term/commands"
	"TermCraft/internal/utils"
	"strings"
)

func GetPythonLocal() string {
	command := commands.TerminalCommand{
		Command: OSpyenvLocal[0], Args: OSpyenvLocal[1:],
	}

	stdOut, stdErr, _ := command.Run()

	if stdErr != "" {
		stdOut = "No python versions activated in this directory."
	}

	return stdOut
}

func GetAvailPythonLocal() []string {
	command := commands.TerminalCommand{
		Command: OSpyenvListPython[0], Args: OSpyenvListPython[1:],
	}

	stdOut, stdErr, _ := command.Run()

	if stdErr != "" {
		stdOut = "No python versions activated in this directory."
	}

	output := utils.Filter(strings.Split(stdOut, "\n"), func(s string) bool {
		return s != "" && s != "system" && s != "System"
	})

	return output
}
