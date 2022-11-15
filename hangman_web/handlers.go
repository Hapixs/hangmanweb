package hangmanweb

import (
	"net/http"
	"text/template"
)

func PostHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		if err := r.ParseForm(); err != nil {
			println("ParseForm() err: %v", err)
			return
		}
	}
	WebInputbuffer.Write([]byte(r.Form.Get("input")))
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func GetHandler(w http.ResponseWriter, r *http.Request) {
	tp := template.Must(template.ParseFiles("../hangman_web/web/index.html"))

	tp.Execute(w, nil)
}
