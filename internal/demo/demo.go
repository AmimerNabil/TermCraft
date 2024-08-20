package demo

import (
	"TermCraft/internal/languages/java"
	"fmt"
)

func DemoGetJavaVersionsRemote() {
	rv, err := java.GetRemoteVersions()
	if err != nil {
	}

	for _, java := range rv {
		fmt.Printf("Vendor: %s, Version: %s, Identifier: %s, Installed: %t\n",
			java.JavaVendor, java.JavaVersion, java.Identifier, java.Installed)
	}
}

func DemoGetJavaVersionsLocally() {
	versions := java.GetLocalJavaVersions()
	for _, info := range versions {
		output := fmt.Sprintf(
			"Java Version Info:\n"+
				"  Active: %v\n"+
				"  Version: %s\n"+
				"  Date: %s\n",
			info.CurrentlyActive, info.JavaVersion, info.JavaVersionDate)
		fmt.Println(output)
	}
}
