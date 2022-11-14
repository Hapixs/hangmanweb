package hangmanweb

import (
	"bufio"
	"fmt"
	"hangman_classic"
	"net/http"
	"os"
	"sync"
	"text/template"
)

func StartServer() {
	fs := http.FileServer(http.Dir("./web/"))
	http.HandleFunc("/hangman", PostHandler)
	http.Handle("/web/", http.StripPrefix("/web/", fs))
	http.HandleFunc("/", GetHandler)
	var wg sync.WaitGroup
	wg.Add(1)
	go worker(&wg, 0)
	hangman_classic.SetConfigItemValue(hangman_classic.ConfigWordsList, "words.txt")
	hangman_classic.InitGame()
	hangman_classic.Executions[3] = overridedExecutionWaitForInput
	hangman_classic.StartGame()
}

func GetHandler(w http.ResponseWriter, r *http.Request) {
	tp := template.Must(template.ParseFiles("web/index.html"))
	tp.Execute(w, nil)
}

var tempInput = ""

func PostHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		fmt.Println("GET")
	case "POST":
		if err := r.ParseForm(); err != nil {
			println("ParseForm() err: %v", err)
			return
		}
	}
	os.Stdin.Write([]byte(r.Form.Get("input") + "\n"))
}

func worker(wg *sync.WaitGroup, id int) {
	defer wg.Done()
	http.ListenAndServe(":8080", nil)
}

var overridedExecutionWaitForInput = hangman_classic.GameExecution{Func: func(userInput *string) bool {
	reader := bufio.NewReader(os.Stdin)
	in, _ := reader.ReadString(byte('\n'))
	println("readed " + in)
	if len(string(in)) <= 0 {
		return true
	}
	*userInput = string(in)
	return false
}}
