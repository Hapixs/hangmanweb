package handlers

import (
	"net/http"
	"objects"
)

func RestartSoloGameHandler(w http.ResponseWriter, r *http.Request) {
	if !objects.IsLogin(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	Game := objects.GetGameFromCookies(w, r)
	if Game != nil && Game.Game != nil {
		objects.Mutex.Lock()
		objects.Sessions[Game.Game.PublicId] = nil
		objects.Mutex.Unlock()
	}
	gameMode := r.FormValue("difficulty")
	objects.StartNewGame(&w, r, gameMode)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
