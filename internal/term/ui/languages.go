package ui

import (
	"TermCraft/internal/languages"
	commandtext "TermCraft/internal/term/ui/command-text"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type AvailableLanguagesList struct {
	El     *tview.List
	Lgs    []string
	Height int
}

func (list *AvailableLanguagesList) Init() {
	list.El = tview.NewList()
	list.El.SetWrapAround(false)
	list.Lgs = languages.SupportedLanguages

	list.Height = len(list.Lgs)*2 + 2

	if list.Height > 7 {
		list.Height = 16
	}

	for index, s := range list.Lgs {
		list.El.AddItem(s, "", rune('a'+index), nil)
	}

	list.El.SetBorder(true)
	list.El.SetTitle("languages")

	list.El.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyRune:
			switch event.Rune() {
			case '[':
				App.SetFocus(systemInfoSection.El)
			case '?':
				commandText.SetText(commandtext.AvailLangCommands)
				commandsPages.ShowPage("Command")
			}
		}
		return event
	})

	list.El.SetChangedFunc(func(index int, mainText string, sec string, shortcut rune) {
		mainContainer.SwitchToPage(mainText)
	})

	list.El.SetSelectedFunc(func(index int, mainText string, sec string, shortcut rune) {
		switch mainText {
		case "java":
			App.SetFocus(jp.Liv)
			// case "python":
			// 	App.SetFocus(pp.pythonsLocal)
		}
	})
}
