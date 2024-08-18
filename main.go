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
				"  Date: %s\n",
			info.CurrentlyActive, info.JavaVersion, info.JavaVersionDate)
		fmt.Println(output)
	}
}
