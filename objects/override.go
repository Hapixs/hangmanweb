package objects

import "github.com/Hapixs/hangmanclassic"

var overridedExecutionWaitForInput = hangmanclassic.GameExecution{Name: string(hangmanclassic.DefaultExecutionWaitForInput), Func: func(userInput *string, game *hangmanclassic.HangmanGame) bool {
	Game := getWebGameFromId(game.PublicId)
	for Game.Input.Len() <= 0 {
		if game.Gamestatus == int(hangmanclassic.ENDED) {
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

var overridedExecutionCheckForRemainingTries = hangmanclassic.GameExecution{Name: string(hangmanclassic.DefaultExecutionCheckForRemainingTries), Func: func(userInput *string, game *hangmanclassic.HangmanGame) bool {
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

var overridedExecutionCheckForWordDiscover = hangmanclassic.GameExecution{Name: string(hangmanclassic.DefaultExecutionCheckForWordDiscover), Func: func(userInput *string, game *hangmanclassic.HangmanGame) bool {
	Game := getWebGameFromId(game.PublicId)
	if !hangmanclassic.HasOccurenceLetter(game.GetGameWord(), '_') {
		Game.IsWin = true
		return true
	}
	return false
}}

var overridedExecutionCheckForWord = hangmanclassic.GameExecution{Name: string(hangmanclassic.DefaultExecutionCheckForWord), Func: func(userInput *string, game *hangmanclassic.HangmanGame) bool {
	Game := getWebGameFromId(game.PublicId)
	if len(*userInput) > 1 {
		if game.GetGameToFind() == *userInput {
			Game.IsWin = true
			return true
		}
		game.AddGameTry()
		game.AddGameTry()
		hangmanclassic.AddInformationHeadMessage("This is not the correct word !")
		return true
	}

	return false
}}

var overridedExecutionCheckForLetterOccurence = hangmanclassic.GameExecution{Name: string(hangmanclassic.DefaultExecutionCheckForLetterOccurence), Func: func(userInput *string, game *hangmanclassic.HangmanGame) bool {
	Game := getWebGameFromId(game.PublicId)
	rn := []rune(*userInput)[0]
	if !hangmanclassic.HasOccurenceLetter(game.GetGameToFind(), rune(rn)) {
		game.AddGameTry()
		hangmanclassic.AddInformationHeadMessage(string(rn) + " is not in this word..")
	} else {
		game.SetGameWord(game.UpdateGameWord(game.GetGameToFind(), game.GetGameWord(), rn))
		Game.User.LetterFind += len(hangmanclassic.GetOccurenceLetter(game.GetGameToFind(), rn))
	}
	return false
}}
