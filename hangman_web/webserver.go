package hangmanweb

import (
	"bytes"
	"hangman_classic"
	"net/http"
	"sync"
)

var WebInputbuffer = bytes.Buffer{}

func StartServer() {
	InitWebHandlers()
	StartWorkers()
	StartHangmanClassic()
}

func InitWebHandlers() {
	http.HandleFunc("/hangman", PostHandler)
	http.HandleFunc("/", GetHandler)
	http.HandleFunc("/reset", ResetHandler)
}

var webgroup sync.WaitGroup
var hangmangroup sync.WaitGroup

func StartWorkers() {
	webgroup.Add(1)
	go WebWorkerFunc(&webgroup, 0)
}

func StartHangmanClassic() {
	hangman_classic.SetConfigItemValue(hangman_classic.ConfigWordsList, "words.txt")
	hangman_classic.InitGame()
	hangman_classic.ReplaceExecution(overridedExecutionWaitForInput, string(hangman_classic.DefaultExecutionWaitForInput))
	hangman_classic.ReplaceExecution(overridedExecutionCheckForRemainingTries, string(hangman_classic.DefaultExecutionCheckForRemainingTries))
	hangman_classic.ReplaceExecution(overridedExecutionCheckForWordDiscover, string(hangman_classic.DefaultExecutionCheckForWordDiscover))
	hangman_classic.StartGame()
}

func RestartHangman() {

	hangman_classic.SetConfigItemValue(hangman_classic.ConfigWordsList, "words.txt")
	hangman_classic.InitGame()
	hangman_classic.ReplaceExecution(overridedExecutionWaitForInput, string(hangman_classic.DefaultExecutionWaitForInput))
	hangman_classic.ReplaceExecution(overridedExecutionCheckForRemainingTries, string(hangman_classic.DefaultExecutionCheckForRemainingTries))
	hangman_classic.ReplaceExecution(overridedExecutionCheckForWordDiscover, string(hangman_classic.DefaultExecutionCheckForWordDiscover))
	hangman_classic.DisplayBody()
}

func WebWorkerFunc(wg *sync.WaitGroup, id int) {
	defer wg.Done()
	http.ListenAndServe(":8080", nil)
}
