package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var (
	systemInfoSection        SystemInfoComponent
	availableLanguesSections AvailableLanguagesList
)

var (
	systemInformationPositions = [7]int{0, 0, 1, 1, 0, 0}
	languageInfoPositions      = [7]int{1, 0, 1, 1, 0, 0}
)

// Define an interface for elements that have SetInputCapture method
type InputCapturable interface {
	SetInputCapture(capture func(event *tcell.EventKey) *tcell.EventKey) *tview.Box
}

func addKeyFunc(el InputCapturable, k tcell.Key, f func()) {
	el.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case k:
			f()
		case tcell.KeyRune:
			switch event.Rune() {
			case rune(k):
				f()
			}
		}
		return event
	})
}
