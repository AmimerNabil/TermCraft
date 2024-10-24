package node

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

func ExecuteCommand(command []string) (string, error) {
	cmd := exec.Command(command[0], command[1:]...)
	output, err := cmd.CombinedOutput()
	return string(output), err
}

func GetVerboseNodeInfo() string {
	var result strings.Builder

	nodeVersionCmd := []string{"fnm", "current"}
	nodeVersion, err := ExecuteCommand(nodeVersionCmd)
	if err != nil {
		result.WriteString(fmt.Sprintf("[red]Error retrieving Node.js version: %s\n", err.Error()))
	} else {
		result.WriteString(fmt.Sprintf("[green]Node.js Version: %s\n", strings.TrimSpace(nodeVersion)))
	}

	npmVersionCmd := []string{"npm", "--version"}
	npmVersion, err := ExecuteCommand(npmVersionCmd)
	if err != nil {
		result.WriteString(fmt.Sprintf("[red]Error retrieving NPM version: %s\n", err.Error()))
	} else {
		result.WriteString(fmt.Sprintf("[green]NPM Version: %s", strings.TrimSpace(npmVersion)))
	}

	npmConfigCmd := []string{"npm", "config", "list"}
	npmConfig, err := ExecuteCommand(npmConfigCmd)
	if err != nil {
		result.WriteString(fmt.Sprintf("[red]Error retrieving NPM config: %s\n", err.Error()))
	} else {
		result.WriteString("\n[lightyellow]NPM Configuration:[white]\n")

		lines := strings.Split(npmConfig, "\n")

		for _, line := range lines {
			parts := strings.SplitN(line, "=", 2)

			if len(parts) == 2 {
				title := strings.TrimSpace(parts[0]) // Things before "="
				value := strings.TrimSpace(parts[1]) // Things after "="

				result.WriteString(fmt.Sprintf("[blue]%s [white]= %s\n", title, value))
			}
		}
		result.WriteString("\n")
	}

	return result.String()
}

func GetRemoteNodeVersions() (map[string]map[string][]string, error) {
	output, err := ExecuteCommand([]string{"fnm", "ls-remote"})
	if err != nil {
		return nil, fmt.Errorf("error retrieving remote Node.js versions: %v", err)
	}

	versionRegex := regexp.MustCompile(`v(\d+)\.(\d+)\.(\d+)`)
	versions := make(map[string]map[string][]string)

	scanner := bufio.NewScanner(strings.NewReader(output))
	for scanner.Scan() {
		line := scanner.Text()

		match := versionRegex.FindStringSubmatch(line)
		if len(match) == 4 {
			major := match[1]       // Major version
			minor := match[2]       // Minor version
			fullVersion := match[0] // Full version (e.g., v16.17.0)

			if _, exists := versions[major]; !exists {
				versions[major] = make(map[string][]string)
			}

			versions[major][minor] = append(versions[major][minor], fullVersion)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error scanning Node.js versions: %v", err)
	}

	return versions, nil
}

func InstallNodeVersion(version string) (string, string, error) {
	cmd := []string{"fnm", "install", version}
	output, err := ExecuteCommand(cmd)
	return output, "", err
}

func UninstallNodeVersion(version string) (string, string, error) {
	cmd := []string{"fnm", "uninstall", version}
	output, err := ExecuteCommand(cmd)
	return output, "", err
}

// Set the global Node.js version for all directories using `fnm default`
func SetNodeVersion(version string) error {
	cmd := []string{"fnm", "use", version}      // Changed from 'use' to 'default'
	cmd2 := []string{"fnm", "default", version} // Changed from 'use' to 'default'

	_, err := ExecuteCommand(cmd)
	ExecuteCommand(cmd2)

	return err
}

// Set the local Node.js version for a specific directory by creating/modifying a .node-version file
func SetLocalNodeVersion(version string) error {
	// Check if .node-version file exists
	fileName := ".node-version"
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return fmt.Errorf("error opening or creating .node-version file: %v", err)
	}
	defer file.Close()

	// Write the version to the file, overwriting existing content
	_, err = file.WriteString(version + "\n")
	if err != nil {
		return fmt.Errorf("error writing version to .node-version file: %v", err)
	}

	fmt.Printf("Local Node.js version set to %s in .node-version file.\n", version)
	return nil
}

// List installed Node.js versions and identify the current one based on `node -v`
func ListInstalledNodeVersions() ([]string, string, error) {
	output, err := ExecuteCommand([]string{"fnm", "ls"})
	if err != nil {
		return nil, "", fmt.Errorf("error retrieving installed Node.js versions: %v", err)
	}

	// Get the currently active version via `node -v`
	currentNodeVersion, err := GetCurrentNodeVersion()
	if err != nil {
		return nil, "", err
	}

	var installedVersions []string
	versionRegex := regexp.MustCompile(`v\d+\.\d+\.\d+`) // Regex to find version numbers

	scanner := bufio.NewScanner(strings.NewReader(output))
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		// Match Node.js version numbers in the list
		if versionRegex.MatchString(line) {
			version := versionRegex.FindString(line)
			versionEntry := version

			// Check if the version is the current version
			if version == currentNodeVersion {
				versionEntry += " -> using"
			}

			// Check if the version is the default version
			if strings.Contains(line, "default") {
				versionEntry += " (default)"
			}

			installedVersions = append(installedVersions, versionEntry)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, "", fmt.Errorf("error scanning installed Node.js versions: %v", err)
	}

	return installedVersions, currentNodeVersion, nil
}

func GetCurrentNodeVersion() (string, error) {
	output, err := ExecuteCommand([]string{"fnm", "current"})
	if err != nil {
		return "", fmt.Errorf("error retrieving current Node.js version: %v", err)
	}
	return strings.TrimSpace(output), nil
}
