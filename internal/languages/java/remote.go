// REMOTE java is responsible for handling java online versions and installing them
// properly. It's job is to fetch, parse and return the available versions

package java

import (
	"TermCraft/internal/term/commands"
	"fmt"
	"os"
)

func InstallSdkMan() (string, error) {
	command := commands.TerminalCommand{
		Command: OSsdkInstall[0],
		Args:    OSsdkInstall[1:],
	}

	script, _, err := command.Run() // what I get here is the sdkman script
	if err != nil {
		fmt.Println("problem fetching script", err)
		return "", err
	}

	command = commands.TerminalCommand{
		Command: "bash",
		Args: []string{
			"-c", script,
		},
	}

	_, _, errr := command.Run()
	if errr != nil {
		fmt.Println("problem exec script", errr)
		return "", errr
	}

	return "", nil
}

func isSDKMANInstalled() bool {
	// Check for the SDKMAN directory
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error getting user home directory:", err)
		return false
	}

	sdkmanDir := homeDir + "/.sdkman"
	if _, err := os.Stat(sdkmanDir); !os.IsNotExist(err) {
		return true
	}

	// Check if the `sdk` command is available

	command := commands.TerminalCommand{
		Command: OSsdkmanVersion[0],
		Args:    OSsdkmanVersion[1:],
	}

	_, _, error := command.Run()
	return error == nil
}
