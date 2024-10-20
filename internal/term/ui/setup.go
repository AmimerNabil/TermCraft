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
	// go func() {
	// 	App.QueueUpdate(func() {
	// 		pp.Init()
	// 		mainContainer.AddPage("python", pp.El, true, false) // Add Python page when done
	// 	})
	// 	pdone <- true // Signal that Python is initialized
	// }()

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

	confirmation = tview.NewFlex().SetDirection(tview.FlexColumnCSS)
	confirmation.SetBorder(true).SetTitle("Confirm")
	confButton1 = tview.NewButton("Yes").SetSelectedFunc(nil)
	confButton2 = tview.NewButton("No").SetSelectedFunc(nil)
	confTextView = tview.NewTextView()
	confirmation.
		AddItem(confTextView, 0, 3, false).
		AddItem(
			tview.NewFlex().SetDirection(tview.FlexRowCSS).
				AddItem(nil, 2, 1, false).
				AddItem(confButton1, 0, 1, false).
				AddItem(nil, 4, 3, false).
				AddItem(confButton2, 0, 1, true).
				AddItem(nil, 2, 1, false),
			0, 2, true)

	currFocus := '2'
	confirmation.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyTab:
			if currFocus == '1' {
				App.SetFocus(confButton2)
				currFocus = '2'
			} else {
				App.SetFocus(confButton1)
				currFocus = '1'
			}
		}
		return event
	})

	commandsPages.AddPage("Main", frame, true, true)
	commandsPages.AddPage("Command", modal(commandText, 70), true, false)
	commandsPages.AddPage("Confirmation", modal(confirmation, 40), true, false)

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

func setConfirmationContent(whatToConfirm string, ifNo func(), ifYes func()) {
	commandsPages.ShowPage("Confirmation")

	confButton1.SetSelectedFunc(func() {
		commandsPages.HidePage("Confirmation")
		ifYes()
	})
	confButton2.SetSelectedFunc(func() {
		commandsPages.HidePage("Confirmation")
		ifNo()
	})
	confTextView.SetText(whatToConfirm)
}
