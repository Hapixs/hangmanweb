package hangmanweb

import (
	"bytes"
	"hangman_classic"
	"net/http"
	"os"
	"text/template"
	"time"
)

type HtmlData struct {
	GetGameTries  int
	GetGameUsed   string
	GetGameWord   string
	GetGameToFind string
}

type WebGame struct {
	Game    *hangman_classic.HangmanGame
	Input   bytes.Buffer
	IsWin   bool
	IsLoose bool
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

var sessions = map[string](*WebGame){}

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

func getGameFromCookies(w http.ResponseWriter, r *http.Request) *WebGame {
	c, err := r.Cookie("sessionid")
	sessionid := ""

	if err != nil || sessions[c.Value] == nil {
		Game := &hangman_classic.HangmanGame{}
		os.Args = append(os.Args, "words.txt")
		Game.InitGame()
		Game.ReplaceExecution(overridedExecutionWaitForInput, string(hangman_classic.DefaultExecutionWaitForInput))
		Game.ReplaceExecution(overridedExecutionCheckForRemainingTries, string(hangman_classic.DefaultExecutionCheckForRemainingTries))
		Game.ReplaceExecution(overridedExecutionCheckForWordDiscover, string(hangman_classic.DefaultExecutionCheckForWordDiscover))
		Game.ReplaceExecution(overridedExecutionCheckForWord, string(hangman_classic.DefaultExecutionCheckForWord))
		Game.Config.SetConfigItemValue(hangman_classic.ConfigMultipleWorkers, true)
		sessions[Game.PublicId] = &WebGame{Game, bytes.Buffer{}, false, false}
		http.SetCookie(w, &http.Cookie{Name: "sessionid", Value: Game.PublicId})
		defer Game.StartGame()
		sessionid = Game.PublicId
	} else {
		sessionid = c.Value
	}

	return sessions[sessionid]
}

func getWebGameFromId(id string) *WebGame {
	return sessions[id]
}
