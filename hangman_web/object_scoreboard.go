package hangmanweb

import (
	"sort"
	"time"
)

const (
	ScoreboardType_win    = "sb_win"
	ScoreboardType_loose  = "sb_loose"
	ScoreboardType_Letter = "sb_letter"
	ScoreboardType_Words  = "sb_words"
	ScoreboardType_Points = "sb_points"
)

type Scoreboard struct {
	Name string
	Type string
	Top  []*User
}

func DynamicScoreboardLoader() {

	time.Sleep(time.Minute)
}

func BuildScoreboard(Type string) *Scoreboard {
	sb := &Scoreboard{}

	sb.Type = Type

	switch Type {
	case ScoreboardType_Words:
		sb.Name = "Top Words finder"
	case ScoreboardType_Letter:
		sb.Name = "Top letter finder"
	case ScoreboardType_Points:
		sb.Name = "Top points"
	case ScoreboardType_loose:
		sb.Name = "Top Loose"
	case ScoreboardType_win:
		sb.Name = "Top win"
	}

	for _, v := range usermap {
		sb.Top = append(sb.Top, v)
	}

	sort.Slice(sb.Top, func(i, j int) bool {
		switch Type {
		case ScoreboardType_Words:
			return sb.Top[i].WordsFind > sb.Top[j].WordsFind
		case ScoreboardType_Letter:
			return sb.Top[i].LetterFind > sb.Top[j].LetterFind
		case ScoreboardType_Points:
			return sb.Top[i].Points > sb.Top[j].Points
		case ScoreboardType_loose:
			return sb.Top[i].Loose > sb.Top[j].Loose
		case ScoreboardType_win:

			return sb.Top[i].Loose > sb.Top[j].Wins
		}
		return false
	})

	return sb
}
