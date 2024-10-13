package python

import (
	"TermCraft/internal/term/commands"
	"fmt"
	"regexp"
	"strings"
)

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
