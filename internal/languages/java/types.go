package java

type JavaProperties struct {
	FileEncoding               string `json:"file.encoding"`
	FileSeparator              string `json:"file.separator"`
	JavaClassPath              string `json:"java.class.path"`
	JavaClassVersion           string `json:"java.class.version"`
	JavaHome                   string `json:"java.home" table:""`
	JavaIOTmpDir               string `json:"java.io.tmpdir"`
	JavaLibraryPath            string `json:"java.library.path"`
	JavaRuntimeName            string `json:"java.runtime.name" table:""`
	JavaRuntimeVersion         string `json:"java.runtime.version"`
	JavaSpecificationName      string `json:"java.specification.name"`
	JavaSpecificationVendor    string `json:"java.specification.vendor"`
	JavaSpecificationVersion   string `json:"java.specification.version"`
	JavaVendor                 string `json:"java.vendor" table:""`
	JavaVendorURL              string `json:"java.vendor.url"`
	JavaVendorURLBug           string `json:"java.vendor.url.bug"`
	JavaVendorVersion          string `json:"java.vendor.version"`
	JavaVersion                string `json:"java.version" table:""`
	JavaVersionDate            string `json:"java.version.date"`
	JavaVMCompressedOopsMode   string `json:"java.vm.compressedOopsMode"`
	JavaVMInfo                 string `json:"java.vm.info"`
	JavaVMName                 string `json:"java.vm.name"`
	JavaVMSpecificationName    string `json:"java.vm.specification.name"`
	JavaVMSpecificationVendor  string `json:"java.vm.specification.vendor"`
	JavaVMSpecificationVersion string `json:"java.vm.specification.version"`
	JavaVMVendor               string `json:"java.vm.vendor"`
	JavaVMVersion              string `json:"java.vm.version"`
	JdkDebug                   string `json:"jdk.debug"`
	LineSeparator              string `json:"line.separator"`
	NativeEncoding             string `json:"native.encoding"`
	OSArch                     string `json:"os.arch"`
	OSName                     string `json:"os.name"`
	OSVersion                  string `json:"os.version"`
	PathSeparator              string `json:"path.separator"`
	SunArchDataModel           string `json:"sun.arch.data.model"`
	SunBootLibraryPath         string `json:"sun.boot.library.path"`
	SunCPUEndian               string `json:"sun.cpu.endian"`
	SunIOUnicodeEncoding       string `json:"sun.io.unicode.encoding"`
	SunJavaLauncher            string `json:"sun.java.launcher"`
	SunJNUEncoding             string `json:"sun.jnu.encoding"`
	SunManagementCompiler      string `json:"sun.management.compiler"`
	SunStderrEncoding          string `json:"sun.stderr.encoding"`
	SunStdoutEncoding          string `json:"sun.stdout.encoding"`
	UserDir                    string `json:"user.dir"`
	UserHome                   string `json:"user.home"`
	UserLanguage               string `json:"user.language"`
	UserName                   string `json:"user.name"`
	CurrentlyActive            bool   `json:"java.active"`
}

type RemoteJavaProperties struct {
	JavaVendor  string `json:"java.vendor"`
	JavaVersion string `json:"java.version"`
	Identifier  string `json:"java.remote.identifier"`
	Installed   bool   `json:"java.installed"`
	InUse       bool   `json:"java.inUse"`
}
