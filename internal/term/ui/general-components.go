package ui

import (
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type Panel struct {
	container *tview.Grid

	currVersionView         *tview.TextView
	currVersionsInstalled   *tview.List
	remoteVersionsAvailable *tview.TreeView

	currVersionHolder          *tview.Flex
	currVersionInstalledHolder *tview.Flex
	remoteVersionHolder        *tview.Flex

	localVersions []string
}

func (pp *Panel) updateLocal(update func() []string) {
	pp.localVersions = update()

	pp.currVersionsInstalled.Clear()

	for i, versions := range pp.localVersions {
		pp.currVersionsInstalled.AddItem(versions, "", rune('a'+i), nil)
	}
}

func (pp *Panel) i(commands string) {
	pp.container = tview.NewGrid()
	pp.container.SetRows(14, 0)
	pp.container.SetColumns(-1, -1)

	pp.container.
		SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
			switch event.Key() {
			case tcell.KeyRune:
				switch event.Rune() {
				case '?':
					commandText.SetText(commands)
					commandsPages.ShowPage("Command")
				}
			}
			return event
		})
}

func (pp *Panel) initCurrVersionInfo(info string) {
	if pp.currVersionView != nil {
		pp.container.RemoveItem(pp.currVersionView)
		pp.container.RemoveItem(pp.currVersionHolder)
		pp.currVersionView = nil
		pp.currVersionHolder = nil
	}

	pp.currVersionView = tview.NewTextView()
	pp.currVersionView.SetText(info).
		SetDynamicColors(true). // Enable dynamic color tags
		SetTextAlign(tview.AlignLeft).
		SetWrap(true)

	flexPaddingH := tview.NewFlex().
		AddItem(nil, 5, 1, false).
		AddItem(pp.currVersionView, 0, 1, true).
		AddItem(nil, 5, 1, false)

	pp.currVersionHolder = tview.NewFlex().SetDirection(tview.FlexColumnCSS).
		AddItem(nil, 1, 1, false).
		AddItem(flexPaddingH, 0, 1, true)

	pp.currVersionHolder.SetBorder(true).SetTitle("Currently Used").SetTitleColor(tview.Styles.TertiaryTextColor)

	pp.container.AddItem(pp.currVersionHolder, 0, 0, 1, 2, 0, 0, false)
}

func (pp *Panel) initCurrVersions(versions []string) {
	if pp.currVersionsInstalled != nil {
		pp.container.RemoveItem(pp.currVersionsInstalled)
		pp.container.RemoveItem(pp.currVersionInstalledHolder)
		pp.currVersionsInstalled = nil
		pp.currVersionInstalledHolder = nil
	}

	list := tview.NewList().ShowSecondaryText(false)
	for i, versions := range versions {
		list.AddItem(versions, "", rune('a'+i), nil)
	}

	flex := tview.NewFlex().
		AddItem(nil, 1, 1, false). // Add 2-unit padding to the left
		AddItem(list, 0, 1, true).
		AddItem(nil, 5, 1, false) // Add 2-unit padding to the right

	pp.currVersionInstalledHolder = tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(nil, 2, 1, false). // Add 2-unit padding to the left
		AddItem(flex, 0, 1, true).
		AddItem(nil, 1, 1, false) // Add 2-unit padding to the right

	pp.currVersionInstalledHolder.SetBorder(true).SetTitle("Currently Installed Versions").SetTitleColor(
		tview.Styles.TertiaryTextColor,
	)

	pp.currVersionsInstalled = list
	pp.container.AddItem(pp.currVersionInstalledHolder, 1, 0, 1, 1, 0, 0, false)
}

func (pp *Panel) initRemoteVersions(updateLocal func() []string, versions map[string]map[string][]string, installFunc func(identifier string) (string, string, error)) {
	vendors := make([]string, 0, len(versions))
	for vendor := range versions {
		vendors = append(vendors, vendor)
	}
	sort.Strings(vendors)

	rootNode := tview.NewTreeNode("Remote Versions")
	for _, vendor := range vendors {
		vendorNode := tview.NewTreeNode(vendor).SetColor(tview.Styles.SecondaryTextColor).SetExpanded(false)
		versionsMap := versions[vendor]

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

				for _, v := range pp.localVersions {
					if v == detail {
						final = fmt.Sprintf("%s *", detail)
						break
					}
				}

				detailNode := tview.NewTreeNode(final).SetColor(tview.Styles.PrimaryTextColor)

				detailNode.SetSelectedFunc(func() {
					if !strings.Contains(final, "*") {
						setConfirmationContent("Are you sure you want to install version "+detail+"?",
							func() {
								App.SetFocus(pp.remoteVersionsAvailable)
							}, func() {
								pp.installRemoteVersion(updateLocal, installFunc, detail, detailNode)
								App.SetFocus(pp.remoteVersionsAvailable)
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
			App.SetFocus(pp.currVersionsInstalled)
			return nil
		}
		return event
	})

	if pp.remoteVersionsAvailable != nil {
		pp.container.RemoveItem(pp.remoteVersionHolder)
		pp.remoteVersionHolder = nil
		pp.remoteVersionHolder = nil
	}

	pp.remoteVersionHolder = tview.NewFlex().
		AddItem(nil, 1, 1, false). // Add 2-unit padding to the left
		AddItem(treeView, 0, 1, true).
		AddItem(nil, 1, 1, false) // Add 2-unit padding to the right

	pp.remoteVersionHolder.SetTitle("Versions Available for download").SetBorder(true).SetTitleColor(tview.Styles.TertiaryTextColor)

	pp.remoteVersionsAvailable = treeView
	pp.container.AddItem(pp.remoteVersionHolder, 1, 1, 1, 1, 0, 0, false)
}

func (pp *Panel) installRemoteVersion(updateLocal func() []string, installerFunc func(identifier string) (string, string, error), identifier string, node *tview.TreeNode) {
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

				pp.updateLocal(updateLocal)

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
		_, stderr, err := installerFunc(strings.TrimSpace(identifier))
		if err != nil || stderr != "" {
			// errtext = stderr
			node.SetText(fmt.Sprintf("%s - Failed to install %s: %s", originalText, identifier, stderr))
		} else {
			node.SetText(fmt.Sprintf("%s - Installed %s successfully!", originalText, identifier))
		}
		done <- true
	}()
}

func (pp *Panel) UninstallPythonVersion(uninstallFunc func(identifier string) (string, string, error), identifier string, index int, list *tview.List) {
	go func() {
		_, stderr, err := uninstallFunc(identifier)
		if err != nil || stderr != "" {
			App.QueueUpdate(func() {
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
