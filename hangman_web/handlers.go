package hangmanweb

import (
	"hangman_classic"
	"net/http"
	"text/template"
	"time"
)

type HtmlData struct {
	GetGameTries  int
	GetGameUsed   string
	GetGameWord   string
	GetGameToFind string
}

var LastHttp http.ResponseWriter
var LastRe *http.Request

func PostHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		if err := r.ParseForm(); err != nil {
			println("ParseForm() err: %v", err)
			return
		}
	}
	WebInputbuffer.Write([]byte(r.Form.Get("input")))
	http.Redirect(w, r, "/", http.StatusSeeOther)
	LastHttp = w
	LastRe = r
}

var IsWin = false
var IsLoose = false

var Game *hangman_classic.HangmanGame

func GetHandler(w http.ResponseWriter, r *http.Request) {
	if IsWin {
		tp := template.Must(template.ParseFiles("web/win.html"))
		tp.Execute(w, nil)
	} else if IsLoose {
		tp := template.Must(template.ParseFiles("web/loose.html"))
		tp.Execute(w, nil)
	} else {
		tp := template.Must(template.ParseFiles("web/index.html"))

		data := HtmlData{
			GetGameTries:  Game.GetGameTries(),
			GetGameUsed:   Game.GetGameUsed(),
			GetGameWord:   Game.GetGameWord(),
			GetGameToFind: Game.GetGameToFind(),
		}

		tp.Execute(w, data)
		LastHttp = w
		LastRe = r
	}
}

func ResetHandler(w http.ResponseWriter, r *http.Request) {
	IsWin = false
	IsLoose = false
	RestartHangman()
	time.Sleep(time.Second / 3)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
