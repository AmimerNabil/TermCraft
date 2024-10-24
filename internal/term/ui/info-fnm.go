package ui

import (
	"TermCraft/internal/languages/node"
	commandtext "TermCraft/internal/term/ui/command-text"
	"strings"

	"github.com/gdamore/tcell/v2"
)

type FNMPanel struct {
	*Panel

	currFNMinfo string
}

func NewFNMPanel() *FNMPanel {
	fp := &FNMPanel{
		Panel: &Panel{},
	}

	fp.i(commandtext.JavaPanel)
	fp.init()

	return fp
}

func (fp *FNMPanel) init() {
	// #1: setup info
	fp.currFNMinfo = node.GetVerboseNodeInfo()
	fp.initCurrVersionInfo(fp.currFNMinfo)

	// #2: get local installed versions
	localVersions, _, _ := node.ListInstalledNodeVersions()

	fp.localVersions = localVersions
	fp.initCurrVersions(fp.localVersions)

	// #3: get remote available options
	remoteVersions, _ := node.GetRemoteNodeVersions()
	fp.initRemoteVersions(fp.updateLocal, remoteVersions, node.InstallNodeVersion)

	fp.currVersionsInstalled.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		index := fp.currVersionsInstalled.GetCurrentItem()
		var text string

		if len(fp.localVersions) > 0 && index >= 0 {
			itemText, _ := fp.currVersionsInstalled.GetItemText(index)
			text = itemText
		}

		cleanText := nodecleanVersionString(text)

		switch event.Key() {
		case tcell.KeyTab:
			App.SetFocus(fp.remoteVersionsAvailable)
			return nil
		case tcell.KeyEscape:
			App.SetFocus(AvailableLanguesSections.El)
		case tcell.KeyRune:
			switch event.Rune() {

			// TODO: figure out how to do local node correctly.
			// case 'L':
			// 	node.SetLocalNodeVersion(cleanText)
			// 	fp.init()
			// 	App.SetFocus(fp.currVersionsInstalled)

			case 'G':
				node.SetNodeVersion(cleanText)
				fp.init()
				App.SetFocus(fp.currVersionsInstalled)
			case 'D':
				if !strings.Contains(text, "using") && !strings.Contains(text, "system") {
					fp.UninstallPythonVersion(node.UninstallNodeVersion, cleanText, index, fp.currVersionsInstalled)
					fp.init()
					App.SetFocus(fp.currVersionsInstalled)
				} else {
					setConfirmationContent("Can't remove this python version. Press enter to go back.",
						func() {
							App.SetFocus(fp.currVersionsInstalled)
						}, func() {
							App.SetFocus(fp.currVersionsInstalled)
						})
				}
			}
		}

		return event
	})
}

func (fp *FNMPanel) updateLocal() []string {
	localVersions, _, _ := node.ListInstalledNodeVersions()
	return localVersions
}

func nodecleanVersionString(version string) string {
	if idx := strings.Index(version, "id:"); idx != -1 {
		idPart := strings.TrimSpace(version[idx+3:])
		idPart = strings.ReplaceAll(idPart, "-> using", "")
		return strings.TrimSpace(idPart)
	}
	return version
}
