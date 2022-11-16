package hangmanweb

import (
	"hangman_classic"
	"net/http"
	"os"
)

var overridedExecutionWaitForInput = hangman_classic.GameExecution{Name: string(hangman_classic.DefaultExecutionWaitForInput), Func: func(userInput *string, game *hangman_classic.HangmanGame) bool {
	for WebInputbuffer.Len() <= 0 {
		if game.Gamestatus == hangman_classic.ENDED {
			return true
		}
	}
	in, _ := WebInputbuffer.ReadString(byte('\n'))
	if len(string(in)) <= 0 {
		return true
	}
	*userInput = string(in)
	os.Stdin.WriteString(in + "\n")
	WebInputbuffer.Reset()
	return false
}}

var overridedExecutionCheckForRemainingTries = hangman_classic.GameExecution{Name: string(hangman_classic.DefaultExecutionCheckForRemainingTries), Func: func(userInput *string, game *hangman_classic.HangmanGame) bool {
	if game.GetGameTries() >= 10 {
		IsLoose = true
		http.Redirect(LastHttp, LastRe, "/", http.StatusSeeOther)
		return true
	}
	return false
}}

var overridedExecutionCheckForWordDiscover = hangman_classic.GameExecution{Name: string(hangman_classic.DefaultExecutionCheckForWordDiscover), Func: func(userInput *string, game *hangman_classic.HangmanGame) bool {
	if !hangman_classic.HasOccurenceLetter(game.GetGameWord(), '_') {
		IsWin = true
		http.Redirect(LastHttp, LastRe, "/", http.StatusSeeOther)
		return true
	}
	return false
}}

var overridedExecutionCheckForWord = hangman_classic.GameExecution{Name: string(hangman_classic.DefaultExecutionCheckForWord), Func: func(userInput *string, game *hangman_classic.HangmanGame) bool {
	if len(*userInput) > 1 {
		if game.GetGameToFind() == *userInput {
			IsWin = true
			http.Redirect(LastHttp, LastRe, "/", http.StatusSeeOther)
			return true
		}
		game.AddGameTry()
		game.AddGameTry()
		hangman_classic.AddInformationHeadMessage("This is not the correct word !")
		return true
	}

	return false
}}
