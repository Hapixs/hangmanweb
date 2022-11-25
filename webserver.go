package main

import (
	"encoding/csv"
	"fmt"
	"handlers"
	"log"
	"net/http"
	"objects"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
	"utils"
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
		SaveUserCSV()
		LoadUserCSV()
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
func autoSaveWorker() {
	for {
		time.Sleep(time.Second * 10)
		SaveUserCSV()
	}
}

var lastUsermapHash [16]byte

func SaveUserCSV() {
	hash := utils.UserMapHash(objects.Usermap)
	if hash == lastUsermapHash {
		return
	}

	records := [][]string{
		{"username", "points", "uniqueid", "password", "wins", "loose", "played", "lettersfind", "wordsfind"},
	}

	for _, v := range objects.Usermap {
		if !v.IsAnnonyme {
			records = append(records,
				[]string{v.Username,
					strconv.Itoa(v.Points),
					strconv.Itoa(v.UniqueId),
					v.Password,
					strconv.Itoa(v.Wins),
					strconv.Itoa(v.Loose),
					strconv.Itoa(v.Played),
					strconv.Itoa(v.LetterFind),
					strconv.Itoa(v.WordsFind)})
		}
	}

	f, err := os.Create("data/users.csv")

	if err != nil {
		log.Fatalln("failed to open file", err)
		return
	}

	w := csv.NewWriter(f)

	if err := w.WriteAll(records); err != nil {
		log.Fatalln("error writing record to file", err)
		return
	}

	f.Close()
	w.Flush()
	println("Saved " + strconv.Itoa(len(records)-1) + " users")
	lastUsermapHash = hash
}
