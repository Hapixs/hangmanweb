package hangmanweb

import (
	"hangman_classic"
	"net/http"
	"text/template"
)

type HtmlData struct {
	GetGameTries  int
	GetGameUsed   string
	GetGameWord   string
	GetGameToFind string
}

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
}

func GetHandler(w http.ResponseWriter, r *http.Request) {
	tp := template.Must(template.ParseFiles("web/index.html"))

	data := HtmlData{
		GetGameTries:  hangman_classic.GetGameTries(),
		GetGameUsed:   hangman_classic.GetGameUsed(),
		GetGameWord:   hangman_classic.GetGameWord(),
		GetGameToFind: hangman_classic.GetGameToFind(),
	}

	tp.Execute(w, data)
}

func ResetHandler(w http.ResponseWriter, r *http.Request) {
	RestartHangman()
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
