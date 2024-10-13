package main

import (
	"TermCraft/internal/demo"

	"github.com/rivo/tview"
)

var App tview.Application

func main() {
	// if !slices.Contains(configs.SupportedOS, runtime.GOOS) {
	// 	log.Panic("Unsupported OS...")
	// }
	//
	// App = *tview.NewApplication()
	// ui.Start(&App)
	//
	// if err := App.Run(); err != nil {
	// 	panic(err)
	// }

	// List of common Python executable names

	demo.DemoGetAvailablePythonVersionsToInstall()
}
