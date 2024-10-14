package java

import (
	"TermCraft/internal/term/commands"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

// public
func GetAllJavaVersionInformation(identifier string) JavaProperties {
	var pathToExec string
	if identifier == "java" {
		pathToExec = "java"
	} else {
		pathToExec = "/home/izelhl/.sdkman/candidates/java" + identifier + "/bin/java"
	}

	command := commands.TerminalCommand{
		Command: pathToExec, Args: OSversionCommand[1:],
	}

	_, output, _ := command.Run()

	// if err != nil {
	// 	// TODO: do something here fmt.Println(err)
	// }

	version := parseProperties(output)

	return version
}

func IsSDKMANInstalled() bool {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error getting user home directory:", err)
		return false
	}

	sdkmanDir := homeDir + "/.sdkman"
	if _, err := os.Stat(sdkmanDir); !os.IsNotExist(err) {
		return true
	}

	command := commands.TerminalCommand{
		Command: OSsdkmanVersion[0],
		Args:    OSsdkmanVersion[1:],
	}

	_, _, error := command.Run()
	return error == nil
}

func InstallSdkMan() (string, error) {
	command := commands.TerminalCommand{
		Command: OSsdkInstall[0],
		Args:    OSsdkInstall[1:],
	}

	script, _, err := command.Run() // what I get here is the sdkman script
	if err != nil {
		fmt.Println("problem fetching script", err)
		return "", err
	}

	command = commands.TerminalCommand{
		Command: "bash",
		Args: []string{
			"-c", script,
		},
	}

	_, _, errr := command.Run()
	if errr != nil {
		fmt.Println("problem exec script", errr)
		return "", errr
	}

	return "successfully Installed", nil
}

func GetRemoteVersions() []RemoteJavaProperties {
	var rv []RemoteJavaProperties

	command := commands.TerminalCommand{
		Command: OSsdkListJava[0],
		Args:    OSsdkListJava[1:],
	}

	versionsSTDO, _, err := command.Run()
	if err != nil {
		fmt.Println("problem fetching versions", err)
		// TODO: handle panic
		log.Panic(err)
	}

	parseJavaOutput(versionsSTDO, &rv)

	return rv
}

func GetLocalJavaVersionsSdk() []RemoteJavaProperties {
	versions := GetRemoteVersions()

	var out []RemoteJavaProperties

	for _, v := range versions {
		if v.Installed {
			out = append(out, v)
		}
	}

	return out
}

// private
func parseJavaOutput(output string, rv *[]RemoteJavaProperties) {
	lines := strings.Split(output, "\n")
	var currentVendor string

	// Regular expression to match Java version lines with or without vendor
	re := regexp.MustCompile(`^\s*([A-Za-z]+)?\s*\|\s*(>>>)?\s*\|\s+(\S+)\s+\|\s+(\S+)\s+\|\s*(installed|local only|)\s*\|\s+(\S+)`)

	for _, line := range lines {
		matches := re.FindStringSubmatch(line)
		if matches != nil {
			vendor := matches[1]
			if vendor == "" {
				vendor = currentVendor
			} else {
				currentVendor = vendor
			}

			version := matches[3]
			status := matches[5]
			identifier := matches[6]

			// Determine if the version is installed
			installed := matches[2] == ">>>" || status == "installed" || status == "local only"

			javaProperties := RemoteJavaProperties{
				JavaVendor:  vendor,
				JavaVersion: version,
				Identifier:  identifier,
				Installed:   installed,
			}

			*rv = append(*rv, javaProperties)
		}
	}
}

func parseProperties(output string) JavaProperties {
	properties := JavaProperties{}
	lines := strings.Split(output, "\n")

	for _, line := range lines {
		if strings.TrimSpace(line) == "" || !strings.Contains(line, "=") {
			continue
		}
		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {
			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])

			switch key {
			case "file.encoding":
				properties.FileEncoding = value
			case "file.separator":
				properties.FileSeparator = value
			case "java.class.path":
				properties.JavaClassPath = value
			case "java.class.version":
				properties.JavaClassVersion = value
			case "java.home":
				properties.JavaHome = value
			case "java.io.tmpdir":
				properties.JavaIOTmpDir = value
			case "java.library.path":
				properties.JavaLibraryPath = value
			case "java.runtime.name":
				properties.JavaRuntimeName = value
			case "java.runtime.version":
				properties.JavaRuntimeVersion = value
			case "java.specification.name":
				properties.JavaSpecificationName = value
			case "java.specification.vendor":
				properties.JavaSpecificationVendor = value
			case "java.specification.version":
				properties.JavaSpecificationVersion = value
			case "java.vendor":
				properties.JavaVendor = value
			case "java.vendor.url":
				properties.JavaVendorURL = value
			case "java.vendor.url.bug":
				properties.JavaVendorURLBug = value
			case "java.vendor.version":
				properties.JavaVendorVersion = value
			case "java.version":
				properties.JavaVersion = value
			case "java.version.date":
				properties.JavaVersionDate = value
			case "java.vm.compressedOopsMode":
				properties.JavaVMCompressedOopsMode = value
			case "java.vm.info":
				properties.JavaVMInfo = value
			case "java.vm.name":
				properties.JavaVMName = value
			case "java.vm.specification.name":
				properties.JavaVMSpecificationName = value
			case "java.vm.specification.vendor":
				properties.JavaVMSpecificationVendor = value
			case "java.vm.specification.version":
				properties.JavaVMSpecificationVersion = value
			case "java.vm.vendor":
				properties.JavaVMVendor = value
			case "java.vm.version":
				properties.JavaVMVersion = value
			case "jdk.debug":
				properties.JdkDebug = value
			case "line.separator":
				properties.LineSeparator = value
			case "native.encoding":
				properties.NativeEncoding = value
			case "os.arch":
				properties.OSArch = value
			case "os.name":
				properties.OSName = value
			case "os.version":
				properties.OSVersion = value
			case "path.separator":
				properties.PathSeparator = value
			case "sun.arch.data.model":
				properties.SunArchDataModel = value
			case "sun.boot.library.path":
				properties.SunBootLibraryPath = value
			case "sun.cpu.endian":
				properties.SunCPUEndian = value
			case "sun.io.unicode.encoding":
				properties.SunIOUnicodeEncoding = value
			case "sun.java.launcher":
				properties.SunJavaLauncher = value
			case "sun.jnu.encoding":
				properties.SunJNUEncoding = value
			case "sun.management.compiler":
				properties.SunManagementCompiler = value
			case "sun.stderr.encoding":
				properties.SunStderrEncoding = value
			case "sun.stdout.encoding":
				properties.SunStdoutEncoding = value
			case "user.dir":
				properties.UserDir = value
			case "user.home":
				properties.UserHome = value
			case "user.language":
				properties.UserLanguage = value
			case "user.name":
				properties.UserName = value
			}
		}
	}
	return properties
}
