package objects

import (
	"crypto/md5"
	"encoding/csv"
	"encoding/json"
	"errors"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"
)

type User struct {
	Username   string
	Points     int
	UniqueId   int
	Password   string
	IsAnnonyme bool
	Wins       int
	Loose      int
	Played     int
	LetterFind int
	WordsFind  int
}

var Usermap = map[int](*User){}

func (u *User) GenerateUniqueId() {
	rand.Seed(time.Now().Unix())
	u.UniqueId = rand.Intn(100000000)
}

func (u *User) SetUpUserCookies(w *http.ResponseWriter) {
	c := http.Cookie{Name: "user_id", Value: strconv.Itoa(u.UniqueId)}
	c.Expires.After(time.Now().Add(time.Hour))
	http.SetCookie(*w, &c)
	Mutex.Lock()
	Usermap[u.UniqueId] = u
	Mutex.Unlock()
}

func (u *User) GetScoreboardPlace(sb *Scoreboard) int {
	for i, v := range sb.Top {
		if v.UniqueId == u.UniqueId {
			return i
		}
	}
	return -1
}

func (u *User) GetWinRatio() float32 {
	return float32(u.Played) / float32(u.Wins)
}

func (u *User) GetLooseRatio() float32 {
	return float32(u.Played) / float32(u.Loose)
}

func IsLogin(r *http.Request) bool {
	c, err := r.Cookie("user_id")
	if err != nil || c.Value == "" {
		return false
	}
	id, _ := strconv.Atoi(c.Value)
	if Usermap[id] == nil {
		return false
	}
	return err == nil
}

func GetUserFromRequest(r *http.Request) (*User, error) {
	if !IsLogin(r) {
		return &User{}, errors.New("User not found")
	}
	uniqueid, _ := r.Cookie("user_id")
	ui, _ := strconv.Atoi(uniqueid.Value)
	user := Usermap[ui]
	return user, nil
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
		Usermap[userId] = &User{Username: records[i][0], Points: userPoint, UniqueId: userId, Password: records[i][3], IsAnnonyme: false, Wins: userWins, Loose: userLoose, Played: Played, LetterFind: LetterFind, WordsFind: WordsFind}
	}

	println("Loaded " + strconv.Itoa(len(Usermap)) + " users from csv")
}

var lastUsermapHash [16]byte

func SaveUserCSV() {
	hash := userMapHash(Usermap)
	if hash == lastUsermapHash {
		return
	}
	records := [][]string{
		{"username", "points", "uniqueid", "password", "wins", "loose", "played", "lettersfind", "wordsfind"},
	}

	for _, v := range Usermap {
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

	os.Mkdir("data/", os.ModePerm)

	f, err := os.Create("data/users.csv")

	if err != nil {
		println("Error during file creation for user.csv")
		println(err.Error())
		os.Exit(1)
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

func userMapHash(m map[int](*User)) [16]byte {
	arrBytes := []byte{}
	for _, item := range m {
		jsonBytes, _ := json.Marshal(item)
		arrBytes = append(arrBytes, jsonBytes...)
	}
	return md5.Sum(arrBytes)
}
