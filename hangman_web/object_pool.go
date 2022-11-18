package hangmanweb

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
	mutex.Lock()
	PoolByUserID[strconv.Itoa(webGame.User.UniqueId)] = gp
	PoolByID[strconv.Itoa(gp.UniqueID)] = gp
	mutex.Unlock()

	return gp
}

func (gp *GamePool) AddUserGame(wg *WebGame) {
	gp.Games = append(gp.Games, wg)
	mutex.Lock()
	PoolByUserID[strconv.Itoa(wg.User.UniqueId)] = gp
	mutex.Unlock()
}

func GetPoolFromGame(wg WebGame) *GamePool {
	return PoolByUserID[strconv.Itoa(wg.User.UniqueId)]
}

func GetPoolFromID(id string) *GamePool {
	return PoolByID[id]
}
