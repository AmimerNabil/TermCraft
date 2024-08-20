package ui

import (
	"TermCraft/internal/languages"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type LanguagePane struct {
	title     string
	languages []string
	list      *tview.List
}

var langPane2 LanguagePane = LanguagePane{
	title:     "Languages",
	languages: languages.SupportedLanguages,
	list:      tview.NewList(),
}

var langPane1 LanguagePane = LanguagePane{
	title:     "Languages",
	languages: languages.SupportedLanguages,
	list:      tview.NewList(),
}

func (l *LanguagePane) getView() *tview.List {
	l.list.SetBorder(true)
	l.list.SetTitle(l.title)
	l.list.SetTitleColor(tcell.ColorWheat)

	for _, v := range languages.SupportedLanguages {
		l.list.AddItem(v, "", rune(v[0]), nil)
	}

	l.list.SetRect(0, 0, 40, 40)

	return l.list
}
