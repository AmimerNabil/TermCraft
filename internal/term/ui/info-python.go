package ui

import (
	"TermCraft/internal/languages/python"
	"fmt"
	"sort"
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type PythonPanel struct {
	El *tview.Grid

	currPython   *tview.TextView
	pythonsLocal *tview.Flex
	confirmation *tview.Flex
	treeRemote   *tview.TreeView

	inUseText string
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
		"[yellow]Python Version:[-] %s\n[yellow]Build Information:[-] %s\n[yellow]Compiler:[-] %s\n[yellow]Version Info:[-] %s\n[yellow]Pip Version:[-] %s",
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

func (pp *PythonPanel) initializeCurrLocalVersions() {
	pythons := python.GetAvailPythonLocal()
	list := tview.NewList().ShowSecondaryText(false)

	for i, python := range pythons {
		var inUse string

		if strings.Contains(python, "*") {
			inUse = "-> using"
		} else {
			inUse = ""
		}

		formatedText := fmt.Sprintf("(%s) %s", python, inUse)
		list.AddItem(formatedText, "", rune('a'+i), nil)
	}

	flex := tview.NewFlex().
		AddItem(nil, 1, 1, false). // Add 2-unit padding to the left
		AddItem(list, 0, 1, true).
		AddItem(nil, 5, 1, false) // Add 2-unit padding to the right

	flexV := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(nil, 2, 1, false). // Add 2-unit padding to the left
		AddItem(flex, 0, 1, true).
		AddItem(nil, 1, 1, false) // Add 2-unit padding to the right

	flexV.SetBorder(true).SetTitle("Currently Installed Versions").SetTitleColor(
		tview.Styles.TertiaryTextColor,
	)

	list.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey { return event })
	pp.El.AddItem(flexV, 1, 0, 1, 1, 0, 0, false)

	pp.pythonsLocal = flexV

	list.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyTab:
			App.SetFocus(pp.treeRemote)
			return nil
		case tcell.KeyEscape:
			App.SetFocus(AvailableLanguesSections.El)
		}
		return event
	})
}

func (pp *PythonPanel) initializeRemotePythons() {
	pythonsRemote := python.GetAvailableRemoteVersionsToInstall() // map[string]map[string][]string

	rootNode := tview.NewTreeNode("Python Versions")

	vendors := make([]string, 0, len(pythonsRemote))
	for vendor := range pythonsRemote {
		vendors = append(vendors, vendor)
	}
	sort.Strings(vendors)

	for _, vendor := range vendors {
		vendorNode := tview.NewTreeNode(vendor).SetColor(tview.Styles.SecondaryTextColor).SetExpanded(false)
		versionsMap := pythonsRemote[vendor]

		versionKeys := make([]string, 0, len(versionsMap))
		for version := range versionsMap {
			versionKeys = append(versionKeys, version)
		}
		sort.Strings(versionKeys)

		for _, version := range versionKeys {
			versionNode := tview.NewTreeNode(version).SetColor(tview.Styles.TertiaryTextColor).SetExpanded(false)
			versionDetails := versionsMap[version]

			for _, detail := range versionDetails {
				detailNode := tview.NewTreeNode(detail).SetColor(tview.Styles.PrimaryTextColor)

				versionNode.AddChild(detailNode)
			}
			vendorNode.AddChild(versionNode)
		}
		rootNode.AddChild(vendorNode)
	}

	treeView := tview.NewTreeView().SetRoot(rootNode).SetCurrentNode(rootNode)
	treeView.SetSelectedFunc(func(node *tview.TreeNode) {
		if node.IsExpanded() {
			node.SetExpanded(false)
		} else {
			node.SetExpanded(true)
		}
	})

	treeView.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyEscape:
			App.SetFocus(AvailableLanguesSections.El)
		case tcell.KeyTab:
			App.SetFocus(pp.pythonsLocal)
			return nil
		case tcell.KeyRune:
			switch event.Rune() {
			}
		}
		return event
	})

	flex := tview.NewFlex().
		AddItem(nil, 1, 1, false). // Add 2-unit padding to the left
		AddItem(treeView, 0, 1, true).
		AddItem(nil, 1, 1, false) // Add 2-unit padding to the right

	flex.SetTitle("Versions Available for download").SetBorder(true).SetTitleColor(tview.Styles.TertiaryTextColor)

	pp.treeRemote = treeView
	pp.El.AddItem(flex, 1, 1, 1, 1, 0, 0, false)
}
