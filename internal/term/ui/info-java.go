package ui

import (
	"TermCraft/internal/languages/java"
	commandtext "TermCraft/internal/term/ui/command-text"
	"fmt"
	"strings"

	"github.com/gdamore/tcell/v2"
)

type JavaPanel struct {
	*Panel

	currJavaInfo string
}

func NewJavaPanel() *JavaPanel {
	jp := &JavaPanel{
		Panel: &Panel{},
	}

	jp.i(commandtext.JavaPanel)
	jp.init()

	return jp
}

func (jp *JavaPanel) init() {
	// #1 : get the info versions and set them up
	javaProperties := java.GetAllJavaVersionInformation("java")
	info := []string{
		fmt.Sprintf("[yellow]Java Home:[-] %s", javaProperties.JavaHome),
		fmt.Sprintf("[yellow]Runtime Name:[-] %s", javaProperties.JavaRuntimeName),
		fmt.Sprintf("[yellow]Java Version:[-] %s", javaProperties.JavaVersion),
		fmt.Sprintf("[yellow]Vendor:[-] %s", javaProperties.JavaVendor),
		fmt.Sprintf("[yellow]VM Name:[-] %s", javaProperties.JavaVMName),
		fmt.Sprintf("[yellow]VM Version:[-] %s", javaProperties.JavaVMVersion),
		fmt.Sprintf("[yellow]OS Architecture:[-] %s", javaProperties.OSArch),
		fmt.Sprintf("[yellow]OS Name:[-] %s", javaProperties.OSName),
		fmt.Sprintf("[yellow]OS Version:[-] %s", javaProperties.OSVersion),
		fmt.Sprintf("[yellow]User Name:[-] %s", javaProperties.UserName),
	}

	jp.currJavaInfo = strings.Join(info, "\n")
	jp.initCurrVersionInfo(jp.currJavaInfo)

	// #2: get localjava properties
	localProperties := java.GetLocalJavaVersionsSdk()

	localPropertiesStrings := []string{}
	for _, java := range localProperties {
		var inUse string

		if java.InUse {
			inUse = "-> using"
		} else {
			inUse = ""
		}

		text := fmt.Sprintf("(%s)\t id: %s %s", java.JavaVendor, java.Identifier, inUse)
		localPropertiesStrings = append(localPropertiesStrings, text)
	}

	jp.localVersions = localPropertiesStrings
	jp.initCurrVersions(jp.localVersions)

	// #3: setup remote versions
	remotes := java.GetRemoteVersions()
	vendorMap := make(map[string]map[string][]string)
	for _, j := range remotes {
		versionParts := strings.Split(j.JavaVersion, ".")
		if len(versionParts) < 2 {
			continue
		}

		majorMinor := fmt.Sprintf("%s.%s", versionParts[0], versionParts[1])

		if _, ok := vendorMap[j.JavaVendor]; !ok {
			vendorMap[j.JavaVendor] = make(map[string][]string)
		}

		var installed string
		if j.Installed {
			installed = "*"
		} else {
			installed = ""
		}
		vendorMap[j.JavaVendor][majorMinor] = append(vendorMap[j.JavaVendor][majorMinor], j.Identifier+" "+installed)
	}

	jp.initRemoteVersions(jp.updateLocal, vendorMap, java.InstallJavaVersion)

	jp.currVersionsInstalled.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		index := jp.currVersionsInstalled.GetCurrentItem()
		var text string

		if len(jp.localVersions) > 0 && index >= 0 {
			itemText, _ := jp.currVersionsInstalled.GetItemText(index)
			text = itemText
		}

		cleanText := javacleanVersionString(text)

		switch event.Key() {
		case tcell.KeyTab:
			App.SetFocus(jp.remoteVersionsAvailable)
			return nil
		case tcell.KeyEscape:
			App.SetFocus(AvailableLanguesSections.El)
		case tcell.KeyEnter:
			java.SetJavaVersion(cleanText)
			jp.init()
			App.SetFocus(jp.currVersionsInstalled)
		case tcell.KeyRune:
			switch event.Rune() {
			case 'D':
				if !strings.Contains(text, "using") && !strings.Contains(text, "system") {
					jp.UninstallPythonVersion(java.UnInstallJavaVersion, cleanText, index, jp.currVersionsInstalled)
					jp.init()
					App.SetFocus(jp.currVersionsInstalled)
				} else {
					setConfirmationContent("Can't remove this python version. Press enter to go back.",
						func() {
							App.SetFocus(jp.currVersionsInstalled)
						}, func() {
							App.SetFocus(jp.currVersionsInstalled)
						})
				}
			}
		}

		return event
	})
}

func (jp *JavaPanel) updateLocal() []string {
	localProperties := java.GetLocalJavaVersionsSdk()

	localPropertiesStrings := []string{}
	for _, java := range localProperties {
		var inUse string

		if java.InUse {
			inUse = "-> using"
		} else {
			inUse = ""
		}

		text := fmt.Sprintf("(%s)\t id: %s %s", java.JavaVendor, java.Identifier, inUse)
		localPropertiesStrings = append(localPropertiesStrings, text)
	}

	return localPropertiesStrings
}

func javacleanVersionString(version string) string {
	if idx := strings.Index(version, "id:"); idx != -1 {
		idPart := strings.TrimSpace(version[idx+3:])
		idPart = strings.ReplaceAll(idPart, "-> using", "")
		return strings.TrimSpace(idPart)
	}
	return version
}
