package ui

import (
	"github.com/rivo/tview"
)

var App *tview.Application

func addItem(root *tview.Grid, el tview.Primitive, positions [7]int, focus bool) {
	root.AddItem(el,
		positions[0], // Row
		positions[1], // Column
		positions[2], // RowSpan (2nd position in array)
		positions[3], // ColSpan (3rd position in array)
		positions[4], // ColSpan (3rd position in array)
		positions[5], // ColSpan (3rd position in array)
		focus)        // Boolean flag
}

func Start(app *tview.Application) {
	App = app

	// initiate sections
	systemInfoSection.Init()
	availableLanguesSections.Init()

	// set columns for sections
	mainGrid := tview.NewGrid().
		SetRows(
			6,
			availableLanguesSections.Height,
		).
		SetColumns(40)

		// add the items on the UI
	addItem(mainGrid, systemInfoSection.El, systemInformationPositions, true)
	addItem(mainGrid, availableLanguesSections.El, languageInfoPositions, false)

	App.SetRoot(mainGrid, true)
}
