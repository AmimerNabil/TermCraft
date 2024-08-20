package java

import "runtime"

var findCommands = map[string][]string{
	"darwin": {
		"find", "/usr", "/home", "/opt",
		"-path", "/usr/sbin/authserver", "-prune", "-o",
		"-type", "f", "-perm", "+111",
		"-name", "java",
		"-print",
	},
	"linux": {
		"find", "/usr", "/home",
		"-path", "/home/root", "-prune", "-o",
		"-type", "f", "-executable",
		"-name", "java", "-o", "-name", "binjava",
	},
}

var versionCommand = map[string][]string{
	"darwin": {
		"java", "-XshowSettings:properties", "-version",
	},
	"linux": {
		"java", "-XshowSettings:properties", "-version",
	},
}

// $ curl -s "https://get.sdkman.io" | bash
var sdkmanInstall = map[string][]string{
	"darwin": {
		"curl", "-s", "https://get.sdkman.io", "|", "bash",
	},
	"linux": {
		"curl", "-s", "https://get.sdkman.io",
	},
}

var sdkmanVersion = map[string][]string{
	"darwin": {
		"sdk", "version",
	},
	"linux": {
		"sdk", "version",
	},
}

var (
	OSversionCommand = versionCommand[runtime.GOOS]
	OSfindCommand    = findCommands[runtime.GOOS]
	OSsdkmanVersion  = sdkmanVersion[runtime.GOOS]
	OSsdkInstall     = sdkmanInstall[runtime.GOOS]
)
