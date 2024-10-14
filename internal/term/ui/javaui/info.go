package javaui

import (
	"TermCraft/internal/languages/java"
	"fmt"
	"sort"
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type JavaPanel struct {
	El *tview.Grid
}

// java pannel elements
var (
	Civ *tview.Flex
	Liv *tview.Flex
	Rvs *tview.Flex
)

var (
	App      *tview.Application
	OutFocus *tview.List
)

func (jp *JavaPanel) Init(app *tview.Application, outFocus *tview.List) *tview.Grid {
	App = app
	OutFocus = outFocus

	Civ = currentlyInstalledVersion()
	Liv = createJavaListView()
	Rvs = CreateJavaTreeView()

	jp.El = tview.NewGrid()
	// jp.El.SetTitle("Java Language Information").SetBorder(true)

	// first sow at the very top, the
	jp.El.SetRows(14, 0)
	jp.El.SetColumns(40, 0)

	jp.El.AddItem(Civ, 0, 0, 1, 2, 0, 0, false)
	jp.El.AddItem(Liv, 1, 0, 1, 1, 0, 0, false)
	jp.El.AddItem(Rvs, 1, 1, 1, 1, 0, 0, false)

	return jp.El
}

func createJavaListView() *tview.Flex {
	// Create a new tview List
	list := tview.NewList()
	list.ShowSecondaryText(false)

	javas := java.GetLocalJavaVersionsSdk()

	for i, java := range javas {
		list.AddItem(fmt.Sprintf("Java %s (%s)", java.JavaVersion, java.JavaVendor), "", rune('a'+i), nil)
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

	list.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyEscape:
			App.SetFocus(OutFocus)
		case tcell.KeyTab:
			App.SetFocus(Rvs)
			return nil
		case tcell.KeyRune:
			switch event.Rune() {
			}
		}
		return event
	})
	return flexV
}

func currentlyInstalledVersion() *tview.Flex {
	textView := tview.NewTextArea()

	props := java.GetAllJavaVersionInformation("java")

	// Format the content
	content := []string{
		fmt.Sprintf("Java Home: %s", props.JavaHome),
		fmt.Sprintf("Runtime Name: %s", props.JavaRuntimeName),
		fmt.Sprintf("Java Version: %s", props.JavaVersion),
		fmt.Sprintf("Vendor: %s", props.JavaVendor),
		fmt.Sprintf("VM Name: %s", props.JavaVMName),
		fmt.Sprintf("VM Version: %s", props.JavaVMVersion),
		fmt.Sprintf("OS Architecture: %s", props.OSArch),
		fmt.Sprintf("OS Name: %s", props.OSName),
		fmt.Sprintf("OS Version: %s", props.OSVersion),
		fmt.Sprintf("User Name: %s", props.UserName),
	}

	// Set the formatted content to the TextView
	textView.SetText(strings.Join(content, "\n"), false)

	// Create a Flex layout to wrap the TextView and add some padding
	flex := tview.NewFlex().
		AddItem(nil, 5, 1, false). // Add 2-unit padding to the left
		AddItem(textView, 0, 1, true).
		AddItem(nil, 5, 1, false) // Add 2-unit padding to the right

	flexV := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(nil, 1, 1, false). // Add 2-unit padding to the left
		AddItem(flex, 0, 1, true).
		AddItem(nil, 1, 1, false) // Add 2-unit padding to the right

	flexV.SetBorder(true).SetTitle("Currently Used").SetTitleColor(tview.Styles.TertiaryTextColor)

	return flexV
}

func CreateJavaTreeView() *tview.Flex {
	javas := java.GetRemoteVersions()

	rootNode := tview.NewTreeNode("Java Versions")
	vendorMap := make(map[string]map[string][]string)

	for _, java := range javas {
		versionParts := strings.Split(java.JavaVersion, ".")
		if len(versionParts) < 2 {
			continue
		}

		majorMinor := fmt.Sprintf("%s.%s", versionParts[0], versionParts[1])

		if _, ok := vendorMap[java.JavaVendor]; !ok {
			vendorMap[java.JavaVendor] = make(map[string][]string)
		}

		vendorMap[java.JavaVendor][majorMinor] = append(vendorMap[java.JavaVendor][majorMinor], java.JavaVersion)
	}

	// Sort vendor keys alphabetically
	vendors := make([]string, 0, len(vendorMap))
	for vendor := range vendorMap {
		vendors = append(vendors, vendor)
	}
	sort.Strings(vendors)

	for _, vendor := range vendors {
		vendorNode := tview.NewTreeNode(vendor).SetColor(tview.Styles.SecondaryTextColor).SetExpanded(false)

		// Sort majorMinor keys alphabetically
		majorMinorVersions := vendorMap[vendor]
		majorMinorKeys := make([]string, 0, len(majorMinorVersions))
		for majorMinor := range majorMinorVersions {
			majorMinorKeys = append(majorMinorKeys, majorMinor)
		}
		sort.Strings(majorMinorKeys)

		for _, majorMinor := range majorMinorKeys {
			majorMinorNode := tview.NewTreeNode(majorMinor).SetColor(tview.Styles.TertiaryTextColor).SetExpanded(false)

			// Sort versions alphabetically
			versions := majorMinorVersions[majorMinor]
			sort.Strings(versions)

			for _, version := range versions {
				versionNode := tview.NewTreeNode(version).SetColor(tview.Styles.PrimaryTextColor)
				majorMinorNode.AddChild(versionNode)
			}

			vendorNode.AddChild(majorMinorNode)
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

	flex := tview.NewFlex().
		AddItem(nil, 1, 1, false). // Add 2-unit padding to the left
		AddItem(treeView, 0, 1, true).
		AddItem(nil, 1, 1, false) // Add 2-unit padding to the right
	flex.SetTitle("Versions Available for download").SetBorder(true).SetTitleColor(tview.Styles.TertiaryTextColor)

	treeView.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyEscape:
			App.SetFocus(OutFocus)
		case tcell.KeyTab:
			App.SetFocus(Liv)
		case tcell.KeyRune:
			switch event.Rune() {
			}
		}
		return event
	})

	return flex
}
