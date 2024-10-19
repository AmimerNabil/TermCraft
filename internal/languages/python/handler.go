package python

import (
	"TermCraft/internal/term/commands"
	"TermCraft/internal/utils"
	"fmt"
	"log"
	"os/exec"
	"regexp"
	"runtime"
	"strings"
)

func GetPyenvLocal() string {
	pythonVersion, err := exec.Command("pyenv", "local").Output()
	if err != nil {
		return "no local vercions handled with pyenv"
	}

	return strings.TrimSpace(string(pythonVersion))
}

func GetPyenvGlobal() string {
	pythonVersion, err := exec.Command("pyenv", "global").Output()
	if err != nil {
		return "no global vercions handled with pyenv"
	}

	return strings.TrimSpace(string(pythonVersion))
}

func GetPythonLocal() string {
	pythonVersion, err := exec.Command("python", "--version").Output()
	if err != nil {
		log.Fatalf("Error getting Python version: %v", err)
	}

	pythonBuild, err := exec.Command("python", "-c", `import platform; print(platform.python_build())`).Output()
	if err != nil {
		log.Fatalf("Error getting Python build: %v", err)
	}

	pythonCompiler, err := exec.Command("python", "-c", `import platform; print(platform.python_compiler())`).Output()
	if err != nil {
		log.Fatalf("Error getting Python compiler: %v", err)
	}

	pythonFullVersion, err := exec.Command("python", "-c", `import sys, platform; print(f"Version Info: {sys.version_info}\nSystem: {platform.system()}\nRelease: {platform.release()}\nProcessor: {platform.processor()}")`).Output()
	if err != nil {
		log.Fatalf("Error getting full Python version details: %v", err)
	}

	pipVersion, err := exec.Command("pip", "--version").Output()
	if err != nil {
		log.Fatalf("Error getting full pip version details: %v", err)
	}

	return string(pythonVersion) + string(pythonBuild) + string(pythonCompiler) + string(pythonFullVersion) + string(pipVersion)
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

func GetAvailableRemoteVersionsToInstall() map[string]map[string][]string {
	command := commands.TerminalCommand{
		Command: OSpyenvInstallList[0], Args: OSpyenvInstallList[1:],
	}

	stdOut, stdErr, _ := command.Run()

	if stdErr != "" {
		stdOut = "No python versions activated in this directory."
	}

	return categorizeVersions(strings.Split(stdOut, "\n"))
}

func UnInstallPythonVersion(identifier string) (string, string, error) {
	cm := OSpyenvUninstallPython(identifier)
	command := commands.TerminalCommand{
		Command: cm[runtime.GOOS][0], Args: cm[runtime.GOOS][1:],
	}

	stdOut, stderr, err := command.Run()

	if stderr != "" || err != nil {
		return "", stderr, err
	}

	return stdOut, "", nil
}

func InstallPythonVersion(identifier string) (string, string, error) {
	cm := OSpyenvInstallPython(identifier)
	command := commands.TerminalCommand{
		Command: cm[runtime.GOOS][0], Args: cm[runtime.GOOS][1:],
	}

	stdOut, stderr, err := command.Run()

	if stderr != "" || err != nil {
		return "", stderr, err
	}

	return stdOut, "", nil
}

func isVersionNumber(version string) bool {
	re := regexp.MustCompile(`^\d+\.\d+\.\d+$`)
	return re.MatchString(version)
}

func categorizeVersions(versions []string) map[string]map[string][]string {
	// Create a map to hold categories and their corresponding versions
	categories := make(map[string]map[string][]string)

	// Iterate over the list of versions
	for _, version := range versions {
		// Trim whitespace from each version
		version = strings.TrimSpace(version)
		if version == "" {
			continue
		}

		// Determine the category based on prefix
		var category string
		switch {
		case strings.HasSuffix(version, "latest"):
			category = "latest"
		case strings.HasPrefix(version, "miniconda2"):
			category = "micropython2"
		case strings.HasPrefix(version, "miniconda3"):
			category = "micropython3"
		case strings.HasPrefix(version, "micropython"):
			category = "micropython"
		case strings.HasPrefix(version, "mambaforge"):
			category = "mambaforge"
		case strings.HasPrefix(version, "graalpy"):
			category = "graalpy"
		case strings.HasPrefix(version, "stackless"):
			category = "stackless"
		case strings.HasPrefix(version, "anaconda3"):
			category = "anaconda3"
		case strings.HasPrefix(version, "anaconda2"):
			category = "anaconda2"
		case strings.HasPrefix(version, "anaconda"):
			category = "anaconda"
		case strings.HasPrefix(version, "pypy2"):
			category = "pypy2"
		case strings.HasPrefix(version, "pypy3"):
			category = "pypy3"
		case strings.HasPrefix(version, "pypy"):
			category = "pypy"
		case strings.HasPrefix(version, "miniforge"):
			category = "miniforge"
		case strings.HasPrefix(version, "nogil"):
			category = "nogil"
		case isVersionNumber(version):
			category = "pyston"
		default:
			category = "others" // Place unknown versions here
		}

		// Initialize the category map if it doesn't exist
		if categories[category] == nil {
			categories[category] = make(map[string][]string)
		}

		// Extract the version number
		verNum := extractVersionNumber(version)

		// Determine the subcategory based on major and minor version
		subcategory := getVersionSubcategory(verNum)

		// Append the version to the corresponding category and subcategory
		categories[category][subcategory] = append(categories[category][subcategory], version)
	}

	return categories
}

// extractVersionNumber extracts the version number part from a string like "stackless-2.7.2"
func extractVersionNumber(version string) string {
	// Extract version number using regex
	re := regexp.MustCompile(`\d+\.\d+(\.\d+)?`)
	match := re.FindString(version)
	return match
}

// getVersionSubcategory returns the subcategory based on major and minor version levels
func getVersionSubcategory(version string) string {
	// Split the version string
	parts := strings.Split(version, ".")
	if len(parts) >= 2 {
		major := parts[0]
		minor := parts[1]

		// Create subcategory like "2.x", "2.7.x"
		if len(parts) == 2 {
			return fmt.Sprintf("%s.x", major)
		} else if len(parts) == 3 {
			return fmt.Sprintf("%s.%s.x", major, minor)
		}
	}
	return "unknown"
}
