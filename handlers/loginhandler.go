package handlers

import (
	"hangmanclassicobjects"
	"net/http"
	"objects"
)

func loginPostHandler(w http.ResponseWriter, r *http.Request) {
	if !objects.IsLogin(r) {
		switch r.Method {
		case "POST":
			if err := r.ParseForm(); err != nil {
				println("ParseForm() err: %v", err)
				return
			}
		}

		encodedPass := string(hangmanclassicobjects.GetEncodedStringInSha256(r.Form.Get("password")))

		for _, v := range objects.Usermap {
			if v.Username == r.Form.Get("username") {
				if v.Password == encodedPass {
					v.SetUpUserCookies(&w)
					http.Redirect(w, r, "/", http.StatusSeeOther)
					return
				}
				http.Redirect(w, r, "/", http.StatusSeeOther)
				return
			}
		}

		user := &objects.User{Username: r.Form.Get("username"), Password: encodedPass, IsAnnonyme: false, Points: 0, Wins: 0, Loose: 0, Played: 0, LetterFind: 0, WordsFind: 0}
		user.GenerateUniqueId()
		user.SetUpUserCookies(&w)
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
