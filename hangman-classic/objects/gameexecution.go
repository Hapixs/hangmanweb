package hangmanclassicobjects

import (
	"bufio"
	"os"
)

type GameExecution struct {
	Name string
	Func func(userInput *string, game *HangmanGame) bool
}

type DefaultExecution string

const (
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

func (g *HangmanGame) InitGameExecutions() {
	g.executions = append(g.executions, executionLookForAutoSave)
	g.executions = append(g.executions, executionDisplayBody)
	g.executions = append(g.executions, executionCheckForRemainingTries)
	g.executions = append(g.executions, executionWaitForInput)
	g.executions = append(g.executions, executionCheckForWord)
	g.executions = append(g.executions, executionCheckForVowel)
	g.executions = append(g.executions, executionCheckLetterIsUsed)
	g.executions = append(g.executions, executionCheckForLetterOccurence)
	g.executions = append(g.executions, executionCheckForWordDiscover)
	g.executions = append(g.executions, executionAddToUsedLetter)
}

func (g *HangmanGame) AddGameExecution(exec GameExecution) {
	g.executions = append(g.executions, exec)
}

func (g *HangmanGame) ReplaceExecution(exec GameExecution, targetName string) {
	for i, e := range g.executions {
		if e.Name == targetName {
			g.executions[i] = exec
			return
		}
	}
	println("Unable to find " + targetName + " execution in the list.")
}

func (g *HangmanGame) AddAfterExecution(exec GameExecution, targetName string) {
	for i, e := range g.executions {
		if e.Name == targetName {
			t := g.executions[i:]
			g.executions = append(append(g.executions[:i], exec), t...)
			return
		}
	}
	println("Unable to find " + targetName + " execution in the list.")
}

func (g *HangmanGame) AddBeforeExecution(exec GameExecution, targetName string) {
	for i, e := range g.executions {
		if e.Name == targetName {
			t := g.executions[i-1:]
			g.executions = append(append(g.executions[:i-1], exec), t...)
			return
		}
	}
	println("Unable to find " + targetName + " execution in the list.")
}

func (g *HangmanGame) RemoveExecution(targetName DefaultExecution) {
	for i, e := range g.executions {
		if e.Name == string(targetName) {
			g.executions = append(g.executions[:i], g.executions[i+1:]...)
			return
		}
	}
}

var executionCheckForRemainingTries = GameExecution{string(DefaultExecutionCheckForRemainingTries), func(userInput *string, game *HangmanGame) bool {
	maxTries, _, _ := game.Config.GetConfigItem(ConfigMaxTries)
	if game.GetGameTries() >= maxTries {
		game.DisplayLooseLogo()
	}
	return false
}}

var executionLookForAutoSave = GameExecution{string(DefaultExecutionLookForAutoSave), func(userInput *string, game *HangmanGame) bool {
	_, autoSaveStatus, _ := game.Config.GetConfigItem(ConfigAutoSave)
	if autoSaveStatus {
		game.SaveGame()
	}
	return false
}}

var executionDisplayBody = GameExecution{string(DefaultExecutionDisplayBody), func(userInput *string, game *HangmanGame) bool {
	game.DisplayBody()
	return false
}}

var executionWaitForInput = GameExecution{string(DefaultExecutionWaitForInput), func(userInput *string, game *HangmanGame) bool {
	reader := bufio.NewReader(os.Stdin)
	in, _ := reader.ReadBytes(byte('\n'))
	if len(string(in)) <= 0 {
		return true
	}
	*userInput = string(in)
	return false
}}

var executionCheckForWord = GameExecution{string(DefaultExecutionCheckForWord), func(userInput *string, game *HangmanGame) bool {
	if len(*userInput) > 1 {
		if string(*userInput) == "STOP" {
			game.SaveGame()
		}
		if game.GetGameToFind() == *userInput {
			game.WinGame()
		}
		game.AddGameTry()
		game.AddGameTry()
		AddInformationHeadMessage("This is not the correct word !")
		return true
	}

	return false
}}

var executionCheckForVowel = GameExecution{string(DefaultExecutionCheckForVowel), func(userInput *string, game *HangmanGame) bool {
	rn := []rune(*userInput)[0]
	gameMode, _, _ := game.Config.GetConfigItem(ConfigGameMode)
	if gameMode == int(HARD) && isVowel(rn) && VowelCount(game.GetGameUsed()) >= 3 {
		AddInformationHeadMessage("You can't use vowel anymore !")
		game.AddGameTry()
		executionAddToUsedLetter.Func(userInput, game)
		return true
	}
	return false
}}

var executionCheckLetterIsUsed = GameExecution{string(DefaultExecutionLetterIsUsed), func(userInput *string, game *HangmanGame) bool {
	rn := []rune(*userInput)[0]
	gameMode, _, _ := game.Config.GetConfigItem(ConfigGameMode)
	if HasOccurenceLetter(game.GetGameUsed(), rn) && gameMode != int(HARD) {
		AddInformationHeadMessage("You already use this letter")
		if gameMode == int(HARD) {
			game.AddGameTry()
			if isVowel(rn) {
				game.AddGameTry()
			}
		}
		return true
	}
	return false
}}

var executionCheckForLetterOccurence = GameExecution{string(DefaultExecutionCheckForLetterOccurence), func(userInput *string, game *HangmanGame) bool {
	rn := []rune(*userInput)[0]
	if !HasOccurenceLetter(game.GetGameToFind(), rune(rn)) {
		game.AddGameTry()
		AddInformationHeadMessage(string(rn) + " is not in this word..")
	} else {
		game.SetGameWord(game.UpdateGameWord(game.GetGameToFind(), game.GetGameWord(), rn))
	}
	return false
}}

var executionCheckForWordDiscover = GameExecution{string(DefaultExecutionCheckForWordDiscover), func(userInput *string, game *HangmanGame) bool {
	if !HasOccurenceLetter(game.GetGameWord(), '_') {
		game.WinGame()
		return true
	}
	return false
}}

var executionAddToUsedLetter = GameExecution{string(DefaultExecutionAddToUsedLetter), func(userInput *string, game *HangmanGame) bool {
	rn := []rune(*userInput)[0]
	game.AddGameUsed(rn)
	return false
}}
