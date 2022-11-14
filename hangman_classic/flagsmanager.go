package hangman_classic

import (
	"os"
	"sort"
)

var flagExecutors = map[string](CommandFlag){
	"-hm":        CommandFlag{flagHardModeExecutor, "Change to hard mode", "-hm", false, "-hm"},
	"--hardmode": CommandFlag{nil, "", "", true, "-hm"},
	"--hard":     CommandFlag{nil, "", "", true, "-hm"},

	"-sw":         CommandFlag{flagStartWithExecutor, "Start from a save file", "-sw <file>", false, "-sw"},
	"--usesave":   CommandFlag{nil, "", "", true, "-sw"},
	"--loadsave":  CommandFlag{nil, "", "", true, "-sw"},
	"--startWith": CommandFlag{nil, "", "", true, "-sw"},

	"-noas":     CommandFlag{flagNoASCII, "Don't use ascii art for letters", "-noas", false, "-noas"},
	"--noascii": CommandFlag{nil, "", "", true, "--noas"},

	"-as":     CommandFlag{flagUseASCII, "Use ascii art for letters", "-as", false, "-as"},
	"--ascii": CommandFlag{nil, "", "", true, "-as"},

	"-ac":         CommandFlag{flagAutoClear, "Auto clear the terminal after each actions", "-ac", false, "-ac"},
	"--clear":     CommandFlag{nil, "", "", true, "-ac"},
	"--autoclear": CommandFlag{nil, "", "", true, "-ac"},

	"-nac": CommandFlag{flagDontAutoClear, "Do not clear the terminal after each actions", "-nac", false, "-nac"},

	"-bh":          CommandFlag{flagBigHangman, "Use a big hangman ascii art (20x20)", "-", false, "-bh"},
	"--bighangman": CommandFlag{flagBigHangman, "", "", true, "-bh"},

	"-lh":             CommandFlag{flagLittleHangman, "Use a hagman ascii art (9x9)", "-lh", false, "-lh"},
	"--littlehangman": CommandFlag{flagLittleHangman, "", "", true, "-lh"},

	"-h":     CommandFlag{nil, "Display help menu", "-h", false, "-h"},
	"--help": CommandFlag{nil, "", "", true, "-h"},

	"-useascii": CommandFlag{flagUseASCIIWithBool, "Define if characters has to be displayed in ascii art", "-useascii [true/false]", false, "-useascii"},
	"-gamemode": CommandFlag{flagSetGameMode, "Define gamemode to use", "-gamemode [hard/normal]", false, "-gamemode"},

	"-autosave": CommandFlag{flagAutoSave, "Auto save the game after each actions", "-autosave", false, "-autosave"},
	"--save":    CommandFlag{nil, "", "", true, "-autosave"},
	"-savefile": CommandFlag{flagSaveFile, "Choose the of the save file.", "-savefile <filename>", false, "-savefile"},

	"--low":  CommandFlag{flagLow, "Use pre config for low quality", "--low", false, "--low"},
	"--high": CommandFlag{flagHigh, "Use pre config for high quality", "--high", false, "--high"},

	"-asciifile": CommandFlag{flagASCIIFile, "Specify the file used for translate character to ascii art", "-asciifile <file>", false, "-asciifile"},

	"-usebetterterm": CommandFlag{flagUseBetterTerm, "", "", false, "-usebetterterm"},
}

func flagHardModeExecutor(args []string) []string {
	SetConfigItemValue(ConfigGameMode, HARD)
	args = append(args[:0], args[1:]...)
	return args
}

func flagStartWithExecutor(args []string) []string {
	if len(args) <= 1 {
		println("[Warn] Please specify a file after --startWith !")
	} else {
		err := LoadSave(args[1])
		if err != nil {
			println("[Warn] Save file not found !")
		}
		FromSave = true
		args = append(args[:0], args[2:]...)
	}
	return args
}

func flagNoASCII(args []string) []string {
	SetConfigItemValue(ConfigUseAscii, false)
	args = append(args[:0], args[1:]...)
	return args
}

func flagUseASCII(args []string) []string {
	SetConfigItemValue(ConfigUseAscii, true)
	args = append(args[:0], args[1:]...)
	return args
}

func flagAutoClear(args []string) []string {
	SetConfigItemValue(ConfigAutoClear, true)
	args = append(args[:0], args[1:]...)
	return args
}

func flagDontAutoClear(args []string) []string {
	SetConfigItemValue(ConfigAutoClear, false)
	args = append(args[:0], args[1:]...)
	return args
}

func flagBigHangman(args []string) []string {
	SetConfigItemValue(ConfigHangmanFile, "bighangman.txt")
	SetConfigItemValue(ConfigHangmanHeight, 20)
	args = append(args[:0], args[1:]...)
	return args
}

func flagLittleHangman(args []string) []string {
	SetConfigItemValue(ConfigHangmanFile, "hangman.txt")
	SetConfigItemValue(ConfigHangmanHeight, 9)
	args = append(args[:0], args[1:]...)
	return args
}

func flagShowHelpMessage(args []string) []string {
	for _, l := range BuildFlagHelpMenu() {
		print(l)
	}
	os.Exit(0)
	return args
}

