package handlers

import (
	"net/http"
	"objects"
)

func annoLoginHandler(w http.ResponseWriter, r *http.Request) {
	ano := objects.User{IsAnnonyme: true, Username: "Annonyme"}
	ano.GenerateUniqueId()
	ano.SetUpUserCookies(&w)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
