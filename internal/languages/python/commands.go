package python

import (
	"fmt"
	"runtime"
)

var versionCommand = map[string][]string{
	"darwin": {
		"python", "--version",
	},
	"linux": {
		"python", "--version",
	},
}

var pyenvVersion = map[string][]string{
	"darwin": {
		"bash", "-c", `
            pyenv --version
        `,
	},
	"linux": {
		"bash", "-c", `
            pyenv --version
        `,
	},
}

var pyenvListPythonVersions = map[string][]string{
	"darwin": {
		"bash", "-c", `
            pyenv versions
        `,
	},
	"linux": {
		"bash", "-c", `
            pyenv versions
        `,
	},
}

var pyenvInstallList = map[string][]string{
	"darwin": {
		"bash", "-c", `
            pyenv install -l
        `,
	},
	"linux": {
		"bash", "-c", `
            pyenv install -l
        `,
	},
}

var pyenvInstall = func(identifier string) map[string][]string {
	return map[string][]string{
		"darwin": {
			"bash", "-c", fmt.Sprintf(`
				pyenv install %s
			`, identifier),
		},
		"linux": {
			"bash", "-c", fmt.Sprintf(`
				pyenv install %s
			`, identifier),
		},
	}
}

var pyenvUse = func(identifier string) map[string][]string {
	return map[string][]string{
		"darwin": {
			"bash", "-c", fmt.Sprintf(`
				pyenv global %s
			`, identifier),
		},
		"linux": {
			"bash", "-c", fmt.Sprintf(`
				pyenv global %s
			`, identifier),
		},
	}
}

var pyenvUninstall = func(identifier string) map[string][]string {
	return map[string][]string{
		"darwin": {
			"bash", "-c", fmt.Sprintf(`
				pyenv uninstall -f %s
			`, identifier),
		},
		"linux": {
			"bash", "-c", fmt.Sprintf(`
				source "$HOME/.sdkman/bin/sdkman-init.sh"
				pyenv uninstall -f %s
			`, identifier),
		},
	}
}

var (
	OSversionCommand       = versionCommand[runtime.GOOS]
	OSpyenvVersion         = pyenvVersion[runtime.GOOS]
	OSpyenvListPython      = pyenvListPythonVersions[runtime.GOOS]
	OSpyenvInstallList     = pyenvInstallList[runtime.GOOS]
	OSpyenvInstallPython   = pyenvInstall
	OSpyenvUninstallPython = pyenvUninstall
)
