package main

import (
	"errors"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

type User struct {
	Username   string
	Points     int
	UniqueId   int
	Password   string
	isAnnonyme bool
	Wins       int
	Loose      int
	Played     int
	LetterFind int
	WordsFind  int
}

var usermap = map[int](*User){}
var usermapHash = HashMap(usermap)

func (u *User) GenerateUniqueId() {
	rand.Seed(time.Now().Unix())
	u.UniqueId = rand.Intn(100000000)
}

func (u *User) SetUpUserCookies(w *http.ResponseWriter) {
	c := http.Cookie{Name: "user_id", Value: strconv.Itoa(u.UniqueId)}
	c.Expires.After(time.Now().Add(time.Hour))
	http.SetCookie(*w, &c)
	mutex.Lock()
	usermap[u.UniqueId] = u
	mutex.Unlock()
}

func (u *User) GetScoreboardPlace(sb Scoreboard) int {
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
	if usermap[id] == nil {
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
	user := usermap[ui]
	return user, nil
}