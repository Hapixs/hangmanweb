package hangman_classic

import (
	"os"
	"strconv"
)

var maxTries, _, _ = GetConfigItem(ConfigMaxTries)
var gameMode, _, _ = GetConfigItem(ConfigGameMode)
var isInit = false

func InitGame() {
	args := os.Args[1:]

	GameProcessArguments(args)

	InitEnvironement()
	InitGameCache()
	InitUI()
	InitGameExecutions()

	maxTries, _, _ = GetConfigItem(ConfigMaxTries)
	gameMode, _, _ = GetConfigItem(ConfigGameMode)
	addInformationHeadMessage("Good Luck, you have " + strconv.Itoa(maxTries-game.Tries) + "  attempts.")
	isInit = true
}

func StartGame() {
	if !isInit {
		InitGame()
	}
	for {
		userInput := ""
		for _, execution := range Executions {
			if execution.Func(&userInput) {
				break
			}
		}
	}
}

func WinGame() {
	DisplayWinLogo()
	StopGame()
}

func StopGame() {
	os.Exit(0)
}
