package hangmanweb

import (
	"hangman_classic"
)

var overridedExecutionWaitForInput = hangman_classic.GameExecution{Name: string(hangman_classic.DefaultExecutionWaitForInput), Func: func(userInput *string, game *hangman_classic.HangmanGame) bool {
	Game := getWebGameFromId(game.PublicId)
	for Game.Input.Len() <= 0 {
		if game.Gamestatus == int(hangman_classic.ENDED) {
			return true
		}
	}
	in, _ := Game.Input.ReadString(byte('\n'))
	if len(string(in)) <= 0 {
		return true
	}
	*userInput = string(in)
	Game.Input.Reset()
	return false
}}

var overridedExecutionCheckForRemainingTries = hangman_classic.GameExecution{Name: string(hangman_classic.DefaultExecutionCheckForRemainingTries), Func: func(userInput *string, game *hangman_classic.HangmanGame) bool {
	if game == nil {
		return true
	}
	Game := getWebGameFromId(game.PublicId)
	if Game == nil {
		return true
	}
	if game.GetGameTries() >= 10 {
		Game.IsLoose = true
		return true
	}
	return false
}}

var overridedExecutionCheckForWordDiscover = hangman_classic.GameExecution{Name: string(hangman_classic.DefaultExecutionCheckForWordDiscover), Func: func(userInput *string, game *hangman_classic.HangmanGame) bool {
	Game := getWebGameFromId(game.PublicId)
	if !hangman_classic.HasOccurenceLetter(game.GetGameWord(), '_') {
		Game.IsWin = true
		return true
	}
	return false
}}

var overridedExecutionCheckForWord = hangman_classic.GameExecution{Name: string(hangman_classic.DefaultExecutionCheckForWord), Func: func(userInput *string, game *hangman_classic.HangmanGame) bool {
	Game := getWebGameFromId(game.PublicId)
	if len(*userInput) > 1 {
		if game.GetGameToFind() == *userInput {
			Game.IsWin = true
			return true
		}
		game.AddGameTry()
		game.AddGameTry()
		hangman_classic.AddInformationHeadMessage("This is not the correct word !")
		return true
	}

	return false
}}
