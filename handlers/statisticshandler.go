package handlers

import (
	"net/http"
	"objects"
	"text/template"
)

func statisticsHandler(w http.ResponseWriter, r *http.Request) {
	if !objects.IsLogin(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	user, _ := objects.GetUserFromRequest(r)
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
