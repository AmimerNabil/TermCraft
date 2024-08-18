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

var (
	OSversionCommand = versionCommand[runtime.GOOS]
	OSfindCommand    = findCommands[runtime.GOOS]
)
