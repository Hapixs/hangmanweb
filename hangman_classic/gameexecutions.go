package hangman_classic

import (
	"bufio"
	"os"
)

var executionCheckForRemainingTries = GameExecution{string(DefaultExecutionCheckForRemainingTries), func(userInput *string) bool {
	if GetGameTries() >= maxTries {
		DisplayLooseLogo()
		StopGame()
	}
	return false
}}

var executionLookForAutoSave = GameExecution{string(DefaultExecutionLookForAutoSave), func(userInput *string) bool {
	_, autoSaveStatus, _ := GetConfigItem(ConfigAutoSave)
	if autoSaveStatus {
		SaveGame()
	}
	return false
}}

var executionDisplayBody = GameExecution{string(DefaultExecutionDisplayBody), func(userInput *string) bool {
	DisplayBody()
	return false
}}

var executionWaitForInput = GameExecution{string(DefaultExecutionWaitForInput), func(userInput *string) bool {
	reader := bufio.NewReader(os.Stdin)
	in, _ := reader.ReadBytes(byte('\n'))
	if len(string(in)) <= 0 {
		return true
	}
	*userInput = string(in)
	return false
}}

var executionCheckForWord = GameExecution{string(DefaultExecutionCheckForWord), func(userInput *string) bool {
	if len(*userInput) > 1 {
		if string(*userInput) == "STOP" {
			SaveGame()
			StopGame()
		}
		if GetGameToFind() == *userInput {
			WinGame()
		}
		AddGameTry()
		AddGameTry()
		addInformationHeadMessage("This is not the correct word !")
		return true
	}

	return false
}}

var executionCheckForVowel = GameExecution{string(DefaultExecutionCheckForVowel), func(userInput *string) bool {
	rn := []rune(*userInput)[0]
	if gameMode == HARD && isVowel(rn) && VowelCount(GetGameUsed()) >= 3 {
		addInformationHeadMessage("You can't use vowel anymore !")
		AddGameTry()
		executionAddToUsedLetter.Func(userInput)
		return true
	}
	return false
}}

var executionCheckLetterIsUsed = GameExecution{string(DefaultExecutionLetterIsUsed), func(userInput *string) bool {
	rn := []rune(*userInput)[0]
	if HasOccurenceLetter(GetGameUsed(), rn) && gameMode != HARD {
		addInformationHeadMessage("You already use this letter")
		if gameMode == HARD {
			AddGameTry()
			if isVowel(rn) {
				AddGameTry()
			}
		}
		return true
	}
	return false
}}

var executionCheckForLetterOccurence = GameExecution{string(DefaultExecutionCheckForLetterOccurence), func(userInput *string) bool {
	rn := []rune(*userInput)[0]
	if !HasOccurenceLetter(GetGameToFind(), rune(rn)) {
		AddGameTry()
		addInformationHeadMessage(string(rn) + " is not in this word..")
	} else {
		SetGameWord(UpdateGameWord(GetGameToFind(), GetGameWord(), rn))
	}
	return false
}}

var executionCheckForWordDiscover = GameExecution{string(DefaultExecutionCheckForWordDiscover), func(userInput *string) bool {
	if !HasOccurenceLetter(GetGameWord(), '_') {
		WinGame()
		return true
	}
	return false
}}

var executionAddToUsedLetter = GameExecution{string(DefaultExecutionAddToUsedLetter), func(userInput *string) bool {
	rn := []rune(*userInput)[0]
	AddGameUsed(rn)
	return false
}}
