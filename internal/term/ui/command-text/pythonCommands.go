package commandtext

var pythonLanguageList = appwide + `
  [::u]Local Version List Commands[::-]
  [yellow] ↓ (Down Arrow)[-]:    Move down
  [yellow] ↑ (Up Arrow)[-]:      Move up
  [yellow] L[-]:                 Select the version to be used locally 
  [yellow] G[-]:                 Select the version to be used globally
  [yellow] Shift + D[-]:         Delete the selected language version 
`

var PythonPanel = pythonLanguageList + `
  [::u]Remote Versions Commands[::-]
  [yellow] ↓ (Down Arrow)[-]:    Move down
  [yellow] ↑ (Up Arrow)[-]:      Move up
  [yellow] Enter[-]:             Open | Install
`
