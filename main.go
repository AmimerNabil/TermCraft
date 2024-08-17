package main

import (
	"TermCraft/internal/languages/java"
	"fmt"
)

func main() {
	versions := java.GetLocalJavaVersions()
	for _, info := range versions {
		output := fmt.Sprintf(
			"Java Version Info:\n"+
				"  Active: %v\n"+
				"  Version: %s\n"+
				"  Build: %s\n"+
				"  Date: %s\n"+
				"  Is LTS: %t\n"+
				"  Path: %s\n"+
				"  Vendor: %s\n",
			info.Active, info.Version, info.Build, info.Date, info.IsLTS, info.Path, info.Vendor,
		)
		fmt.Println(output)
	}
}
