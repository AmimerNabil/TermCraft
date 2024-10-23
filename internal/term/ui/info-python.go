package ui

import (
	"TermCraft/internal/languages/python"
	"fmt"
	"log"
	"os/exec"
	"regexp"
	"strings"

	"github.com/gdamore/tcell/v2"
)

type PythonPanel struct {
	*Panel

	remotePythons map[string]map[string][]string

	currGlobal     string
	currLocal      string
	currPythonInfo string

	localVersions []string
}

func NewPythonPanel() *PythonPanel {
	pp := &PythonPanel{
		Panel: &Panel{},
	}

	pp.i()
	pp.init()

	return pp
}

func (pp *PythonPanel) init() {
	// format data before constructing:

	// 1st: get local and global
	pp.currLocal = python.GetPyenvLocal()
	pp.currGlobal = python.GetPyenvGlobal()

	// 2nd: color info
	info := formatPythonInfo(python.GetPythonLocal())
	pp.currPythonInfo = info

	// 3rd: place local and global on list
	pp.localVersions = markVersions(python.GetAvailPythonLocal(), strings.TrimSpace(pp.currGlobal), strings.TrimSpace(pp.currLocal))

	pp.remotePythons = python.GetAvailableRemoteVersionsToInstall()

	pp.initCurrVersionInfo(pp.currPythonInfo)
	pp.initCurrVersions(pp.localVersions)
	pp.initRemoteVersions(pp.remotePythons, python.InstallPythonVersion)

	// set input captures
	pp.currVersionsInstalled.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		index := pp.currVersionsInstalled.GetCurrentItem()
		var text string
		if len(pp.localVersions) > 0 && index >= 0 {
			itemText, _ := pp.currVersionsInstalled.GetItemText(index)
			text = itemText
		}

		cleanText := CleanVersionString(text)

		switch event.Key() {
		case tcell.KeyTab:
			App.SetFocus(pp.remoteVersionsAvailable)
			return nil
		case tcell.KeyEscape:
			App.SetFocus(AvailableLanguesSections.El)
		case tcell.KeyRune:
			switch event.Rune() {
			case 'G':
				if !strings.Contains(text, "global") && !strings.Contains(text, "system") {
					_, err := exec.Command("pyenv", "global", cleanText).Output()
					if err != nil {
						log.Fatalf("no global vercion %v", err)
					}

					pp.init()
					App.SetFocus(pp.currVersionsInstalled)
				}
			case 'L':
				if !strings.Contains(text, "using") && !strings.Contains(text, "system") {
					_, err := exec.Command("pyenv", "local", cleanText).Output()
					if err != nil {
						log.Fatalf("no local vercion %v", err)
					}

					pp.init()
					App.SetFocus(pp.currVersionsInstalled)
				}
			case 'D':
				if !strings.Contains(text, "using") && !strings.Contains(text, "system") {
					pp.UninstallPythonVersion(python.UnInstallPythonVersion, CleanVersionString(text), index, pp.currVersionsInstalled)
				} else {
					setConfirmationContent("Can't remove this python version. Press enter to go back.",
						func() {
							App.SetFocus(pp.currVersionsInstalled)
						}, func() {
							App.SetFocus(pp.currVersionsInstalled)
						})
				}
			}
		}

		return event
	})
}

func formatPythonInfo(input string) string {
	var result strings.Builder

	// Define colors (tview syntax)
	headerColor := "[yellow]"
	normalColor := "[white]"
	resetColor := "[-]"

	// Split the input into lines
	lines := strings.Split(input, "\n")

	// Track which section we are processing
	for _, line := range lines {
		line = strings.TrimSpace(line)

		switch {
		case strings.HasPrefix(line, "Python"):
			result.WriteString(fmt.Sprintf("%sPython Version:%s %s%s\n", headerColor, resetColor, normalColor, line))

		case strings.HasPrefix(line, "'main'"):
			result.WriteString(fmt.Sprintf("%sBuild Info:%s %s%s\n", headerColor, resetColor, normalColor, line))

		case strings.HasPrefix(line, "GCC"):
			result.WriteString(fmt.Sprintf("%sGCC Version:%s %s%s\n", headerColor, resetColor, normalColor, line))

		case strings.HasPrefix(line, "Version Info:"):
			result.WriteString(fmt.Sprintf("%sVersion Info:%s %s%s\n", headerColor, resetColor, normalColor, line))

		case strings.HasPrefix(line, "System:"):
			result.WriteString(fmt.Sprintf("%sSystem:%s %s%s\n", headerColor, resetColor, normalColor, line))

		case strings.HasPrefix(line, "Release:"):
			result.WriteString(fmt.Sprintf("%sRelease:%s %s%s\n", headerColor, resetColor, normalColor, line))

		case strings.HasPrefix(line, "Processor:"):
			result.WriteString(fmt.Sprintf("%sProcessor:%s %s%s\n", headerColor, resetColor, normalColor, line))

		case strings.HasPrefix(line, "pip"):
			result.WriteString(fmt.Sprintf("%sPip Version:%s %s%s\n", headerColor, resetColor, normalColor, line))

		default:
			// If the line doesn't match a known section, just skip or add it as plain text.
		}
	}

	return result.String()
}

func markVersions(versions []string, globalVersion string, localVersion string) []string {
	var result []string

	// Loop over the list of versions and append the appropriate marker
	for _, version := range versions {
		v := strings.ReplaceAll(version, "*", "")
		v = strings.TrimSpace(v)

		if v == globalVersion && v == localVersion {
			result = append(result, fmt.Sprintf("%s (global) -> using", version))
		} else if v == globalVersion {
			result = append(result, fmt.Sprintf("%s (global)", version))
		} else if v == localVersion {
			result = append(result, fmt.Sprintf("%s -> using", version))
		} else {
			result = append(result, version) // Unmarked versions
		}
	}

	return result
}

func CleanVersionString(version string) string {
	cleaned := strings.TrimSpace(version)
	cleaned = strings.TrimPrefix(cleaned, "*")
	re := regexp.MustCompile(`\s*\(.*\)$`)
	cleaned = re.ReplaceAllString(cleaned, "")

	return strings.TrimSpace(cleaned)
}
