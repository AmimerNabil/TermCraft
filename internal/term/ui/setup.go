package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type Pane struct {
	flexPane  *tview.Flex
	currFocus int
}

type ConfigPane struct {
	title       string
	configFiles []string
}

var configPane ConfigPane = ConfigPane{
	title:       "Configuration Files",
	configFiles: []string{},
}

var (
	rootFlex  *tview.Flex = tview.NewFlex()
	leftFlex  *tview.Flex = tview.NewFlex()
	rightFlex *tview.Flex = tview.NewFlex()
)

var (
	rootPane *Pane = &Pane{
		flexPane: rootFlex,
	}
	leftPane *Pane = &Pane{
		flexPane: leftFlex,
	}
	rightPane *Pane = &Pane{
		flexPane: rightFlex,
	}
)

var App *tview.Application

func Start(app *tview.Application) {
	App = app

	rootPane.flexPane.AddItem(leftFlex, 0, 1, true)
	rootPane.flexPane.AddItem(rightFlex, 0, 1, false)

	leftPane.flexPane.SetDirection(tview.FlexRow)
	leftPane.flexPane.
		AddItem(langPane1.getView(), 0, 2, true).
		AddItem(langPane2.getView(), 0, 1, false)
	leftPane.flexPane.SetInputCapture(leftPaneInputCapture)

	app.SetRoot(rootFlex, true)
}

func leftPaneInputCapture(event *tcell.EventKey) *tcell.EventKey {
	switch event.Key() {
	case tcell.KeyTab:
		switch {
		case langPane1.list.HasFocus():
			App.SetFocus(langPane2.list)
		case langPane2.list.HasFocus():
			App.SetFocus(langPane1.list)
		}
	}
	return event
}
