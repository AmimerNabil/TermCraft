package java

import (
	"fmt"
	"runtime"
)

var versionCommand = map[string][]string{
	"darwin": {
		"java", "-XshowSettings:properties", "-version",
	},
	"linux": {
		"java", "-XshowSettings:properties", "-version",
	},
}

var sdkmanInstall = map[string][]string{
	"darwin": {
		"curl", "-s", "https://get.sdkman.io",
	},
	"linux": {
		"curl", "-s", "https://get.sdkman.io",
	},
}

var sdkmanVersion = map[string][]string{
	"darwin": {
		"bash", "-c", `
        	source "$HOME/.sdkman/bin/sdkman-init.sh"
            sdk version
        `,
	},
	"linux": {
		"bash", "-c", `
        	source "$HOME/.sdkman/bin/sdkman-init.sh"
            sdk version
        `,
	},
}

var sdkListJavaVersions = map[string][]string{
	"darwin": {
		"bash", "-c", `
        	source "$HOME/.sdkman/bin/sdkman-init.sh"
            sdk list java
        `,
	},
	"linux": {
		"bash", "-c", `
        	source "$HOME/.sdkman/bin/sdkman-init.sh"
            sdk list java
        `,
	},
}

var sdkInstallJavaVersion = func(identifier string) map[string][]string {
	return map[string][]string{
		"darwin": {
			"bash", "-c", fmt.Sprintf(`
				source "$HOME/.sdkman/bin/sdkman-init.sh"
				yes n | sdk install java %s
			`, identifier),
		},
		"linux": {
			"bash", "-c", fmt.Sprintf(`
				source "$HOME/.sdkman/bin/sdkman-init.sh"
				yes n | sdk install java %s
			`, identifier),
		},
	}
}

var sdkSetJava = func(identifier string) map[string][]string {
	return map[string][]string{
		"darwin": {
			"bash", "-c", fmt.Sprintf(`
				source "$HOME/.sdkman/bin/sdkman-init.sh"
				sdk default java %s
			`, identifier),
		},
		"linux": {
			"bash", "-c", fmt.Sprintf(`
				source "$HOME/.sdkman/bin/sdkman-init.sh"
				sdk default java %s
			`, identifier),
		},
	}
}

var sdkUninstallJavaVersion = func(identifier string) map[string][]string {
	return map[string][]string{
		"darwin": {
			"bash", "-c", fmt.Sprintf(`
				source "$HOME/.sdkman/bin/sdkman-init.sh"
				sdk uninstall java %s
			`, identifier),
		},
		"linux": {
			"bash", "-c", fmt.Sprintf(`
				source "$HOME/.sdkman/bin/sdkman-init.sh"
				sdk uninstall java %s
			`, identifier),
		},
	}
}

var (
	OSversionCommand = versionCommand[runtime.GOOS]
	OSsdkmanVersion  = sdkmanVersion[runtime.GOOS]
	OSsdkInstall     = sdkmanInstall[runtime.GOOS]
	OSsdkListJava    = sdkListJavaVersions[runtime.GOOS]
	OSinstallJava    = sdkInstallJavaVersion
	OSUninstallJava  = sdkUninstallJavaVersion
	OSSetJava        = sdkSetJava
)
