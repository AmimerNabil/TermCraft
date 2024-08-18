package java

import (
	"TermCraft/internal/term/commands"
	"fmt"
	"log"
	"runtime"
	"strings"
)

func GetLocalRunningJava() (string, error) {
	command := commands.TerminalCommand{
		Command: OSversionCommand[0], Args: OSversionCommand[1:],
	}

	_, errOut, err := command.Run()
	if err != nil {
		log.Panicln(err)
	}

	return errOut, nil
}

func GetLocalJavaVersions() []JavaProperties {
	// get executables of java in computer
	executables, err := GetLocalJavaExecutables()
	if err != nil {
		log.Panicln(err)
	}

	currVersion, err := GetLocalRunningJava()
	if err != nil {
		log.Panicln(err)
	}

	var versions []JavaProperties

	for _, v := range executables {

		command := commands.TerminalCommand{
			Command: v, Args: OSversionCommand[1:],
		}
		_, output, err := command.Run()
		if err != nil {
			fmt.Println(err)
			continue
		}

		version := parseProperties(output)
		version.CurrentlyActive = currVersion == output
		versions = append(versions, version)
	}

	return versions
}

func GetLocalJavaExecutables() ([]string, error) {
	commandToRun, ok := findCommands[runtime.GOOS]
	if !ok {
		log.Panic("unsupported OS")
	}

	command := commands.TerminalCommand{
		Command: commandToRun[0],
		Args:    commandToRun[1:],
	}

	stdOut, _, error := command.Run()
	if error != nil {
		return nil, error
	}
	javaLocations := strings.Split(stdOut, "\n")
	var finalLocations []string

	for _, v := range javaLocations {
		trimmed := strings.TrimSpace(v)
		if trimmed != "" && strings.Contains(trimmed, "bin/java") {
			finalLocations = append(finalLocations, trimmed)
		}
	}

	return finalLocations, nil
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
