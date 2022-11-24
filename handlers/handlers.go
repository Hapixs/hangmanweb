package handlers

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
