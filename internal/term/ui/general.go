package ui

import (
	"github.com/rivo/tview"
)

var (
	// main
	systemInfoSection        SystemInfoComponent
	AvailableLanguesSections AvailableLanguagesList
	mainContainer            *tview.Pages
	commandsPages            *tview.Pages
	commandText              *tview.TextView
	oldFocus                 tview.Primitive

	// language specific
	jp JavaPanel
	pp PythonPanel

	// config specific
	// TODO:config for later things like .bashrc .zshrc
)

var (
	systemInformationPositions  = [7]int{0, 0, 1, 1, 0, 0}
	availableLanguagesPositions = [7]int{1, 0, 1, 1, 0, 0}
)
