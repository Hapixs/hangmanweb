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

	data := HtmlData{}

	InitHtmlDataFragments(&data)

	data.GetUserName = user.Username
	data.Played = user.Played
	data.Wins = user.Wins
	data.Loose = user.Loose
	data.Letters = user.LetterFind
	data.Words = user.WordsFind
	data.WinRatio = user.GetWinRatio()
	data.LooseRatio = user.GetLooseRatio()

	tp.Execute(w, data)
}
