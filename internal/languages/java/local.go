package java

import (
	"log"
	"os/exec"
	"regexp"
	"strings"
)

type LocalJavaInstallation struct {
	Version string
	Build   string
	Date    string
	IsLTS   bool
	Path    string
	Vendor  string
	Active  bool
}

func GetLocalRunningJava() (string, error) {
	cmd := exec.Command("java", "--version")
	std, err := cmd.Output()
	if err != nil {
		// TODO: handle parsing of java version error
		return "", err
	}
	return string(std), nil
}

func GetLocalJavaVersions() []LocalJavaInstallation {
	// get executables of java in computer
	executables, err := GetLocalJavaExecutables()
	if err != nil {
		log.Panicln(err)
	}

	currVersion, err := GetLocalRunningJava()
	if err != nil {
		log.Panicln(err)
	}

	var versions []LocalJavaInstallation

	for _, v := range executables {
		cmd := exec.Command(v, "--version")
		std, err := cmd.Output()
		if err != nil {
			// TODO: handle parsing of java version error
			log.Panic(err)
		}

		output := string(std)
		var version LocalJavaInstallation = parseJavaVersion(output)
		version.Path = v
		version.Active = output == currVersion

		versions = append(versions, version)
	}

	return versions
}

// Function to parse the `java --version` output
func parseJavaVersion(output string) LocalJavaInstallation {
	var info LocalJavaInstallation
	lines := strings.Split(output, "\n")

	// Extracting Version, Build, Date, and Vendor from the output
	reVersion := regexp.MustCompile(`openjdk (\d+\.\d+\.\d+)`)
	reBuild := regexp.MustCompile(`build (\d+\.\d+\.\d+\+[\d\w\-]+)`)
	reDate := regexp.MustCompile(`(\d{4}-\d{2}-\d{2})`)
	reLTS := regexp.MustCompile(`LTS`)

	for _, line := range lines {
		if strings.HasPrefix(line, "openjdk") {
			if match := reVersion.FindStringSubmatch(line); match != nil {
				info.Version = match[1]
			}
		}

		if match := reBuild.FindStringSubmatch(line); match != nil {
			info.Build = match[1]
		}

		if match := reDate.FindStringSubmatch(line); match != nil {
			info.Date = match[1]
		}

		if reLTS.MatchString(line) {
			info.IsLTS = true
		}

		if strings.Contains(line, "Runtime Environment") {
			parts := strings.Fields(line)
			info.Vendor = strings.Join(parts[0:len(parts)-4], " ")
		}
	}

	return info
}

func GetLocalJavaExecutables() ([]string, error) {
	cmd := exec.Command("find", "/usr", "/home", "-path", "/home/root", "-prune", "-o", "-type", "f", "-executable", "-name", "java", "-o", "-name", "binjava")
	std, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	javaLocations := strings.Split(string(std), "\n")
	var finalLocations []string

	for _, v := range javaLocations {
		trimmed := strings.TrimSpace(v)
		if trimmed != "" && strings.Contains(trimmed, "bin/java") {
			finalLocations = append(finalLocations, trimmed)
		}
	}

	return finalLocations, nil
}
