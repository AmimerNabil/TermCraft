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

	systemInfoSection.Init()
	AvailableLanguesSections.Init()
	lgs = make(map[string]*Panel)

	initializePanels()

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

	go func() {
		<-jdone
		<-pdone
		<-ndone
		app.Draw()
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

func initializePanels() {
	go func() {
		key := "java"
		jp = NewJavaPanel()
		lgs[key] = jp.Panel

		App.QueueUpdate(func() {
			mainContainer.AddPage(key, jp.container, true, true) // Add Python page when done
		})

		jdone <- true // Signal that Python is initialized
	}()

	go func() {
		key := "python"
		pp = NewPythonPanel()
		lgs[key] = pp.Panel

		App.QueueUpdate(func() {
			mainContainer.AddPage(key, pp.container, true, false) // Add Python page when done
		})

		pdone <- true // Signal that Python is initialized
	}()

	go func() {
		key := "node"
		np = NewFNMPanel()
		lgs[key] = np.Panel

		App.QueueUpdate(func() {
			mainContainer.AddPage(key, np.container, true, false) // Add Python page when done
		})

		ndone <- true // Signal that Python is initialized
	}()
}