func flagUseASCIIWithBool(args []string) []string {
	if len(args) < 1 {
		args = append(args[:0], args[1:]...)
		return args
	}
	switch args[1] {
	case "false", "f", "0":
		SetConfigItemValue(ConfigUseAscii, true)
	case "true", "t", "1":
		SetConfigItemValue(ConfigUseAscii, false)
	}
	args = append(args[:0], args[2:]...)
	return args
}

func flagSetGameMode(args []string) []string {
	if len(args) < 1 {
		args = append(args[:0], args[1:]...)
		return args
	}
	switch args[1] {
	case "1", "hard", "h":
		SetConfigItemValue(ConfigGameMode, HARD)
	case "0", "normal", "n":
		SetConfigItemValue(ConfigGameMode, NORMAL)
	}
	args = append(args[:0], args[2:]...)
	return args
}

func flagAutoSave(args []string) []string {
	SetConfigItemValue(ConfigAutoSave, true)
	args = append(args[:0], args[1:]...)
	return args
}
func flagSaveFile(args []string) []string {
	if len(args) < 1 {
		args = append(args[:0], args[1:]...)
		return args
	}
	SetConfigItemValue(ConfigSaveFile, args[1])
	args = append(args[:0], args[2:]...)
	return args
}

func flagLow(args []string) []string {
	SetConfigItemValue(ConfigAutoClear, false)
	SetConfigItemValue(ConfigHangmanFile, "hangman.txt")
	SetConfigItemValue(ConfigHangmanHeight, 8)
	SetConfigItemValue(ConfigUseAscii, false)
	args = append(args[:0], args[1:]...)
	return args
}

func flagHigh(args []string) []string {
	SetConfigItemValue(ConfigAutoClear, true)
	SetConfigItemValue(ConfigHangmanFile, "bighangman.txt")
	SetConfigItemValue(ConfigHangmanHeight, 20)
	SetConfigItemValue(ConfigUseAscii, true)
	args = append(args[:0], args[1:]...)
	return args
}

func flagASCIIFile(args []string) []string {
	if len(args) <= 1 {
		println("[Warn] Please specify a file after -asciifile (using standard.txt instead)!")
		args = append(args[:0], args[1:]...)
	} else {
		SetConfigItemValue(ConfigASCIIFile, args[1])
		args = append(args[:0], args[2:]...)
	}
	return args
}

func flagUseBetterTerm(args []string) []string {
	SetConfigItemValue(ConfigBetterTerminal, true)
	args = append(args[:0], args[1:]...)
	return args
}

func GameProcessArguments(args []string) {
	for len(args) > 0 {
		arg := args[0]
		if arg[0] == "-"[0] {
			if arg == "--help" || arg == "-h" {
				flagShowHelpMessage(args)
				return
			}
			if val, ok := flagExecutors[arg]; ok {
				cmdFlag := val
				if cmdFlag.IsAliase {
					cmdFlag = flagExecutors[cmdFlag.AliaseOf]
				}
				args = cmdFlag.FlagExecutor(args)
			} else {
				args = append(args[:0], args[1:]...)
				println("Can't find argument " + arg)
			}
		} else {
			SetConfigItemValue(ConfigWordsList, arg)
			args = append(args[:0], args[1:]...)
		}
	}
}

type FlagHelpContent struct {
	Principal   string
	Description string
	Usage       string
	Aliases     []string
}

func BuildFlagHelpMenu() []string {
	flags := map[string](FlagHelpContent){}
	for k, v := range flagExecutors {
		if !v.IsAliase {
			flagHelpContent := FlagHelpContent{k, v.Description, v.Usage, []string{k}}
			flagHelpContent.Description = v.Description
			flags[flagHelpContent.Principal] = flagHelpContent
		}
	}
	for k, v := range flagExecutors {
		if v.IsAliase {
			if val, ok := flags[v.AliaseOf]; ok {
				val.Aliases = append(val.Aliases, k)
				flags[val.Principal] = val
			}
		}
	}
	flagList := []FlagHelpContent{}
	for _, v := range flags {
		flagList = append(flagList, v)
	}
	sort.SliceStable(flagList, func(i, j int) bool {
		return flagList[i].Principal < flagList[j].Principal
	})
	helpBlock := []string{}
	for _, v := range flagList {
		helpBlock = append(helpBlock, "\t"+v.Principal+"\n")
		helpBlock = append(helpBlock, "\t\tDescription: "+v.Description+"\n")
		helpBlock = append(helpBlock, "\t\tUsage: "+v.Usage+"\n")
		helpBlock = append(helpBlock, "\t\tAliases: ")
		argAliasesStr := ""
		sort.SliceStable(v.Aliases, func(i, j int) bool {
			return v.Aliases[i] < v.Aliases[j]
		})
		for i, a := range v.Aliases {
			argAliasesStr += a
			if len(v.Aliases)-1 > i {
				argAliasesStr += ", "
			}
		}
		helpBlock = append(helpBlock, argAliasesStr+"\n\n")
	}
	return helpBlock
}
