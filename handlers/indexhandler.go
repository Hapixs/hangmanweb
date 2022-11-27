package handlers

import (
	"net/http"
	"objects"
	"text/template"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	var data HtmlData = HtmlData{}

	InitHtmlDataFragments(&data)

	var tp *template.Template
	if !objects.IsLogin(r) {
		tp = template.Must(template.ParseFiles(templatePathLogin))
		data.IsLogin = false
	} else {
		Game := objects.GetGameFromCookies(w, r)

		data.IsLogin = true

		if Game.Game == nil {
			tp = template.Must(template.ParseFiles(templatePathIndex))
			data.IsInGame = false
		} else {

			data.GameMode = Game.Gamemode
			data.GetGameToFind = Game.Game.GetGameToFind()
			data.GetGameTries = Game.Game.GetGameTries()
			data.GetGameUsed = Game.Game.GetGameUsed()
			data.GetGameWord = Game.Game.GetGameWord()
			data.GetGameToFind = Game.Game.GetGameToFind()
			data.GetUserName = Game.User.Username

			if Game.IsWin {
				tp = template.Must(template.ParseFiles(templatePathWin))
				Game.User.Wins++
				Game.User.WordsFind++
				Game.User.Points++
				Game.User.Played++
			} else if Game.IsLoose {
				tp = template.Must(template.ParseFiles(templatePathLoose))
				Game.User.Loose++
				Game.User.Played++
			} else {
				tp = template.Must(template.ParseFiles(templatePathIndex))
				data.IsInGame = true
			}
		}
	}
	tp.Execute(w, data)
}
