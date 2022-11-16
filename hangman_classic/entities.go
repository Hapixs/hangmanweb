package hangman_classic

type HangmanGame struct {
	Tries         int
	Word          string
	EncodedToFind string
	tofind        string
	Used          string
	Config        GameConfig

	Gamestatus int
	cache      Gamecache
	executions []GameExecution
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

type CommandFlag struct {
	FlagExecutor func(game *HangmanGame, args []string) []string
	Description  string
	Usage        string
	IsAliase     bool
	AliaseOf     string
}

type GameExecution struct {
	Name string
	Func func(userInput *string, game *HangmanGame) bool
}

const (
	NORMAL = 0
	HARD   = 1
)

const (
	PLAYING = 0
	ENDED   = 1
)

type ConfigKey string
type DefaultExecution string

const (
	ConfigWordsList      ConfigKey = "wordsListFileName"
	ConfigHangmanFile    ConfigKey = "hangmanFileName"
	ConfigASCIIFile      ConfigKey = "asciiRuneFileName"
	ConfigGameMode       ConfigKey = "gameMode"
	ConfigMaxTries       ConfigKey = "maxTries"
	ConfigUseAscii       ConfigKey = "gameDesigneUseAscii"
	ConfigAutoClear      ConfigKey = "autoClear"
	ConfigHangmanHeight  ConfigKey = "hangmanDisplayHeight"
	ConfigASCIIHeight    ConfigKey = "asciiRuneDisplayHeight"
	ConfigAutoSave       ConfigKey = "autoSave"
	ConfigSaveFile       ConfigKey = "saveFileName"
	ConfigBetterTerminal ConfigKey = "betterTerminal"

	DefaultExecutionCheckForRemainingTries  DefaultExecution = "checkforremainingtries"
	DefaultExecutionLookForAutoSave         DefaultExecution = "lookforautosave"
	DefaultExecutionDisplayBody             DefaultExecution = "displaybody"
	DefaultExecutionWaitForInput            DefaultExecution = "waitforinput"
	DefaultExecutionCheckForWord            DefaultExecution = "checkforword"
	DefaultExecutionCheckForVowel           DefaultExecution = "checkforvowel"
	DefaultExecutionLetterIsUsed            DefaultExecution = "letterisused"
	DefaultExecutionCheckForLetterOccurence DefaultExecution = "checkforletteroccurence"
	DefaultExecutionCheckForWordDiscover    DefaultExecution = "checkforworddiscover"
	DefaultExecutionAddToUsedLetter         DefaultExecution = "addtousedletter"
)
