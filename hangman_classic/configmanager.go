package hangman_classic

import "errors"

func (gc *GameConfig) InitConfig() {
	gc.IntItems = map[ConfigKey]ConfigItemInt{
		ConfigMaxTries:      {10, 10},
		ConfigGameMode:      {NORMAL, NORMAL},
		ConfigHangmanHeight: {8, 8},
		ConfigASCIIHeight:   {9, 9},
	}

	gc.BoolItems = map[ConfigKey]ConfigItemBoolean{
		ConfigUseAscii:       {true, true},
		ConfigAutoClear:      {false, false},
		ConfigAutoSave:       {false, false},
		ConfigBetterTerminal: {false, false},
	}

	gc.StringItems = map[ConfigKey]ConfigItemString{
		ConfigWordsList:   {"", ""},
		ConfigHangmanFile: {"hangman.txt", "hangman.txt"},
		ConfigASCIIFile:   {"standard.txt", "standard.txt"},
		ConfigSaveFile:    {"save.txt", "save.txt"},
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
