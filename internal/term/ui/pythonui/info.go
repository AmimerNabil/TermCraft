package pythonui

import "github.com/rivo/tview"

type PythonPanel struct {
	El *tview.Flex
}

func (jp *PythonPanel) Init() *tview.Flex {
	jp.El = tview.NewFlex()
	jp.El.SetTitle("Python Language Information").SetBorder(true)

	return jp.El
}
