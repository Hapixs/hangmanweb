package main

import (
	"net/http"
	"strconv"
	"text/template"

	"github.com/Hapixs/hangmanclassic"
)

type HtmlData struct {
	GetGameTries  int
	GetGameUsed   string
	GetGameWord   string
	GetGameToFind string
	GetUserName   string
	IsInGame      bool
	GameMode      string
}

const (
	templatePathIndex = "static/templates/index.html"
	templatePathWin   = "static/templates/win.html"
	templatePathLoose = "static/templates/loose.html"
	templatePathLogin = "static/templates/login.html"
	templateStats     = "static/templates/statistics.html"
)

func HangmanPostHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		if err := r.ParseForm(); err != nil {
			println("ParseForm() err: %v", err)
			return
		}
		Game := getGameFromCookies(w, r)
		Game.Input.Write([]byte(r.Form.Get("input")))
	}
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

		encodedPass := string(hangmanclassic.GetEncodedStringInSha256(r.Form.Get("password")))

		for _, v := range usermap {
			if v.Username == r.Form.Get("username") {
				if v.Password == encodedPass {
					v.SetUpUserCookies(&w)
					http.Redirect(w, r, "/", http.StatusSeeOther)
					return
				}
				http.Redirect(w, r, "/", http.StatusSeeOther)
				return
			}
		}

		user := &User{Username: r.Form.Get("username"), Password: encodedPass, isAnnonyme: false, Points: 0, Wins: 0, Loose: 0, Played: 0, LetterFind: 0, WordsFind: 0}
		user.GenerateUniqueId()
		user.SetUpUserCookies(&w)
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func AnnoLoginHandler(w http.ResponseWriter, r *http.Request) {
	ano := User{isAnnonyme: true, Username: "Annonyme"}
	ano.GenerateUniqueId()
	ano.SetUpUserCookies(&w)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func DisconnectHandler(w http.ResponseWriter, r *http.Request) {
	c := http.Cookie{Name: "user_id", Value: ""}
	http.SetCookie(w, &c)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func StartSoloPageHandler(w http.ResponseWriter, r *http.Request) {
	if !IsLogin(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	gameMode := r.FormValue("difficulty")
	StartNewGame(&w, r, gameMode)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func RestartSoloGameHandler(w http.ResponseWriter, r *http.Request) {
	if !IsLogin(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	Game := getGameFromCookies(w, r)
	if Game != nil && Game.Game != nil {
		mutex.Lock()
		sessions[Game.Game.PublicId] = nil
		mutex.Unlock()
	}
	gameMode := r.FormValue("difficulty")
	StartNewGame(&w, r, gameMode)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	var data HtmlData
	var tp *template.Template
	if !IsLogin(r) {
		tp = template.Must(template.ParseFiles(templatePathLogin))
	} else {
		Game := getGameFromCookies(w, r)
		if Game.Game == nil {
			tp = template.Must(template.ParseFiles(templatePathIndex))
			data = HtmlData{
				GetUserName: Game.User.Username,
				IsInGame:    false,
			}
		} else if Game.IsWin {
			tp = template.Must(template.ParseFiles(templatePathWin))
			Game.User.Wins++
			Game.User.WordsFind++
			Game.User.Points++
			Game.User.Played++
			data = HtmlData{
				GameMode:      Game.Gamemode,
				GetGameToFind: Game.Game.GetGameToFind(),
			}
		} else if Game.IsLoose {
			tp = template.Must(template.ParseFiles(templatePathLoose))
			Game.User.Loose++
			Game.User.Played++
			data = HtmlData{
				GameMode:      Game.Gamemode,
				GetGameToFind: Game.Game.GetGameToFind(),
			}
		} else {
			tp = template.Must(template.ParseFiles(templatePathIndex))
			data = HtmlData{
				GetGameTries:  Game.Game.GetGameTries(),
				GetGameUsed:   Game.Game.GetGameUsed(),
				GetGameWord:   Game.Game.GetGameWord(),
				GetGameToFind: Game.Game.GetGameToFind(),
				GetUserName:   Game.User.Username,
				IsInGame:      true,
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
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

type StatsHtmlData struct {
	Username   string
	Played     int
	Wins       int
	Loose      int
	Letters    int
	Words      int
	WinRatio   float32
	LooseRatio float32
}

func StatisticsHandler(w http.ResponseWriter, r *http.Request) {
	if !IsLogin(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	user, _ := GetUserFromRequest(r)
	tp := template.Must(template.ParseFiles(templateStats))

	data := StatsHtmlData{
		Username:   user.Username,
		Played:     user.Played,
		Wins:       user.Wins,
		Loose:      user.Loose,
		Letters:    user.LetterFind,
		Words:      user.WordsFind,
		WinRatio:   user.GetWinRatio(),
		LooseRatio: user.GetLooseRatio(),
	}

	tp.Execute(w, data)
}

type ScoreboardHtmlData struct {
	SB_NAME  string
	SB_TYPE  string
	SB_USERS []string
}

func ScoreboardHandler(w http.ResponseWriter, r *http.Request) {
	sbt := ScoreboardType_Letter

	if r.Method == "POST" {
		sbt = r.FormValue("sb_type")
	}

	tp := template.Must(template.ParseFiles("web/scoreboard.html"))

	sb := BuildScoreboard(sbt)
	sbd := ScoreboardHtmlData{}
	sbd.SB_NAME = sb.Name
	sbd.SB_TYPE = sbt

	for i, v := range sb.Top {
		switch sbt {
		case ScoreboardType_Letter:
			sbd.SB_USERS = append(sbd.SB_USERS, "#"+strconv.Itoa(i+1)+": "+v.Username+" - "+strconv.Itoa(v.LetterFind))
		case ScoreboardType_Words:
			sbd.SB_USERS = append(sbd.SB_USERS, "#"+strconv.Itoa(i+1)+": "+v.Username+" - "+strconv.Itoa(v.WordsFind))
		case ScoreboardType_Points:
			sbd.SB_USERS = append(sbd.SB_USERS, "#"+strconv.Itoa(i+1)+": "+v.Username+" - "+strconv.Itoa(v.Points))
		case ScoreboardType_loose:
			sbd.SB_USERS = append(sbd.SB_USERS, "#"+strconv.Itoa(i+1)+": "+v.Username+" - "+strconv.Itoa(v.Loose))
		case ScoreboardType_win:
			sbd.SB_USERS = append(sbd.SB_USERS, "#"+strconv.Itoa(i+1)+": "+v.Username+" - "+strconv.Itoa(v.Wins))
		}
	}

	tp.Execute(w, sbd)
}
