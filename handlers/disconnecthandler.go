package handlers

import "net/http"

func disconnectHandler(w http.ResponseWriter, r *http.Request) {
	c := http.Cookie{Name: "user_id", Value: ""}
	http.SetCookie(w, &c)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
