package ui

import (
	"TermCraft/internal/languages/python"
	"fmt"
	"log"
	"os/exec"
	"regexp"
	"sort"
	"strings"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type PythonPanel struct {
	El *tview.Grid

	currPython   *tview.TextView
	treeRemote   *tview.TreeView
	pythonsLocal *tview.Flex
	confirmation *tview.Flex

	localPythons []string

	currGlobal string
	currLocal  string
}

func CleanVersionString(version string) string {
	cleaned := strings.TrimSpace(version)
	cleaned = strings.TrimPrefix(cleaned, "*")
	re := regexp.MustCompile(`\s*\(.*\)$`)
	cleaned = re.ReplaceAllString(cleaned, "")

	return strings.TrimSpace(cleaned)
}

func (pp *PythonPanel) Init() *tview.Grid {
	pp.El = tview.NewGrid()
	pp.localPythons = python.GetAvailPythonLocal()

	pp.currLocal = python.GetPyenvLocal()
	pp.currGlobal = python.GetPyenvGlobal()

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
	list := tview.NewList().ShowSecondaryText(false)

	for i, python := range pp.localPythons {
		inUse := ""
		global := ""

		if strings.Contains(python, "*") {
			inUse = "-> using"
		}
		if strings.Contains(python, pp.currGlobal) {
			global = "(global)"
		}

		formatedText := fmt.Sprintf("%s %s %s", python, global, inUse)
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

	pp.El.AddItem(flexV, 1, 0, 1, 1, 0, 0, false)

	pp.pythonsLocal = flexV

	list.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		index := list.GetCurrentItem()
		var text string
		if len(pp.localPythons) > 0 && index >= 0 {
			itemText, _ := list.GetItemText(index)
			text = itemText
		}

		switch event.Key() {
		case tcell.KeyTab:
			App.SetFocus(pp.treeRemote)
			return nil
		case tcell.KeyEscape:
			App.SetFocus(AvailableLanguesSections.El)
		case tcell.KeyRune:
			switch event.Rune() {
			case 'G':
				if !strings.Contains(text, "global") && !strings.Contains(text, "system") {
					pp.UsePythonVersionGlobal(CleanVersionString(text))
				}
			case 'L':
				if !strings.Contains(text, "using") && !strings.Contains(text, "system") {
					pp.UsePythonVersionLocal(CleanVersionString(text))
				}
			case 'D':
				if !strings.Contains(text, "using") && !strings.Contains(text, "system") {
					pp.UninstallPythonVersion(CleanVersionString(text), index, list)
				} else {
					setConfirmationContent("Can't remove this python version. Press enter to go back.",
						func() {
							App.SetFocus(pp.pythonsLocal)
						}, func() {
							App.SetFocus(pp.pythonsLocal)
						})
				}
			}
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
				final := detail

				for _, v := range pp.localPythons {
					cleanedVersion := CleanVersionString(v)
					if cleanedVersion == detail {
						final = fmt.Sprintf("%s *", detail)
						break
					}
				}

				detailNode := tview.NewTreeNode(final).SetColor(tview.Styles.PrimaryTextColor)

				detailNode.SetSelectedFunc(func() {
					if !strings.Contains(final, "*") {
						setConfirmationContent("Are you sure you want to install python version "+detail+"?",
							func() {
								App.SetFocus(pp.treeRemote)
							}, func() {
								pp.InstallPythonVersion(detail, detailNode)
								App.SetFocus(pp.treeRemote)
							})
					}
				})

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

func (pp *PythonPanel) InstallPythonVersion(identifier string, node *tview.TreeNode) {
	done := make(chan bool)
	originalText := node.GetText()

	go func() {
		rotation := []string{"|", "/", "-", "\\"} // Symbols for the spinner
		i := 0
		for {
			select {
			case <-done:
				node.SetText(fmt.Sprintf("%s *", originalText))
				node.SetSelectedFunc(func() {})
				pp.El.RemoveItem(pp.pythonsLocal)
				pp.localPythons = python.GetAvailPythonLocal()
				pp.initializeCurrLocalVersions()
				pp.El.AddItem(pp.pythonsLocal, 1, 0, 1, 1, 0, 0, false)
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
		_, stderr, err := python.InstallPythonVersion(identifier)
		if err != nil || stderr != "" {
			node.SetText(fmt.Sprintf("%s - Failed to install %s: %s", originalText, identifier, stderr))
		} else {
			node.SetText(fmt.Sprintf("%s - Installed %s successfully!", originalText, identifier))
		}
		done <- true
	}()
}

func (pp *PythonPanel) UninstallPythonVersion(identifier string, index int, list *tview.List) {
	// originalText, _ := list.GetItemText(index)

	go func() {
		_, stderr, err := python.UnInstallPythonVersion(identifier)

		if err != nil || stderr != "" {
			App.QueueUpdate(func() {
				pp.currPython.SetText("identifier : " + identifier)

				list.SetItemText(index, fmt.Sprintf("%s: %s", identifier, stderr), "")
			})
			App.Draw()
		} else {
			App.QueueUpdate(func() {
				list.RemoveItem(index)
			})
			App.Draw()
		}
	}()
}

func (pp *PythonPanel) UsePythonVersionLocal(identifier string) {
	_, err := exec.Command("pyenv", "local", identifier).Output()
	if err != nil {
		log.Fatalf("no local vercion %v", err)
	}

	pp.localPythons = python.GetAvailPythonLocal()

	pp.currLocal = python.GetPyenvLocal()
	pp.currGlobal = python.GetPyenvGlobal()

	pp.initializeCurrPython()
	pp.initializeCurrLocalVersions()
	App.SetFocus(pp.pythonsLocal)
}

func (pp *PythonPanel) UsePythonVersionGlobal(identifier string) {
	identifier = strings.TrimSpace(strings.ReplaceAll(identifier, "-> using", ""))

	_, err := exec.Command("pyenv", "global", identifier).Output()
	if err != nil {
		log.Fatalf("no global vercion %s %v", identifier, err)
	}

	pp.localPythons = python.GetAvailPythonLocal()

	pp.currLocal = python.GetPyenvLocal()
	pp.currGlobal = python.GetPyenvGlobal()

	pp.initializeCurrPython()
	pp.initializeCurrLocalVersions()
	App.SetFocus(pp.pythonsLocal)
}
