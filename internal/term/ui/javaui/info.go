package javaui

import (
	"TermCraft/internal/languages/java"
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type JavaPanel struct {
	El           *tview.Grid
	Civ          *tview.Flex
	Liv          *tview.Flex
	Rvs          *tview.Flex
	confirmation *tview.Flex
}

var (
	App        *tview.Application
	OutFocus   *tview.List
	indexInUse int
	inUseText  string
)

func (jp *JavaPanel) Init(app *tview.Application, outFocus *tview.List) *tview.Grid {
	App = app
	OutFocus = outFocus

	jp.reload()

	return jp.El
}

func (jp *JavaPanel) reload() {
	jp.Civ = jp.currentlyInstalledVersion()
	jp.Liv = jp.createJavaListView()
	jp.Rvs = jp.CreateJavaTreeView()

	jp.El = tview.NewGrid()

	// first sow at the very top, the
	jp.El.SetRows(14, 0)
	jp.El.SetColumns(-1, -1)

	jp.El.AddItem(jp.Civ, 0, 0, 1, 2, 0, 0, false)
	jp.El.AddItem(jp.Liv, 1, 0, 1, 1, 0, 0, false)
	jp.El.AddItem(jp.Rvs, 1, 1, 1, 1, 0, 0, false)
}

func (jp *JavaPanel) createJavaListView() *tview.Flex {
	// Create a new tview List
	list := tview.NewList()
	list.ShowSecondaryText(false)

	javas := java.GetLocalJavaVersionsSdk()

	for i, java := range javas {
		var inUse string

		if java.InUse {
			inUse = "-> using"
			indexInUse = i
		} else {
			inUse = ""
		}

		inUseText = fmt.Sprintf("(%s)\t id: %s %s", java.JavaVendor, java.Identifier, inUse)
		list.AddItem(inUseText, "", rune('a'+i), nil)
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
		index := list.GetCurrentItem()
		var text string
		if len(javas) > 0 && index >= 0 {
			itemText, _ := list.GetItemText(index)
			text = itemText
		}

		switch event.Key() {
		case tcell.KeyEnter:
			jp.UseVersion(getVersionFromID(text), index, list)
		case tcell.KeyEscape:
			App.SetFocus(OutFocus)
		case tcell.KeyTab:
			App.SetFocus(jp.Rvs)
			return nil
		case tcell.KeyRune:
			switch event.Rune() {
			case 'D':
				jp.DeleteVersion(getVersionFromID(text), index, list)
			}
		}
		return event
	})

	return flexV
}

func (jp *JavaPanel) currentlyInstalledVersion() *tview.Flex {
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

	flexV := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(nil, 1, 1, false). // Add 2-unit padding to the left
		AddItem(textView, 0, 1, true).
		AddItem(nil, 1, 1, false) // Add 2-unit padding to the right

	flex := tview.NewFlex().
		AddItem(nil, 5, 1, false). // Add 2-unit padding to the left
		AddItem(flexV, 0, 1, true).
		AddItem(nil, 5, 1, false) // Add 2-unit padding to the right

	flex.SetBorder(true).SetTitle("Currently Used").SetTitleColor(tview.Styles.TertiaryTextColor)

	return flex
}

func (jp *JavaPanel) CreateJavaTreeView() *tview.Flex {
	javas := java.GetRemoteVersions()

	rootNode := tview.NewTreeNode("Java Versions")
	vendorMap := make(map[string]map[string][]java.RemoteJavaProperties)

	for _, j := range javas {
		versionParts := strings.Split(j.JavaVersion, ".")
		if len(versionParts) < 2 {
			continue
		}

		majorMinor := fmt.Sprintf("%s.%s", versionParts[0], versionParts[1])

		if _, ok := vendorMap[j.JavaVendor]; !ok {
			vendorMap[j.JavaVendor] = make(map[string][]java.RemoteJavaProperties)
		}

		vendorMap[j.JavaVendor][majorMinor] = append(vendorMap[j.JavaVendor][majorMinor], j)
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
			// sort.Strings(versions)

			for _, version := range versions {
				var installed string
				if version.Installed {
					installed = "*"
				} else {
					installed = ""
				}

				versionNode := tview.NewTreeNode(version.JavaVersion + " " + installed).SetColor(tview.Styles.PrimaryTextColor)
				versionNode.SetSelectedFunc(func() {
					// ask for confirmation
					jp.confirmation = tview.NewFlex().SetDirection(tview.FlexColumnCSS)
					currFocus := '2'
					b1 := tview.NewButton("Yes").SetSelectedFunc(func() {
						jp.Civ.RemoveItem(jp.confirmation)
						jp.InstallVersion(version.Identifier, versionNode)
						App.SetFocus(jp.Rvs)
					})
					b2 := tview.NewButton("No").SetSelectedFunc(func() {
						jp.Civ.RemoveItem(jp.confirmation)
						App.SetFocus(jp.Rvs)
					})

					jp.confirmation.
						AddItem(nil, 4, 1, false).
						AddItem(
							tview.NewTextArea().
								SetText("Are you sure you want to Install: "+version.Identifier, false),
							0, 3, false).
						AddItem(
							tview.NewFlex().SetDirection(tview.FlexRowCSS).
								AddItem(b1, 0, 1, false).
								AddItem(nil, 4, 3, false).
								AddItem(b2, 0, 1, true),
							0, 2, true)

					jp.confirmation.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
						switch event.Key() {
						case tcell.KeyTab:
							if currFocus == '1' {
								App.SetFocus(b2)
								currFocus = '2'
							} else {
								App.SetFocus(b1)
								currFocus = '1'
							}
						}
						return event
					})

					jp.Civ.AddItem(jp.confirmation, 0, 1, false)

					App.SetFocus(jp.confirmation)
				})
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
			App.SetFocus(jp.Liv)
		case tcell.KeyRune:
			switch event.Rune() {
			}
		}
		return event
	})

	return flex
}

