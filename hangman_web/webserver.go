package hangmanweb

import (
	"net/http"
)

func StartServer() {
	InitWebHandlers()
	http.ListenAndServe(":8080", nil)
}

func InitWebHandlers() {
	http.HandleFunc("/hangman", PostHandler)
	http.HandleFunc("/", GetHandler)
	http.HandleFunc("/reset", ResetHandler)
}
