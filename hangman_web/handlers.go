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
	GetUserName   string
}

const (
	templatePathIndex = "web/index.html"
	templatePathWin   = "web/win.html"
	templatePathLoose = "web/loose.html"
	templatePathLogin = "web/login.html"
)

func HangmanPostHandler(w http.ResponseWriter, r *http.Request) {
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

func LoginPostHandler(w http.ResponseWriter, r *http.Request) {
	if !IsLogin(r) {
		switch r.Method {
		case "POST":
			if err := r.ParseForm(); err != nil {
				println("ParseForm() err: %v", err)
				return
			}
		}

		encodedPass := string(hangman_classic.GetEncodedStringInSha256(r.Form.Get("password")))

		for _, v := range usermap {
			if v.Username == r.Form.Get("username") {
				if v.Password == encodedPass {
					v.SetUpUserCookies(&w)
					return
				}
				// TODO: send error
				return
			}
		}

		user := &User{Username: r.Form.Get("username"), Password: encodedPass}
		println(r.Form.Get("username"))
		user.GenerateUniqueId()
		user.SetUpUserCookies(&w)
		println("Registered " + user.Username)
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	var data HtmlData
	var tp *template.Template
	if !IsLogin(r) {
		tp = template.Must(template.ParseFiles(templatePathLogin))
	} else {
		Game := getGameFromCookies(w, r)
		if Game.IsWin {
			tp = template.Must(template.ParseFiles(templatePathWin))
		} else if Game.IsLoose {
			tp = template.Must(template.ParseFiles(templatePathLoose))
		} else {
			tp = template.Must(template.ParseFiles(templatePathIndex))
			data = HtmlData{
				GetGameTries:  Game.Game.GetGameTries(),
				GetGameUsed:   Game.Game.GetGameUsed(),
				GetGameWord:   Game.Game.GetGameWord(),
				GetGameToFind: Game.Game.GetGameToFind(),
				GetUserName:   Game.User.Username,
			}
		}
	}
	tp.Execute(w, data)
}

func ResetHandler(w http.ResponseWriter, r *http.Request) {
	Game := getGameFromCookies(w, r)
	Game.Game.Kill()
	mutex.Lock()
	sessions[Game.Game.PublicId] = nil
	mutex.Unlock()
	time.Sleep(time.Second / 2)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
