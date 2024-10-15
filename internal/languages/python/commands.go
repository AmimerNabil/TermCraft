package python

import "runtime"

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

var pyenvLocal = map[string][]string{
	"darwin": {
		"bash", "-c", `
            pyenv local
        `,
	},
	"linux": {
		"bash", "-c", `
            pyenv local
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

var (
	OSpyenvLocal       = pyenvLocal[runtime.GOOS]
	OSversionCommand   = versionCommand[runtime.GOOS]
	OSpyenvVersion     = pyenvVersion[runtime.GOOS]
	OSpyenvListPython  = pyenvListPythonVersions[runtime.GOOS]
	OSpyenvInstallList = pyenvInstallList[runtime.GOOS]
)
