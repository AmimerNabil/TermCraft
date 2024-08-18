package java

import "runtime"

var findCommands = map[string][]string{
	"darwin": {
		"find", "/usr", "/home", "/Library", "/opt",
		"-path", "/usr/sbin/authserver", "-prune", "-o",
		"-path", "/Library/Caches", "-prune", "-o",
		"-path", "/Library/Trial", "-prune", "-o",
		"-path", "/Library/Application Support", "-prune", "-o",
		"-type", "f", "-perm", "+111",
		"-name", "java", "-o",
		"-name", "binjava",
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
