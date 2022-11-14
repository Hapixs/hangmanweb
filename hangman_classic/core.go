package hangman_classic

import (
	"encoding/base64"
	"os"
	"strconv"
)

var maxTries, _, _ = GetConfigItem(configMaxTries)
var gameMode, _, _ = GetConfigItem(configGameMode)

var executions = []GameExecution{}

func StartGame() {
	args := os.Args[1:]

	GameProcessArguments(args)

	InitEnvironement()
	InitGameCache()
	InitUI()

	executions = append(executions, executionLookForAutoSave)
	executions = append(executions, executionDisplayBody)
	executions = append(executions, executionCheckForRemainingTries)
	executions = append(executions, executionWaitForInput)
	executions = append(executions, executionCheckForWord)
	executions = append(executions, executionCheckForVowel)
	executions = append(executions, executionCheckLetterIsUsed)
	executions = append(executions, executionCheckForLetterOccurence)
	executions = append(executions, executionCheckForWordDiscover)
	executions = append(executions, executionAddToUsedLetter)

	maxTries, _, _ = GetConfigItem(configMaxTries)
	gameMode, _, _ = GetConfigItem(configGameMode)
	addInformationHeadMessage("Good Luck, you have " + strconv.Itoa(maxTries-game.Tries) + "  attempts.")
	ContinueGame()
}

func ContinueGame() {
	for {
		userInput := ""
		for _, execution := range executions {
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
