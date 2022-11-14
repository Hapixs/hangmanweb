package hangman_classic

import "errors"

var configItemsInt = map[ConfigKey]ConfigItemInt{
	ConfigMaxTries:      {10, 10},
	ConfigGameMode:      {NORMAL, NORMAL},
	ConfigHangmanHeight: {8, 8},
	ConfigASCIIHeight:   {9, 9},
}

var configItemsBool = map[ConfigKey]ConfigItemBoolean{
	ConfigUseAscii:       {true, true},
	ConfigAutoClear:      {false, false},
	ConfigAutoSave:       {false, false},
	ConfigBetterTerminal: {false, false},
}

var configItemsString = map[ConfigKey]ConfigItemString{
	ConfigWordsList:   {"", ""},
	ConfigHangmanFile: {"hangman.txt", "hangman.txt"},
	ConfigASCIIFile:   {"standard.txt", "standard.txt"},
	ConfigSaveFile:    {"save.txt", "save.txt"},
}

func GetConfigItem(key ConfigKey) (int, bool, string) {
	if k, ok := configItemsBool[key]; ok {
		return 0, k.Value, ""
	} else if k, ok := configItemsInt[key]; ok {
		return k.Value, false, ""
	} else if k, ok := configItemsString[key]; ok {
		return 0, false, k.Value
	}
	panic(errors.New("Unable to find " + string(key) + " config key !"))
}

func SetConfigItemValue(key ConfigKey, keyValue interface{}) {
	if k, ok := configItemsBool[key]; ok {
		k.Value = keyValue.(bool)
		configItemsBool[key] = k
		return
	} else if k, ok := configItemsInt[key]; ok {
		k.Value = keyValue.(int)
		configItemsInt[key] = k
		return
	} else if k, ok := configItemsString[key]; ok {
		k.Value = keyValue.(string)
		configItemsString[key] = k
		return
	}
	panic(errors.New("Unable to find " + string(key) + " config key !"))
}

func GetGameConfig() GameConfig {
	return GameConfig{configItemsBool, configItemsInt, configItemsString}
}

func SetGameConfig(config GameConfig) {
	configItemsInt = config.IntItems
	configItemsBool = config.BoolItems
	configItemsString = config.StringItems
}
