package objects

import (
	"bytes"
	"net/http"
	"sync"

	"github.com/Hapixs/hangmanclassic"
)

type WebGame struct {
	Game     *hangmanclassic.HangmanGame
	Input    bytes.Buffer
	IsWin    bool
	IsLoose  bool
	User     *User
	Gamemode string
}

var mutex = &sync.Mutex{}

func getGameFromCookies(w http.ResponseWriter, r *http.Request) *WebGame {
	if !IsLogin(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}

	user, _ := GetUserFromRequest(r)
	c, err := r.Cookie("sessionid")
	sessionid := ""
	if err != nil || sessions[c.Value] == nil || sessions[c.Value].User == nil {
		return &WebGame{nil, bytes.Buffer{}, false, false, user, "easy"}
	} else {
		sessionid = c.Value
	}

	return sessions[sessionid]
}

func StartNewGame(w *http.ResponseWriter, r *http.Request, gameMode string) {
	Game := &hangmanclassic.HangmanGame{}
	args := []string{}
	switch gameMode {
	default:
		args = append(args, "words.txt")
	case "medium":
		args = append(args, "words2.txt")
	case "hard":
		args = append(args, "words3.txt")
		args = append(args, "--hard")
	}
	prepareGameForWeb(Game, args)
	user, err := GetUserFromRequest(r)
	if err != nil {
		user = &User{Username: "Unknown"}
		user.GenerateUniqueId()
		user.SetUpUserCookies(w)
	}
	mutex.Lock()
	sessions[Game.PublicId] = &WebGame{Game, bytes.Buffer{}, false, false, user, gameMode}
	mutex.Unlock()
	http.SetCookie(*w, &http.Cookie{Name: "sessionid", Value: Game.PublicId})
	defer Game.StartGame()
}

func getWebGameFromId(id string) *WebGame {
	mutex.Lock()
	s := sessions[id]
	mutex.Unlock()

	return s
}

func prepareGameForWeb(Game *hangmanclassic.HangmanGame, args []string) {
	Game.InitGame(args)
	Game.ReplaceExecution(overridedExecutionWaitForInput, string(hangmanclassic.DefaultExecutionWaitForInput))
	Game.ReplaceExecution(overridedExecutionCheckForRemainingTries, string(hangmanclassic.DefaultExecutionCheckForRemainingTries))
	Game.ReplaceExecution(overridedExecutionCheckForWordDiscover, string(hangmanclassic.DefaultExecutionCheckForWordDiscover))
	Game.ReplaceExecution(overridedExecutionCheckForWord, string(hangmanclassic.DefaultExecutionCheckForWord))
	Game.ReplaceExecution(overridedExecutionCheckForLetterOccurence, string(hangmanclassic.DefaultExecutionCheckForLetterOccurence))
	Game.Config.SetConfigItemValue(hangmanclassic.ConfigMultipleWorkers, true)
	Game.RemoveExecution(hangmanclassic.DefaultExecutionDisplayBody)
}
