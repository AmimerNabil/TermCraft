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

	confirmation *tview.Flex
	confButton1  *tview.Button
	confButton2  *tview.Button
	confTextView *tview.TextView

	// language specific
	jp *JavaPanel
	pp *PythonPanel

	// mapping
	lgs map[string]*Panel

	// synchro channels
	pdone = make(chan bool)
	jdone = make(chan bool)
)

var (
	systemInformationPositions  = [7]int{0, 0, 1, 1, 0, 0}
	availableLanguagesPositions = [7]int{1, 0, 1, 1, 0, 0}
)