func (jp *JavaPanel) DeleteVersion(identifier string, index int, list *tview.List) {
	done := make(chan bool)
	originalText, _ := list.GetItemText(index)

	go func() {
		rotation := []string{"|", "/", "-", "\\"}
		i := 0
		for {
			select {
			case <-done:
				return
			default:
				list.SetItemText(index, fmt.Sprintf("%s %s", originalText, rotation[i%len(rotation)]), "")
				i++
				App.Draw()
				time.Sleep(200 * time.Millisecond)
			}
		}
	}()

	go func() {
		_, stderr, err := java.UnInstallJavaVersion(identifier)

		if err != nil || stderr != "" {
			list.SetItemText(index, fmt.Sprintf("%s - Failed to delete %s: %s", originalText, identifier, stderr), "")
			App.Draw()
		} else {
			list.RemoveItem(index)
			App.Draw()
		}

		done <- true
	}()
}

func (jp *JavaPanel) UseVersion(identifier string, index int, list *tview.List) {
	done := make(chan bool)
	originalText, _ := list.GetItemText(index)

	if indexInUse == index {
		return
	}

	go func() {
		rotation := []string{"|", "/", "-", "\\"}
		i := 0
		for {
			select {
			case <-done:

				temp := fmt.Sprintf("%s-> using ", originalText)
				var newOld string
				newOld = strings.ReplaceAll(inUseText, "-> using", "")

				list.SetItemText(index, temp, "")
				list.SetItemText(indexInUse, newOld, "")

				indexInUse = index
				inUseText = temp

				jp.El.RemoveItem(jp.Civ)
				jp.Civ = jp.currentlyInstalledVersion()
				jp.El.AddItem(jp.Civ, 0, 0, 1, 2, 0, 0, false)

				App.Draw()

				return
			default:
				list.SetItemText(index, fmt.Sprintf("%s %s", originalText, rotation[i%len(rotation)]), "")
				i++
				App.Draw()
				time.Sleep(200 * time.Millisecond)
			}
		}
	}()

	go func() {
		_, stderr, err := java.SetJavaVersion(identifier)

		if err != nil || stderr != "" {
		} else {
			App.Draw()
		}

		done <- true
	}()
}

func (jp *JavaPanel) InstallVersion(identifier string, node *tview.TreeNode) {
	done := make(chan bool)
	originalText := node.GetText()

	go func() {
		rotation := []string{"|", "/", "-", "\\"} // Symbols for the spinner
		i := 0
		for {
			select {
			case <-done:
				node.SetText(fmt.Sprintf("%s *", originalText))
				jp.El.RemoveItem(jp.Liv)
				jp.Liv = jp.createJavaListView()
				jp.El.AddItem(jp.Liv, 1, 0, 1, 1, 0, 0, false)
				App.Draw()
				return
			default:
				// Update the node text, appending the spinner to the original text
				node.SetText(fmt.Sprintf("%s %s", originalText, rotation[i%len(rotation)]))
				App.Draw()
				i++
				time.Sleep(200 * time.Millisecond) // Controls the speed of the spinner
			}
		}
	}()

	go func() {
		_, stderr, err := java.InstallJavaVersion(identifier)
		if err != nil || stderr != "" {
			node.SetText(fmt.Sprintf("%s - Failed to install %s: %s", originalText, identifier, stderr))
		} else {
			node.SetText(fmt.Sprintf("%s - Installed %s successfully!", originalText, identifier))
		}
		done <- true
	}()
}

func getVersionFromID(input string) string {
	parts := strings.SplitN(input, "id: ", 2)

	if len(parts) == 2 {
		return strings.TrimSpace(parts[1]) // Trim any extra spaces
	}

	return ""
}
