package main

import (
	"encoding/csv"
	"log"
	"objects"
	"os"
	"strconv"
	"time"
	"utils"
)

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
