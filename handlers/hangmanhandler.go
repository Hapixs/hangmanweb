package handlers

import (
	"net/http"
	"objects"
)

func HangmanPostHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		if err := r.ParseForm(); err != nil {
			println("ParseForm() err: %v", err)
			return
		}
		Game := objects.GetGameFromCookies(w, r)
		Game.Input.Write([]byte(r.Form.Get("input")))
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
