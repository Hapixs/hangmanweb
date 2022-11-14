package hangman_classic

type HangmanGame struct {
	Tries         int
	Word          string
	EncodedToFind string
	tofind        string
	Used          string
}

type ConfigItemBoolean struct {
	Value   bool
	Default bool
}
type ConfigItemInt struct {
	Value   int
	Default int
}

type ConfigItemString struct {
	Value   string
	Default string
}

type GameConfig struct {
	BoolItems   map[ConfigKey]ConfigItemBoolean
	IntItems    map[ConfigKey]ConfigItemInt
	StringItems map[ConfigKey]ConfigItemString
}

type GameSave struct {
	Game   HangmanGame
	Config GameConfig
}

type CommandFlag struct {
	FlagExecutor func(args []string) []string
	Description  string
	Usage        string
	IsAliase     bool
	AliaseOf     string
}

type GameExecution struct {
	Func func(userInput *string) bool
}

const (
	NORMAL = 0
	HARD   = 1
)

type ConfigKey string

const (
	configWordsList      ConfigKey = "wordsListFileName"
	configHangmanFile    ConfigKey = "hangmanFileName"
	configASCIIFile      ConfigKey = "asciiRuneFileName"
	configGameMode       ConfigKey = "gameMode"
	configMaxTries       ConfigKey = "maxTries"
	configUseAscii       ConfigKey = "gameDesigneUseAscii"
	configAutoClear      ConfigKey = "autoClear"
	configHangmanHeight  ConfigKey = "hangmanDisplayHeight"
	configASCIIHeight    ConfigKey = "asciiRuneDisplayHeight"
	configAutoSave       ConfigKey = "autoSave"
	configSaveFile       ConfigKey = "saveFileName"
	configBetterTerminal ConfigKey = "betterTerminal"
)
