package ui

import (
	"TermCraft/internal/term/ui/javaui"
	"TermCraft/internal/term/ui/pythonui"

	"github.com/rivo/tview"
)

var (
	// main
	systemInfoSection        SystemInfoComponent
	AvailableLanguesSections AvailableLanguagesList
	mainContainer            *tview.Pages

	// language specific
	jp javaui.JavaPanel
	pp pythonui.PythonPanel

	// config specific
	// todo

)

var (
	systemInformationPositions  = [7]int{0, 0, 1, 1, 0, 0}
	availableLanguagesPositions = [7]int{1, 0, 1, 1, 0, 0}
)
