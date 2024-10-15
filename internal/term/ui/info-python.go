package ui

import (
	"TermCraft/internal/languages/python"
	"fmt"
	"strings"

	"github.com/rivo/tview"
)

type PythonPanel struct {
	El *tview.Grid

	currPython *tview.TextView
}

func (pp *PythonPanel) Init() *tview.Grid {
	pp.El = tview.NewGrid()

	pp.initializeCurrPython()
	pp.initializeRemotePythons()
	pp.initializeCurrLocalVersions()

	// first sow at the very top, the
	pp.El.SetRows(14, 0)
	pp.El.SetColumns(-1, -1)

	return pp.El
}

func (pp *PythonPanel) initializeCurrPython() {
	info := python.GetPythonLocal()

	if pp.currPython != nil {
		pp.El.RemoveItem(pp.currPython)
		pp.currPython = nil
	}
	lines := strings.Split(info, "\n")

	// Assign each part of the string to corresponding variables
	pythonVersion := lines[0]
	pythonBuild := lines[1]
	pythonCompiler := lines[2]
	versionInfo := lines[3]
	pipVersion := lines[7]

	// Format the text to make it more readable
	formattedText := fmt.Sprintf(
		"[::b]Python Version:[::-]%s\n[::b]Build Information:[::-]%s\n[::b]Compiler:[::-]%s\n[::b]Version Info:[::-]%s\n[::b]Pip Version:[::-]\n%s",
		pythonVersion,
		pythonBuild,
		pythonCompiler,
		versionInfo,
		pipVersion,
	)

	pp.currPython = tview.NewTextView()
	pp.currPython.SetText(formattedText).
		SetDynamicColors(true). // Enable dynamic color tags
		SetTextAlign(tview.AlignLeft).
		SetWrap(true)

	flexPaddingH := tview.NewFlex().
		AddItem(nil, 5, 1, false).
		AddItem(pp.currPython, 0, 1, true).
		AddItem(nil, 5, 1, false)

	flexPaddingV := tview.NewFlex().SetDirection(tview.FlexColumnCSS).
		AddItem(nil, 1, 1, false).
		AddItem(flexPaddingH, 0, 1, true).
		AddItem(nil, 2, 1, false)

	flexPaddingV.SetBorder(true).SetTitle("Currently Used").SetTitleColor(tview.Styles.TertiaryTextColor)

	pp.El.AddItem(flexPaddingV, 0, 0, 1, 2, 0, 0, false)
}

func (pp *PythonPanel) initializeRemotePythons() {
}

func (pp *PythonPanel) initializeCurrLocalVersions() {
}
