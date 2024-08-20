// REMOTE java is responsible for handling java online versions and installing them
// properly. It's job is to fetch, parse and return the available versions

package java

import (
	"TermCraft/internal/term/commands"
	"fmt"
	"os"
	"regexp"
	"strings"
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

	return "successfully Installed", nil
}

func GetRemoteVersions() ([]RemoteJavaProperties, error) {
	var rv []RemoteJavaProperties

	command := commands.TerminalCommand{
		Command: OSsdkListJava[0],
		Args:    OSsdkListJava[1:],
	}

	versionsSTDO, _, err := command.Run()
	if err != nil {
		fmt.Println("problem fetching versions", err)
		return nil, err
	}

	parseJavaOutput(versionsSTDO, &rv)

	return rv, nil
}

func IsSDKMANInstalled() bool {
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

func parseJavaOutput(output string, rv *[]RemoteJavaProperties) {
	lines := strings.Split(output, "\n")
	var currentVendor string

	// Regular expression to match Java version lines with or without vendor
	re := regexp.MustCompile(`^\s*([A-Za-z]+)?\s*\|\s*(>>>)?\s*\|\s+(\S+)\s+\|\s+(\S+)\s+\|\s*(installed|local only|)\s*\|\s+(\S+)`)

	for _, line := range lines {
		matches := re.FindStringSubmatch(line)
		if matches != nil {
			vendor := matches[1]
			if vendor == "" {
				vendor = currentVendor
			} else {
				currentVendor = vendor
			}

			version := matches[3]
			status := matches[5]
			identifier := matches[6]

			// Determine if the version is installed
			installed := matches[2] == ">>>" || status == "installed" || status == "local only"

			javaProperties := RemoteJavaProperties{
				JavaVendor:  vendor,
				JavaVersion: version,
				Identifier:  identifier,
				Installed:   installed,
			}

			*rv = append(*rv, javaProperties)
		}
	}
}
