package main

import (
	"TermCraft/configs"
	"TermCraft/internal/term/ui"
	"log"
	"runtime"
	"slices"

	"github.com/rivo/tview"
)

var App tview.Application

func main() {
	// out, serr, err := python.UnInstallPythonVersion("2.1.3")
	// fmt.Println("out " + out)
	// fmt.Println("err " + serr)
	// fmt.Println(err)

	if !slices.Contains(configs.SupportedOS, runtime.GOOS) {
		log.Panic("Unsupported OS...")
	}

	App = *tview.NewApplication()
	ui.Start(&App)

	if err := App.Run(); err != nil {
		panic(err)
	}
}
