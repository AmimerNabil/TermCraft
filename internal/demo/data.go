package demo

import (
	"TermCraft/internal/languages/java"
)

// Example JavaProperties object slice
var javaProps []java.JavaProperties = []java.JavaProperties{
	{
		JavaVendor:      "Oracle Corporation",
		JavaVersion:     "1.8.0_292",
		JavaHome:        "/usr/lib/jvm/java-8-oracle",
		JavaRuntimeName: "Java(TM) SE Runtime Environment",
	},
	{
		JavaVendor:      "AdoptOpenJDK",
		JavaVersion:     "11.0.10",
		JavaHome:        "/usr/lib/jvm/java-11-adoptopenjdk",
		JavaRuntimeName: "OpenJDK Runtime Environment",
	},
}

// Python properties
var pythonProps = []any{
	map[string]interface{}{
		"Vendor":         "Python Software Foundation",
		"Version":        "3.10.5",
		"InstallPath":    "/usr/local/bin/python3",
		"RuntimeName":    "CPython",
		"PackageManager": "pip",
		"VirtualEnv":     true,
	},
	map[string]interface{}{
		"Vendor":         "Anaconda, Inc.",
		"Version":        "3.9.7",
		"InstallPath":    "/opt/anaconda3/bin/python",
		"RuntimeName":    "Anaconda Python",
		"PackageManager": "conda",
		"VirtualEnv":     false,
	},
}

// Go properties
var goProps = []any{
	map[string]interface{}{
		"Vendor":         "Google",
		"Version":        "1.20.4",
		"InstallPath":    "/usr/local/go",
		"RuntimeName":    "Go Runtime",
		"PackageManager": "go get",
		"ModulesEnabled": true,
	},
	map[string]interface{}{
		"Vendor":         "GoLang (Open Source)",
		"Version":        "1.19.7",
		"InstallPath":    "/home/user/go",
		"RuntimeName":    "GoLang SDK",
		"PackageManager": "go mod",
		"ModulesEnabled": true,
	},
}

// Node.js properties
var nodeProps = []any{
	map[string]interface{}{
		"Vendor":         "OpenJS Foundation",
		"Version":        "18.9.0",
		"InstallPath":    "/usr/local/bin/node",
		"RuntimeName":    "Node.js",
		"PackageManager": "npm",
		"LTS":            false,
	},
	map[string]interface{}{
		"Vendor":         "NodeSource",
		"Version":        "16.14.2",
		"InstallPath":    "/opt/nodesource/bin/node",
		"RuntimeName":    "Node.js (LTS)",
		"PackageManager": "npm",
		"LTS":            true,
	},
}

// Rust properties
var rustProps = []any{
	map[string]interface{}{
		"Vendor":         "Mozilla",
		"Version":        "1.73.0",
		"InstallPath":    "/usr/local/cargo",
		"RuntimeName":    "Rust Compiler",
		"PackageManager": "cargo",
		"Stable":         true,
	},
	map[string]interface{}{
		"Vendor":         "Rust Foundation",
		"Version":        "1.72.1",
		"InstallPath":    "/home/user/.cargo",
		"RuntimeName":    "Cargo Build System",
		"PackageManager": "cargo",
		"Stable":         true,
	},
}

// C/C++ properties
var cppProps = []any{
	map[string]interface{}{
		"Vendor":      "GNU Project",
		"Version":     "11.3.0",
		"InstallPath": "/usr/local/bin/g++",
		"RuntimeName": "GCC",
		"BuildSystem": "make",
		"Standard":    "C++17",
	},
	map[string]interface{}{
		"Vendor":      "LLVM",
		"Version":     "14.0.6",
		"InstallPath": "/usr/local/llvm/bin/clang++",
		"RuntimeName": "Clang",
		"BuildSystem": "cmake",
		"Standard":    "C++20",
	},
}

// Kotlin properties
var kotlinProps = []any{
	map[string]interface{}{
		"Vendor":      "JetBrains",
		"Version":     "1.8.0",
		"InstallPath": "/opt/kotlin",
		"RuntimeName": "Kotlin Compiler",
		"BuildSystem": "Gradle",
		"JVMTarget":   "1.8",
	},
	map[string]interface{}{
		"Vendor":      "JetBrains",
		"Version":     "1.7.10",
		"InstallPath": "/usr/local/kotlin",
		"RuntimeName": "Kotlin SDK",
		"BuildSystem": "Maven",
		"JVMTarget":   "11",
	},
}
