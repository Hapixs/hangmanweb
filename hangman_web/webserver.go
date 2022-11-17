package hangmanweb

import (
	"encoding/csv"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"
)

var sessions = map[string](*WebGame){}

func StartServer() {
	InitWebHandlers()
	fs := http.FileServer(http.Dir("./web/"))
	http.Handle("/web/", http.StripPrefix("/web/", fs))

	wg.Add(1)
	go AutoSaveWorker(&wg)
	go LoadUserCSV()

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
}

var wg = sync.WaitGroup{}

func AutoSaveWorker(wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		time.Sleep(time.Second * 10)
		println("Saving user data ...")
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
		usermap[userId] = &User{records[i][0], userPoint, userId, records[i][3], false}
	}

	println("Loaded " + strconv.Itoa(len(usermap)) + " users from csv")
}

func SaveUserCSV() {
	records := [][]string{
		{"username", "points", "uniqueid", "password"},
	}
	for _, v := range usermap {
		if v.isAnnonyme {
			continue
		}
		records = append(records, []string{v.Username, strconv.Itoa(v.Points), strconv.Itoa(v.UniqueId), v.Password})
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
	println("User data saved !")
}
