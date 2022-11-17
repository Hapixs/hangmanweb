package hangmanweb

import (
	"bytes"
	"hangman_classic"
	"net/http"
	"os"
	"sync"
)

var sessions = map[string](*WebGame){}

type WebGame struct {
	Game    *hangman_classic.HangmanGame
	Input   bytes.Buffer
	IsWin   bool
	IsLoose bool
	User    *User
	PoolId  string
}

func StartServer() {
	InitWebHandlers()
	fs := http.FileServer(http.Dir("./web/"))
	http.Handle("/web/", http.StripPrefix("/web/", fs))
	http.ListenAndServe(":8080", nil)
}

func InitWebHandlers() {
	http.HandleFunc("/hangman", PostHandler)
	http.HandleFunc("/", GetHandler)
	http.HandleFunc("/reset", ResetHandler)
	http.HandleFunc("/login", PostLogin)
}

var mutex = &sync.Mutex{}

func getGameFromCookies(w http.ResponseWriter, r *http.Request) *WebGame {
	if !IsLogin(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
	c, err := r.Cookie("sessionid")
	sessionid := ""
	if err != nil || sessions[c.Value] == nil || sessions[c.Value].User == nil {
		Game := &hangman_classic.HangmanGame{}
		prepareGameForWeb(Game)
		user, err := GetUserFromRequest(r)
		if err != nil {
			user = &User{Username: "Unknown"}
			user.GenerateUniqueId()
			user.SetUpUserCookies(&w)
		}
		mutex.Lock()
		sessions[Game.PublicId] = &WebGame{Game, bytes.Buffer{}, false, false, user, ""}
		mutex.Unlock()
		http.SetCookie(w, &http.Cookie{Name: "sessionid", Value: Game.PublicId})
		defer Game.StartGame()
		sessionid = Game.PublicId
	} else {
		sessionid = c.Value
	}

	return sessions[sessionid]
}

func getWebGameFromId(id string) *WebGame {
	mutex.Lock()
	s := sessions[id]
	mutex.Unlock()
	return s
}

func prepareGameForWeb(Game *hangman_classic.HangmanGame) {
	os.Args = append(os.Args, "words.txt")
	Game.InitGame()
	Game.ReplaceExecution(overridedExecutionWaitForInput, string(hangman_classic.DefaultExecutionWaitForInput))
	Game.ReplaceExecution(overridedExecutionCheckForRemainingTries, string(hangman_classic.DefaultExecutionCheckForRemainingTries))
	Game.ReplaceExecution(overridedExecutionCheckForWordDiscover, string(hangman_classic.DefaultExecutionCheckForWordDiscover))
	Game.ReplaceExecution(overridedExecutionCheckForWord, string(hangman_classic.DefaultExecutionCheckForWord))
	Game.Config.SetConfigItemValue(hangman_classic.ConfigMultipleWorkers, true)
	Game.RemoveExecution(hangman_classic.DefaultExecutionDisplayBody)
	println("Prepared " + Game.PublicId)
}

// Hmm Un multijoueur ?
// jvais dormir et demain jte le code
// tqt on est la
// :*
type Pool struct {
	PublicId string
	Users    []User
}

var PoolMap = map[string](Pool){}
