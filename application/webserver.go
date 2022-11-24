package main

import (
	"encoding/csv"
	"fmt"
	"handlers"
	"net/http"
	"objects"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

func main() {
	StartServer()
}

func StartServer() {
	InitWebServer()

	go autoSaveWorker()

	go LoadUserCSV()

	sigchnl := make(chan os.Signal, 1)
	signal.Notify(sigchnl)

	go func() {
		for {
			s := <-sigchnl
			osSignalHandler(s)
		}
	}()

	http.ListenAndServe(":8080", nil)
}

func InitWebServer() {
	fs := http.FileServer(http.Dir("./static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/hangman", handlers.HangmanPostHandler)
	http.HandleFunc("/", handlers.IndexHandler)
	http.HandleFunc("/reset", handlers.ResetHandler)
	http.HandleFunc("/login", handlers.LoginPostHandler)
	http.HandleFunc("/startsologame", handlers.StartSoloPageHandler)
	http.HandleFunc("/nolog", handlers.AnnoLoginHandler)
	http.HandleFunc("/logout", handlers.DisconnectHandler)
	http.HandleFunc("/restartsologame", handlers.RestartSoloGameHandler)
	http.HandleFunc("/statistics", handlers.StatisticsHandler)
	http.HandleFunc("/scoreboard", handlers.ScoreboardHandler)
}

func LoadUserCSV() {

	f, err := os.Open("data/users.csv")

	if err != nil {
		println("No users.csv found.")
		return
	}

	r := csv.NewReader(f)
	records, err := r.ReadAll()

	if err != nil {
		println("CSV READ FAILED !")
		return
	}

	for i := 1; i < len(records); i++ {
		userId, _ := strconv.Atoi(records[i][2])
		userPoint, _ := strconv.Atoi(records[i][1])
		userWins, _ := strconv.Atoi(records[i][4])
		userLoose, _ := strconv.Atoi(records[i][5])
		Played, _ := strconv.Atoi(records[i][6])
		LetterFind, _ := strconv.Atoi(records[i][7])
		WordsFind, _ := strconv.Atoi(records[i][8])
		objects.Usermap[userId] = &objects.User{Username: records[i][0], Points: userPoint, UniqueId: userId, Password: records[i][3], IsAnnonyme: false, Wins: userWins, Loose: userLoose, Played: Played, LetterFind: LetterFind, WordsFind: WordsFind}
	}

	println("Loaded " + strconv.Itoa(len(objects.Usermap)) + " users from csv")
}

func osSignalHandler(signal os.Signal) {
	if signal == syscall.SIGTERM || signal == syscall.SIGINT {
		fmt.Println("Program will terminate now.")
		SaveUserCSV()
		os.Exit(0)
	}
}
