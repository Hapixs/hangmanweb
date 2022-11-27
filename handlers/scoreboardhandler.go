package handlers

import (
	"net/http"
	"objects"
	"strconv"
	"text/template"
)

func scoreboardHandler(w http.ResponseWriter, r *http.Request) {
	sbt := objects.ScoreboardType_Letter

	if r.Method == "POST" {
		sbt = r.FormValue("sb_type")
	}

	tp := template.Must(template.ParseFiles(templateScoreboard))

	sb := objects.BuildScoreboard(sbt)
	sbd := HtmlData{}
	InitHtmlDataFragments(&sbd)

	sbd.SB_NAME = sb.Name
	sbd.SB_TYPE = sbt

	if objects.IsLogin(r) {
		sbd.IsLogin = true
		user, err := objects.GetUserFromRequest(r)
		if err == nil {
			sbd.SB_USER_PLACE = user.GetScoreboardPlace(sb) + 1
		}
	}

	for i, v := range sb.Top {
		switch sbt {
		case objects.ScoreboardType_Letter:
			sbd.SB_USERS = append(sbd.SB_USERS, "#"+strconv.Itoa(i+1)+": "+v.Username+" - "+strconv.Itoa(v.LetterFind))
		case objects.ScoreboardType_Words:
			sbd.SB_USERS = append(sbd.SB_USERS, "#"+strconv.Itoa(i+1)+": "+v.Username+" - "+strconv.Itoa(v.WordsFind))
		case objects.ScoreboardType_Points:
			sbd.SB_USERS = append(sbd.SB_USERS, "#"+strconv.Itoa(i+1)+": "+v.Username+" - "+strconv.Itoa(v.Points))
		case objects.ScoreboardType_loose:
			sbd.SB_USERS = append(sbd.SB_USERS, "#"+strconv.Itoa(i+1)+": "+v.Username+" - "+strconv.Itoa(v.Loose))
		case objects.ScoreboardType_win:
			sbd.SB_USERS = append(sbd.SB_USERS, "#"+strconv.Itoa(i+1)+": "+v.Username+" - "+strconv.Itoa(v.Wins))
		}
	}

	tp.Execute(w, sbd)
}
