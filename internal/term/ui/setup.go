package ui

import (
	"TermCraft/internal/languages/java"

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

	// Example JavaProperties object
	javaProps := []java.JavaProperties{
		{
			JavaVendor:      "Oracle Corporation",
			JavaVersion:     "1.8.0_292",
			JavaHome:        "/usr/lib/jvm/java-8-oracle",
			JavaRuntimeName: "Java(TM) SE Runtime Environment",
		},
		{
			JavaVendor:      "Oracle Corporation",
			JavaVersion:     "1.8.0_292",
			JavaHome:        "/usr/lib/jvm/java-8-oracle",
			JavaRuntimeName: "Java(TM) SE Runtime Environment",
		},
		{
			JavaVendor:      "Oracle Corporation",
			JavaVersion:     "1.8.0_292",
			JavaHome:        "/usr/lib/jvm/java-8-oracle",
			JavaRuntimeName: "Java(TM) SE Runtime Environment",
		},
	}

	// initiate sections
	systemInfoSection.Init()
	availableLanguesSections.Init()
	languageInfoSection.init(convertToInterfaces(javaProps))

	// set columns for sections
	mainGrid := tview.NewGrid().
		SetRows(
			6,
			availableLanguesSections.Height,
		).
		SetColumns(40, 0)

		// add the items on the UI
	addItem(mainGrid, systemInfoSection.El, systemInformationPositions, true)
	addItem(mainGrid, availableLanguesSections.El, availableLanguagesPositions, false)
	addItem(mainGrid, languageInfoSection.El, languageInfoPositions, false)

	App.SetRoot(mainGrid, true)
}
