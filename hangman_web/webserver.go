package hangmanweb

import (
	"bytes"
	"hangman_classic"
	"net/http"
	"sync"
)

var WebInputbuffer = bytes.Buffer{}

func StartServer() {
	fs := http.FileServer(http.Dir("./web/"))
	http.HandleFunc("/hangman", PostHandler)
	http.Handle("/web/", http.StripPrefix("/web/", fs))
	http.HandleFunc("/", GetHandler)

	var wg sync.WaitGroup
	wg.Add(1)
	go WebWorkerFunc(&wg, 0)

	hangman_classic.SetConfigItemValue(hangman_classic.ConfigWordsList, "words.txt")
	hangman_classic.InitGame()
	hangman_classic.ReplaceExecution(overridedExecutionWaitForInput, string(hangman_classic.DefaultExecutionWaitForInput))
	hangman_classic.StartGame()
}

func WebWorkerFunc(wg *sync.WaitGroup, id int) {
	defer wg.Done()
	http.ListenAndServe(":8080", nil)
}
