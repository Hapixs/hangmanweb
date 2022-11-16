package hangmanweb

import (
	"bytes"
	"hangman_classic"
	"net/http"
	"os"
	"sync"
)

var WebInputbuffer = bytes.Buffer{}

func StartServer() {
	InitWebHandlers()
	StartWorkers()
	http.ListenAndServe(":8080", nil)
}

func InitWebHandlers() {
	http.HandleFunc("/hangman", PostHandler)
	http.HandleFunc("/", GetHandler)
	http.HandleFunc("/reset", ResetHandler)
}

var webgroup sync.WaitGroup

func StartWorkers() {
	webgroup.Add(1)
	go WebWorkerFunc(&webgroup, 0)
}

func StartHangmanClassic() {
	Game = &hangman_classic.HangmanGame{}
	os.Args = append(os.Args, "words.txt")
	Game.InitGame()
	Game.ReplaceExecution(overridedExecutionWaitForInput, string(hangman_classic.DefaultExecutionWaitForInput))
	Game.ReplaceExecution(overridedExecutionCheckForRemainingTries, string(hangman_classic.DefaultExecutionCheckForRemainingTries))
	Game.ReplaceExecution(overridedExecutionCheckForWordDiscover, string(hangman_classic.DefaultExecutionCheckForWordDiscover))
	Game.ReplaceExecution(overridedExecutionCheckForWord, string(hangman_classic.DefaultExecutionCheckForWord))
	Game.Config.SetConfigItemValue(hangman_classic.ConfigMultipleWorkers, true)
	Game.StartGame()
}

func RestartHangman() {
	Game.Kill()
	StartWorkers()
}

func WebWorkerFunc(wg *sync.WaitGroup, id int) {
	defer wg.Done()
	StartHangmanClassic()
}
