package hangman_classic

import (
	"encoding/base64"
	"os"
	"strconv"
)

var maxTries, _, _ = GetConfigItem(ConfigMaxTries)
var gameMode, _, _ = GetConfigItem(ConfigGameMode)

var Executions = []GameExecution{}

var isInit = false

func InitGame() {
	args := os.Args[1:]

	GameProcessArguments(args)

	InitEnvironement()
	InitGameCache()
	InitUI()

	Executions = append(Executions, executionLookForAutoSave)
	Executions = append(Executions, executionDisplayBody)
	Executions = append(Executions, executionCheckForRemainingTries)
	Executions = append(Executions, executionWaitForInput)
	Executions = append(Executions, executionCheckForWord)
	Executions = append(Executions, executionCheckForVowel)
	Executions = append(Executions, executionCheckLetterIsUsed)
	Executions = append(Executions, executionCheckForLetterOccurence)
	Executions = append(Executions, executionCheckForWordDiscover)
	Executions = append(Executions, executionAddToUsedLetter)

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

func EncodeStrInBase64(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}

func DecodeStrInBase64(str string) string {
	decoded, _ := base64.StdEncoding.DecodeString(str)
	return string(decoded)
}
