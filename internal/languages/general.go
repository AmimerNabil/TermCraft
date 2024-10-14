package languages

// / this is the most general functions that a language has to implement
type LanguagePack[L any, R any] interface {
	getLocalVersions() []L
	getRemoteVersions() []R
}

var SupportedLanguages []string = []string{
	"java",
	// "python",
	// "go",
	// "node",
	// "rust",
	// "c/c++",
	// "kotlin",
}

var defaultLang = "java"
