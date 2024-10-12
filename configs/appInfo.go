package configs

var (
	AppName    string = "TermCraft"
	AppVersion string = "0.0.1"
	AppLogo    string = `
 ____  ____  ____  _  _   ___  ____   __   ____  ____
(_  _)(  __)(  _ \( \/ ) / __)(  _ \ / _\ (  __)(_  _)
  )(   ) _)  )   // \/ \( (__  )   //    \ ) _)   )(  
 (__) (____)(__\_)\_)(_/ \___)(__\_)\_/\_/(__)   (__)
`
	AppDescription string = `
        Termcraft is a terminal-based text UI application designed for developers 
        to easily manage their configuration files and language installations. 
        It streamlines the setup and customization process, providing a unified, 
        efficient interface directly within the terminal.
    `
	SupportedOS []string = []string{
		"darwin", "linux",
	}
)
