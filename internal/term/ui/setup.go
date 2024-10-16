package ui

import (
	"github.com/gdamore/tcell/v2"
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

	// Create a channel to signal when both initializations are done
	jdone := make(chan bool)
	pdone := make(chan bool)

	systemInfoSection.Init()
	AvailableLanguesSections.Init()

	// Run jp.Init() asynchronously
	go func() {
		App.QueueUpdate(func() {
			jp.Init()
			mainContainer.AddPage("java", jp.El, true, true) // Add Java page when done
		})
		jdone <- true // Signal that Java is initialized
	}()

	// Run pp.Init() asynchronously
	go func() {
		App.QueueUpdate(func() {
			pp.Init()
			mainContainer.AddPage("python", pp.El, true, false) // Add Python page when done
		})
		pdone <- true // Signal that Python is initialized
	}()

	commandsPages = tview.NewPages()
	mainContainer = tview.NewPages()

	mainGrid := tview.NewGrid().
		SetRows(
			6,
			0,
		).
		SetColumns(30, 0)

	addItem(mainGrid, systemInfoSection.El, systemInformationPositions, false)
	addItem(mainGrid, AvailableLanguesSections.El, availableLanguagesPositions, true)

	addItem(mainGrid, mainContainer, [7]int{
		0, 1, 2, 1, 0, 0,
	}, false)

	frame := tview.NewFrame(mainGrid)

	modal := func(p tview.Primitive, width int) tview.Primitive {
		return tview.NewFlex().
			AddItem(nil, 0, 1, false).
			AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
				AddItem(nil, 0, 1, false).
				AddItem(p, 0, 1, true).
				AddItem(nil, 0, 1, false), width, 1, true).
			AddItem(nil, 0, 1, false)
	}

	commandText = tview.NewTextView()
	commandText.
		SetDynamicColors(true).
		SetTextAlign(tview.AlignLeft).
		SetBorder(true).
		SetTitle("Help - Commands")

	commandsPages.AddPage("Main", frame, true, true)
	commandsPages.AddPage("Command", modal(commandText, 70), true, false)

	commandText.
		SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
			switch event.Key() {
			case tcell.KeyRune:
				switch event.Rune() {
				case 'q':
					commandsPages.HidePage("Command")
				}
			}
			return event
		})

	App.SetRoot(commandsPages, true)

	// Run a goroutine to wait for both jp and pp initialization to finish
	go func() {
		<-jdone    // Wait for jp.Init() to finish
		<-pdone    // Wait for pp.Init() to finish
		app.Draw() // Redraw the application once both are done
	}()
}
