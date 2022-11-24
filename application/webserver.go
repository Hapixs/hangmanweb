package main

import (
	"crypto/md5"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

func main() {
	StartServer()
}

var sessions = map[string](*WebGame){}

func StartServer() {
	InitWebHandlers()
	fs := http.FileServer(http.Dir("./web/"))
	http.Handle("/web/", http.StripPrefix("/web/", fs))

	go AutoSaveWorker()
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

func InitWebHandlers() {
	http.HandleFunc("/hangman", HangmanPostHandler)
	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/reset", ResetHandler)
	http.HandleFunc("/login", LoginPostHandler)
	http.HandleFunc("/startsologame", StartSoloPageHandler)
	http.HandleFunc("/nolog", AnnoLoginHandler)
	http.HandleFunc("/logout", DisconnectHandler)
	http.HandleFunc("/restartsologame", RestartSoloGameHandler)
	http.HandleFunc("/statistics", StatisticsHandler)
	http.HandleFunc("/scoreboard", ScoreboardHandler)
}

func AutoSaveWorker() {
	for {
		time.Sleep(time.Second * 10)
		SaveUserCSV()
	}
}

func LoadUserCSV() {

	f, err := os.Open("users.csv")

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
		usermap[userId] = &User{Username: records[i][0], Points: userPoint, UniqueId: userId, Password: records[i][3], isAnnonyme: false, Wins: userWins, Loose: userLoose, Played: Played, LetterFind: LetterFind, WordsFind: WordsFind}
	}

	println("Loaded " + strconv.Itoa(len(usermap)) + " users from csv")
}

func SaveUserCSV() {
	hash := HashMap(usermap)
	if hash == usermapHash {
		return
	}
	records := [][]string{
		{"username", "points", "uniqueid", "password", "wins", "loose", "played", "lettersfind", "wordsfind"},
	}
	for _, v := range usermap {
		if v.isAnnonyme {
			continue
		}
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

	f, err := os.Create("users.csv")

	if err != nil {
		log.Fatalln("failed to open file", err)
	}

	w := csv.NewWriter(f)

	if err := w.WriteAll(records); err != nil {
		log.Fatalln("error writing record to file", err)
	}
	f.Close()
	w.Flush()
	println("Saved " + strconv.Itoa(len(records)-1) + " users")
	usermapHash = hash
}

func osSignalHandler(signal os.Signal) {
	if signal == syscall.SIGTERM || signal == syscall.SIGINT {
		fmt.Println("Program will terminate now.")
		SaveUserCSV()
		os.Exit(0)
	}
}

func HashMap(m map[int](*User)) [16]byte {
	arrBytes := []byte{}
	for _, item := range m {
		jsonBytes, _ := json.Marshal(item)
		arrBytes = append(arrBytes, jsonBytes...)
	}
	return md5.Sum(arrBytes)
}
