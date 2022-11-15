package hangmanweb

import (
	"hangman_classic"
	"os"
)

var overridedExecutionWaitForInput = hangman_classic.GameExecution{Name: string(hangman_classic.DefaultExecutionWaitForInput), Func: func(userInput *string) bool {
	for WebInputbuffer.Len() <= 0 {
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

var overridedExecutionCheckForRemainingTries = hangman_classic.GameExecution{Name: string(hangman_classic.DefaultExecutionCheckForRemainingTries), Func: func(userInput *string) bool {
	if hangman_classic.GetGameTries() >= 10 {
		hangman_classic.DisplayLooseLogo()
		hangman_classic.StopGame()
	}
	return false
}}

var overridedExecutionCheckForWordDiscover = hangman_classic.GameExecution{Name: string(hangman_classic.DefaultExecutionCheckForWordDiscover), Func: func(userInput *string) bool {
	if !hangman_classic.HasOccurenceLetter(hangman_classic.GetGameWord(), '_') {
		hangman_classic.WinGame()
		return true
	}
	return false
}}
