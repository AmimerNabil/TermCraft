package main

import (
	"TermCraft/configs"
	"TermCraft/internal/term/ui"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"slices"

	"github.com/rivo/tview"
)

var App tview.Application

func main() {
	// Define the update flags
	updateFlagShort := flag.Bool("U", false, "Update Termcraft to the latest version")
	updateFlagLong := flag.Bool("update", false, "Update Termcraft to the latest version")

	// Parse command-line arguments
	flag.Parse()

	// If either -U or --update is provided, run the update script
	if *updateFlagShort || *updateFlagLong {
		err := runUpdateScript()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error updating Termcraft: %v\n", err)
			os.Exit(1)
		}
		return
	}

	if !slices.Contains(configs.SupportedOS, runtime.GOOS) {
		log.Panic("Unsupported OS...")
	}

	App = *tview.NewApplication()
	ui.Start(&App)

	if err := App.Run(); err != nil {
		panic(err)
	}
}

// runUpdateScript executes the update.sh script to update Termcraft
func runUpdateScript() error {
	fmt.Println("Updating Termcraft...")

	// Change the path to your update script if it's located elsewhere
	cmd := exec.Command("bash", "$HOME/.termcraft/update.sh")

	// Redirect output to the console
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Run the update script
	return cmd.Run()
}
