package handlers

import (
	"net/http"
	"objects"
	"text/template"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	var data HtmlData
	var tp *template.Template
	if !objects.IsLogin(r) {
		tp = template.Must(template.ParseFiles(templatePathLogin))
	} else {
		Game := objects.GetGameFromCookies(w, r)
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
