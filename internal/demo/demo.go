package demo

import (
	"TermCraft/internal/languages/java"
	"TermCraft/internal/languages/python"
	"fmt"
	"strconv"
)

func DemoInstallJavaVersion(id string) {
	java.InstallJavaVersion(id)
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
