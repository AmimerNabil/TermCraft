package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var (
	systemInfoSection        SystemInfoComponent
	availableLanguesSections AvailableLanguagesList
	languageInfoSection      LanguageInfo
)

var (
	systemInformationPositions  = [7]int{0, 0, 1, 1, 0, 0}
	availableLanguagesPositions = [7]int{1, 0, 1, 1, 0, 0}
	languageInfoPositions       = [7]int{0, 1, 2, 1, 0, 0}
)

// Define an interface for elements that have SetInputCapture method
type InputCapturable interface {
	SetInputCapture(capture func(event *tcell.EventKey) *tcell.EventKey) *tview.Box
}
