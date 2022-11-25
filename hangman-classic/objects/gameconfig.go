package hangmanclassicobjects

import "errors"

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

type ConfigKey string

const (
	ConfigWordsList       ConfigKey = "wordsListFileName"
	ConfigHangmanFile     ConfigKey = "hangmanFileName"
	ConfigASCIIFile       ConfigKey = "asciiRuneFileName"
	ConfigGameMode        ConfigKey = "gameMode"
	ConfigMaxTries        ConfigKey = "maxTries"
	ConfigUseAscii        ConfigKey = "gameDesigneUseAscii"
	ConfigAutoClear       ConfigKey = "autoClear"
	ConfigHangmanHeight   ConfigKey = "hangmanDisplayHeight"
	ConfigASCIIHeight     ConfigKey = "asciiRuneDisplayHeight"
	ConfigAutoSave        ConfigKey = "autoSave"
	ConfigSaveFile        ConfigKey = "saveFileName"
	ConfigBetterTerminal  ConfigKey = "betterTerminal"
	ConfigMultipleWorkers ConfigKey = "usemultipleworkers"
	ConfigIsInit          ConfigKey = "isinit"
)

func (gc *GameConfig) InitConfig() {
	gc.IntItems = map[ConfigKey]ConfigItemInt{
		ConfigMaxTries:      {10, 10},
		ConfigGameMode:      {int(NORMAL), int(NORMAL)},
		ConfigHangmanHeight: {8, 8},
		ConfigASCIIHeight:   {9, 9},
	}

	gc.BoolItems = map[ConfigKey]ConfigItemBoolean{
		ConfigUseAscii:        {true, true},
		ConfigAutoClear:       {false, false},
		ConfigAutoSave:        {false, false},
		ConfigBetterTerminal:  {false, false},
		ConfigMultipleWorkers: {false, false},
		ConfigIsInit:          {false, false},
	}

	gc.StringItems = map[ConfigKey]ConfigItemString{
		ConfigWordsList:   {"", ""},
		ConfigHangmanFile: {"assets/hangman.txt", "assets/hangman.txt"},
		ConfigASCIIFile:   {"assets/standard.txt", "assets/standard.txt"},
		ConfigSaveFile:    {"saves/save.txt", "saves/save.txt"},
	}
}

func (gc *GameConfig) GetConfigItem(key ConfigKey) (int, bool, string) {
	if k, ok := gc.BoolItems[key]; ok {
		return 0, k.Value, ""
	} else if k, ok := gc.IntItems[key]; ok {
		return k.Value, false, ""
	} else if k, ok := gc.StringItems[key]; ok {
		return 0, false, k.Value
	}
	panic(errors.New("Unable to find " + string(key) + " config key !"))
}

func (gc *GameConfig) SetConfigItemValue(key ConfigKey, keyValue interface{}) {
	if k, ok := gc.BoolItems[key]; ok {
		k.Value = keyValue.(bool)
		gc.BoolItems[key] = k
		return
	} else if k, ok := gc.IntItems[key]; ok {
		k.Value = keyValue.(int)
		gc.IntItems[key] = k
		return
	} else if k, ok := gc.StringItems[key]; ok {
		k.Value = keyValue.(string)
		gc.StringItems[key] = k
		return
	}
	panic(errors.New("Unable to find " + string(key) + " config key !"))
}
