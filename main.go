package main

import (
	"TermCraft/configs"
	"TermCraft/internal/demo"
	"log"
	"runtime"
	"slices"

	"github.com/rivo/tview"
)

var App tview.Application

func main() {
	if !slices.Contains(configs.SupportedOS, runtime.GOOS) {
		log.Panic("Unsupported OS...")
	}

	demo.DemoGetJavaVersionsRemote()

	// fmt.Println(java.IsSDKMANInstalled())
	// App = *tview.NewApplication()
	// ui.Start(&App)
	//
	// if err := App.Run(); err != nil {
	// 	panic(err)
	// }
}
