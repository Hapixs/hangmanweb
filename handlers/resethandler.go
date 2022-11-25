package handlers

import (
	"net/http"
	"objects"
)

func resetHandler(w http.ResponseWriter, r *http.Request) {
	Game := objects.GetGameFromCookies(w, r)
	Game.Game.Kill()
	objects.Mutex.Lock()
	objects.Sessions[Game.Game.PublicId] = nil
	objects.Mutex.Unlock()
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
