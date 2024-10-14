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
		positions[4], // mintWidth (3rd position in array)
		positions[5], // mingHeight (3rd position in array)
		focus)        // Boolean flag
}

func Start(app *tview.Application) {
	App = app

	// initiate sections
	systemInfoSection.Init()
	AvailableLanguesSections.Init()

	jp.Init(App, AvailableLanguesSections.El)
	pp.Init()

	mainContainer = tview.NewPages()
	// mainContainer.SetTitle("MainContainer").SetTitleAlign(tview.AlignCenter).SetBorder(true)

	// set the different NewPages
	mainContainer.AddPage("java", jp.El, true, true)
	mainContainer.AddPage("python", pp.El, true, false)

	// set columns for sections
	mainGrid := tview.NewGrid().
		SetRows(
			6,
			0,
		).
		SetColumns(30, 0)

		// add the items on the UI
	addItem(mainGrid, systemInfoSection.El, systemInformationPositions, false)
	addItem(mainGrid, AvailableLanguesSections.El, availableLanguagesPositions, true)

	addItem(mainGrid, mainContainer, [7]int{
		0, 1, 2, 1, 0, 0,
	}, false)

	frame := tview.NewFrame(mainGrid)

	App.SetRoot(frame, true)
}
