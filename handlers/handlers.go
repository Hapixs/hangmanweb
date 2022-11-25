package handlers

import "net/http"

const (
	templatePathIndex  = "static/templates/index.html"
	templatePathWin    = "static/templates/win.html"
	templatePathLoose  = "static/templates/loose.html"
	templatePathLogin  = "static/templates/login.html"
	templateStats      = "static/templates/statistics.html"
	templateScoreboard = "static/templates/scoreboard.html"
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

type ScoreboardHtmlData struct {
	SB_NAME  string
	SB_TYPE  string
	SB_USERS []string
}

func InitWebServer() {
	fs := http.FileServer(http.Dir("./static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/hangman", hangmanPostHandler)
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/reset", resetHandler)
	http.HandleFunc("/login", loginPostHandler)
	http.HandleFunc("/startsologame", startSoloPageHandler)
	http.HandleFunc("/nolog", annoLoginHandler)
	http.HandleFunc("/logout", disconnectHandler)
	http.HandleFunc("/restartsologame", restartSoloGameHandler)
	http.HandleFunc("/statistics", statisticsHandler)
	http.HandleFunc("/scoreboard", scoreboardHandler)
}
