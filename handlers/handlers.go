package handlers

import (
	"net/http"
	"os"
)

const (
	templatePathIndex  = "static/templates/index.html"
	templatePathWin    = "static/templates/win.html"
	templatePathLoose  = "static/templates/loose.html"
	templatePathLogin  = "static/templates/login.html"
	templateStats      = "static/templates/statistics.html"
	templateScoreboard = "static/templates/scoreboard.html"

	restartButtonFragment = "restart_button"
	navigatorFragment     = "navigator"
)

type HtmlData struct {
	GetGameTries  int
	GetGameUsed   string
	GetGameWord   string
	GetGameToFind string
	IsInGame      bool
	GameMode      string

	IsLogin     bool
	GetUserName string
	Played      int
	Wins        int
	Loose       int
	Letters     int
	Words       int
	WinRatio    float32
	LooseRatio  float32

	NotifMessage string

	DifficultySelectorFragment string
	RestartButtonFragment      string
	NavigatorFragment          string

	SB_NAME       string
	SB_TYPE       string
	SB_USERS      []string
	SB_USER_PLACE int
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

func loadHtmlFragment(uri string) string {
	uri = "static/templates/fragments/" + uri + ".html"
	c, err := os.ReadFile(uri)

	if err != nil {
		println("Error when reading " + uri)
		return ""
	}
	return string(c)
}

func InitHtmlDataFragments(data *HtmlData) {
	data.RestartButtonFragment = loadHtmlFragment(restartButtonFragment)
	data.NavigatorFragment = loadHtmlFragment(navigatorFragment)
}
