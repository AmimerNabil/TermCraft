package demo

import (
	"TermCraft/internal/languages/java"
	"TermCraft/internal/languages/python"
	"fmt"
	"strconv"
)

var jlp = java.JavaLangPack{}

func DemoGetJavaVersionsRemote() {
	rv := jlp.GetRemoteVersions()
	for _, java := range rv {
		fmt.Printf("Vendor: %s, Version: %s, Identifier: %s, Installed: %t\n",
			java.JavaVendor, java.JavaVersion, java.Identifier, java.Installed)
	}
}

func DemoGetJavaVersionsLocally() {
	versions := jlp.GetLocalVersions()
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

func DemoGetPyenvLocal() {
	out := python.GetPythonLocal()
	fmt.Println(out)
}

func DemoGetAvailPythonLocal() {
	out := python.GetAvailPythonLocal()
	for i, version := range out {
		fmt.Println(strconv.Itoa(i) + " " + version)
	}
}

func DemoGetAvailablePythonVersionsToInstall() {
	out := python.GetAvailableRemoteVersionsToInstall()
	for k, v := range out {

		fmt.Println(k)
		fmt.Println(v)
	}
}
