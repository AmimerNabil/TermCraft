package node

import (
	"fmt"
	"runtime"
)

// Command for getting the Node.js version
var nodeVersionCommand = map[string][]string{
	"darwin": {
		"node", "--version",
	},
	"linux": {
		"node", "--version",
	},
}

// Command for installing nvm
var nvmInstall = map[string][]string{
	"darwin": {
		"curl", "-o-", "https://raw.githubusercontent.com/nvm-sh/nvm/v0.39.4/install.sh",
	},
	"linux": {
		"curl", "-o-", "https://raw.githubusercontent.com/nvm-sh/nvm/v0.39.4/install.sh",
	},
}

// Command for getting the nvm version
var nvmVersion = map[string][]string{
	"darwin": {
		"bash", "-c", `
        	source "$HOME/.nvm/nvm.sh"
            nvm --version
        `,
	},
	"linux": {
		"bash", "-c", `
        	source "$HOME/.nvm/nvm.sh"
            nvm --version
        `,
	},
}

// Command for listing remote Node.js versions via nvm
var nvmListNodeVersions = map[string][]string{
	"darwin": {
		"bash", "-c", `
        	source "$HOME/.nvm/nvm.sh"
            nvm ls-remote
        `,
	},
	"linux": {
		"bash", "-c", `
        	source "$HOME/.nvm/nvm.sh"
            nvm ls-remote
        `,
	},
}

// Command for installing a specific Node.js version via nvm
var nvmInstallNodeVersion = func(identifier string) map[string][]string {
	return map[string][]string{
		"darwin": {
			"bash", "-c", fmt.Sprintf(`
				source "$HOME/.nvm/nvm.sh"
				nvm install %s
			`, identifier),
		},
		"linux": {
			"bash", "-c", fmt.Sprintf(`
				source "$HOME/.nvm/nvm.sh"
				nvm install %s
			`, identifier),
		},
	}
}

// Command for setting a specific Node.js version as default via nvm
var nvmSetNode = func(identifier string) map[string][]string {
	return map[string][]string{
		"darwin": {
			"bash", "-c", fmt.Sprintf(`
				fnm default %s
			`, identifier),
		},
		"linux": {
			"bash", "-c", fmt.Sprintf(`
				fnm default %s
			`, identifier),
		},
	}
}

// Command for uninstalling a specific Node.js version via nvm
var nvmUninstallNodeVersion = func(identifier string) map[string][]string {
	return map[string][]string{
		"darwin": {
			"bash", "-c", fmt.Sprintf(`
				source "$HOME/.nvm/nvm.sh"
				nvm uninstall %s
			`, identifier),
		},
		"linux": {
			"bash", "-c", fmt.Sprintf(`
				source "$HOME/.nvm/nvm.sh"
				nvm uninstall %s
			`, identifier),
		},
	}
}

// Variables based on the current OS
var (
	OSnodeVersionCommand  = nodeVersionCommand[runtime.GOOS]
	OSnvmVersion          = nvmVersion[runtime.GOOS]
	OSnvmInstall          = nvmInstall[runtime.GOOS]
	OSnvmListNodeVersions = nvmListNodeVersions[runtime.GOOS]
	OSnvmInstallNode      = nvmInstallNodeVersion
	OSnvmUninstallNode    = nvmUninstallNodeVersion
	OSnvmSetNode          = nvmSetNode
)
