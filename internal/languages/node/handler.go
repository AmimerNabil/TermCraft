package node

import (
	"bufio"
	"fmt"
	"os/exec"
	"regexp"
	"strings"
)

// ExecuteCommand executes the given command and returns the output as a string
func ExecuteCommand(command []string) (string, error) {
	cmd := exec.Command(command[0], command[1:]...)
	output, err := cmd.CombinedOutput()
	return string(output), err
}

// GetVerboseNodeInfo returns detailed Node.js information formatted for tview
func GetVerboseNodeInfo() string {
	var result strings.Builder

	// Title
	result.WriteString("[yellow]Node.js Detailed Information\n\n")

	// Get Node.js version
	nodeVersionCmd := []string{"node", "--version"}
	nodeVersion, err := ExecuteCommand(nodeVersionCmd)
	if err != nil {
		result.WriteString(fmt.Sprintf("[red]Error retrieving Node.js version: %s\n", err.Error()))
	} else {
		result.WriteString(fmt.Sprintf("[green]Node.js Version: %s\n", strings.TrimSpace(nodeVersion)))
	}

	// Get NPM version
	npmVersionCmd := []string{"npm", "--version"}
	npmVersion, err := ExecuteCommand(npmVersionCmd)
	if err != nil {
		result.WriteString(fmt.Sprintf("[red]Error retrieving NPM version: %s\n", err.Error()))
	} else {
		result.WriteString(fmt.Sprintf("[green]NPM Version: %s\n", strings.TrimSpace(npmVersion)))
	}

	// Get NPM configuration
	npmConfigCmd := []string{"npm", "config", "list"}
	npmConfig, err := ExecuteCommand(npmConfigCmd)
	if err != nil {
		result.WriteString(fmt.Sprintf("[red]Error retrieving NPM config: %s\n", err.Error()))
	} else {
		result.WriteString("\n[white]NPM Configuration:\n")
		result.WriteString(strings.TrimSpace(npmConfig) + "\n")
	}

	// Get Node.js environment variables
	envVarsCmd := []string{"node", "-p", "JSON.stringify(process.env, null, 2)"}
	envVars, err := ExecuteCommand(envVarsCmd)
	if err != nil {
		result.WriteString(fmt.Sprintf("[red]Error retrieving Node.js environment variables: %s\n", err.Error()))
	} else {
		result.WriteString("\n[white]Node.js Environment Variables:\n")
		result.WriteString(strings.TrimSpace(envVars) + "\n")
	}

	return result.String()
}

// GetRemoteNodeVersions fetches and organizes Node.js remote versions as map[major][minor]
func GetRemoteNodeVersions() (map[string]map[string]string, error) {
	// Run the nvm ls-remote command to fetch available Node.js versions
	output, err := ExecuteCommand(OSnvmListNodeVersions)
	if err != nil {
		return nil, fmt.Errorf("error retrieving remote Node.js versions: %v", err)
	}

	// Prepare a regular expression to match version strings like "v16.17.0"
	versionRegex := regexp.MustCompile(`v(\d+)\.(\d+)\.(\d+)`)
	versions := make(map[string]map[string]string)

	// Use a scanner to go through the command output line by line
	scanner := bufio.NewScanner(strings.NewReader(output))
	for scanner.Scan() {
		line := scanner.Text()

		// Find a match using the regex
		match := versionRegex.FindStringSubmatch(line)
		if len(match) == 4 {
			major := match[1]       // Major version
			minor := match[2]       // Minor version
			fullVersion := match[0] // Full version (e.g., v16.17.0)

			// Initialize the map for the major version if it doesn't exist
			if _, exists := versions[major]; !exists {
				versions[major] = make(map[string]string)
			}

			// Add the minor version to the map
			versions[major][minor] = fullVersion
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error scanning Node.js versions: %v", err)
	}

	return versions, nil
}

// InstallNodeVersion installs a specific version of Node.js using nvm
func InstallNodeVersion(version string) error {
	cmd := OSnvmInstallNode(version)
	_, err := ExecuteCommand(cmd["darwin"]) // Adjust this if needed for Linux
	return err
}

// UninstallNodeVersion uninstalls a specific version of Node.js using nvm
func UninstallNodeVersion(version string) error {
	cmd := OSnvmUninstallNode(version)
	_, err := ExecuteCommand(cmd["darwin"]) // Adjust this if needed for Linux
	return err
}

// SetNodeVersion sets a specific Node.js version as the current version using nvm
func SetNodeVersion(version string) error {
	cmd := OSnvmSetNode(version)
	_, err := ExecuteCommand(cmd["darwin"]) // Adjust this if needed for Linux
	return err
}

func ListInstalledNodeVersions() (map[string]string, string, error) {
	output, err := ExecuteCommand([]string{"bash", "-c", "source $HOME/.nvm/nvm.sh && nvm ls"})
	if err != nil {
		return nil, "", fmt.Errorf("error retrieving installed Node.js versions: %v", err)
	}

	installedVersions := make(map[string]string)
	var currentVersion string
	versionRegex := regexp.MustCompile(`->\s*(v\d+\.\d+\.\d+)`) // Regex to find currently active version

	scanner := bufio.NewScanner(strings.NewReader(output))
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "v") { // Only process lines that start with 'v'
			version := strings.Fields(line)[0]
			installedVersions[version] = version

			// Check if the line indicates the currently active version
			if versionRegex.MatchString(line) {
				currentVersion = versionRegex.FindStringSubmatch(line)[1] // Get the version from regex match
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, "", fmt.Errorf("error scanning installed Node.js versions: %v", err)
	}

	return installedVersions, currentVersion, nil
}

// GetCurrentNodeVersion retrieves the currently active Node.js version using nvm
func GetCurrentNodeVersion() (string, error) {
	output, err := ExecuteCommand([]string{"bash", "-c", "source $HOME/.nvm/nvm.sh && nvm current"})
	if err != nil {
		return "", fmt.Errorf("error retrieving current Node.js version: %v", err)
	}
	return strings.TrimSpace(output), nil
}
