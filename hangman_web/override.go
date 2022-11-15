package hangmanweb

import (
	"hangman_classic"
	"os"
)

var overridedExecutionWaitForInput = hangman_classic.GameExecution{Func: func(userInput *string) bool {
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
