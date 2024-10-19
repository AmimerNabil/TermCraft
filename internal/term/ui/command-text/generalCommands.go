package commandtext

var appwide = `
  [::b]Available Commands[::-]

  [::u]App-Wide Commands[::-]
  [yellow] q[-]:                 Quit the help panel
  [yellow] ?[-]:                 Show commands
  [yellow] Ctrl + C[-]:          Quit the application
    `

var AvailLangCommands = appwide + `
  [::u]Language Panel Commands[::-]
  [yellow] ↓ (Down Arrow)[-]:    Move down
  [yellow] ↑ (Up Arrow)[-]:      Move up
  [yellow] Enter[-]:             Enter the main UI for the selected language
`
