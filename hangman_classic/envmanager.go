package hangman_classic

import (
	"math/rand"
	"time"
)

func InitEnvironement() {
	rand.Seed(time.Now().UnixMicro())
}
