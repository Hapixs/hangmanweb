package objects

import "hangmanclassicobjects"

var overridedExecutionWaitForInput = hangmanclassicobjects.GameExecution{Name: string(hangmanclassicobjects.DefaultExecutionWaitForInput), Func: func(userInput *string, game *hangmanclassicobjects.HangmanGame) bool {
	Game := getWebGameFromId(game.PublicId)
	for Game.Input.Len() <= 0 {
		if game.Gamestatus == int(hangmanclassicobjects.ENDED) {
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

var overridedExecutionCheckForRemainingTries = hangmanclassicobjects.GameExecution{Name: string(hangmanclassicobjects.DefaultExecutionCheckForRemainingTries), Func: func(userInput *string, game *hangmanclassicobjects.HangmanGame) bool {
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

var overridedExecutionCheckForWordDiscover = hangmanclassicobjects.GameExecution{Name: string(hangmanclassicobjects.DefaultExecutionCheckForWordDiscover), Func: func(userInput *string, game *hangmanclassicobjects.HangmanGame) bool {
	Game := getWebGameFromId(game.PublicId)
	if !hangmanclassicobjects.HasOccurenceLetter(game.GetGameWord(), '_') {
		Game.IsWin = true
		return true
	}
	return false
}}

var overridedExecutionCheckForWord = hangmanclassicobjects.GameExecution{Name: string(hangmanclassicobjects.DefaultExecutionCheckForWord), Func: func(userInput *string, game *hangmanclassicobjects.HangmanGame) bool {
	Game := getWebGameFromId(game.PublicId)
	if len(*userInput) > 1 {
		if game.GetGameToFind() == *userInput {
			Game.IsWin = true
			return true
		}
		game.AddGameTry()
		game.AddGameTry()
		hangmanclassicobjects.AddInformationHeadMessage("This is not the correct word !")
		return true
	}

	return false
}}

var overridedExecutionCheckForLetterOccurence = hangmanclassicobjects.GameExecution{Name: string(hangmanclassicobjects.DefaultExecutionCheckForLetterOccurence), Func: func(userInput *string, game *hangmanclassicobjects.HangmanGame) bool {
	Game := getWebGameFromId(game.PublicId)
	rn := []rune(*userInput)[0]
	if !hangmanclassicobjects.HasOccurenceLetter(game.GetGameToFind(), rune(rn)) {
		game.AddGameTry()
		hangmanclassicobjects.AddInformationHeadMessage(string(rn) + " is not in this word..")
	} else {
		game.SetGameWord(game.UpdateGameWord(game.GetGameToFind(), game.GetGameWord(), rn))
		Game.User.LetterFind += len(hangmanclassicobjects.GetOccurenceLetter(game.GetGameToFind(), rn))
	}
	return false
}}
