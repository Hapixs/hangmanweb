package hangmanweb

import (
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

func PostHandler(w http.ResponseWriter, r *http.Request) {
	Game := getGameFromCookies(w, r)
	switch r.Method {
	case "POST":
		if err := r.ParseForm(); err != nil {
			println("ParseForm() err: %v", err)
			return
		}
	}
	Game.Input.Write([]byte(r.Form.Get("input")))
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func GetHandler(w http.ResponseWriter, r *http.Request) {
	Game := getGameFromCookies(w, r)
	if Game.IsWin {
		tp := template.Must(template.ParseFiles("web/win.html"))
		tp.Execute(w, nil)
	} else if Game.IsLoose {
		tp := template.Must(template.ParseFiles("web/loose.html"))
		tp.Execute(w, nil)
	} else {
		tp := template.Must(template.ParseFiles("web/index.html"))

		data := HtmlData{
			GetGameTries:  Game.Game.GetGameTries(),
			GetGameUsed:   Game.Game.GetGameUsed(),
			GetGameWord:   Game.Game.GetGameWord(),
			GetGameToFind: Game.Game.GetGameToFind(),
		}
		tp.Execute(w, data)
	}
}

func ResetHandler(w http.ResponseWriter, r *http.Request) {
	Game := getGameFromCookies(w, r)
	Game.Game.Kill()
	sessions[Game.Game.PublicId] = nil
	time.Sleep(time.Second / 2)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
