package ui

import (
	"TermCraft/internal/languages/java"
	"fmt"
	"strings"
)

type JavaPanel struct {
	*Panel

	currJavaInfo string
}

func NewJavaPanel() *JavaPanel {
	jp := &JavaPanel{
		Panel: &Panel{},
	}

	jp.i()
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
}
