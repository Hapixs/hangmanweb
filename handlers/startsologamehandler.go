package handlers

import (
	"net/http"
	"objects"
)

func StartSoloPageHandler(w http.ResponseWriter, r *http.Request) {
	if !objects.IsLogin(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	gameMode := r.FormValue("difficulty")
	objects.StartNewGame(&w, r, gameMode)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
