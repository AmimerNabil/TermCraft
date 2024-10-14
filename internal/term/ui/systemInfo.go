package ui

import (
	"TermCraft/configs"
	"os"
	"runtime"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type SystemInfoComponent struct {
	El *tview.Grid
}

func (si *SystemInfoComponent) Init() {
	si.El = tview.NewGrid().
		SetRows(1, 1, 1, 1) // Three rows for OS info

	osName := runtime.GOOS
	hostname, _ := os.Hostname()
	goVersion := runtime.Version()

	si.El.AddItem(tview.NewTextView().SetText("OS: "+osName), 0, 0, 1, 1, 0, 0, false)
	si.El.AddItem(tview.NewTextView().SetText("Hostname: "+hostname), 1, 0, 1, 1, 0, 0, false)
	si.El.AddItem(tview.NewTextView().SetText("Go Version: "+goVersion), 2, 0, 1, 1, 0, 0, false)
	si.El.AddItem(tview.NewTextView().SetText("TermCraft Version: "+configs.AppVersion), 3, 0, 1, 1, 0, 0, false)

	si.El.SetTitle("System Info")
	si.El.SetBorder(true)

	si.El.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyDown, tcell.KeyUp, tcell.KeyLeft, tcell.KeyRight:
			return nil
		case tcell.KeyRune:
			switch event.Rune() {
			case 'j', 'l', 'h', 'k':
				return nil
			case ']':
				App.SetFocus(AvailableLanguesSections.El)
			}
		}
		return event
	})
}
