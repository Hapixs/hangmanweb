package objects

import (
	"math/rand"
	"strconv"
)

// HERE IT IS ! THE MULTIPLAYER
// Not available for today

type GamePool struct {
	UniqueID int
	Games    []*WebGame
}

var PoolByUserID = map[string](*GamePool){}
var PoolByID = map[string](*GamePool){}

func CreatePool(webGame WebGame) *GamePool {
	gp := &GamePool{}
	gp.UniqueID = rand.Intn(10000000000)

	gp.Games = append(gp.Games, &webGame)
	Mutex.Lock()
	PoolByUserID[strconv.Itoa(webGame.User.UniqueId)] = gp
	PoolByID[strconv.Itoa(gp.UniqueID)] = gp
	Mutex.Unlock()

	return gp
}

func (gp *GamePool) AddUserGame(wg *WebGame) {
	gp.Games = append(gp.Games, wg)
	Mutex.Lock()
	PoolByUserID[strconv.Itoa(wg.User.UniqueId)] = gp
	Mutex.Unlock()
}

func GetPoolFromGame(wg WebGame) *GamePool {
	return PoolByUserID[strconv.Itoa(wg.User.UniqueId)]
}

func GetPoolFromID(id string) *GamePool {
	return PoolByID[id]
}
