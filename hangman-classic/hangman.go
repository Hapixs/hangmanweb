package main

import (
	"objects"
	"os"
)

func main() {
	game := objects.HangmanGame{}
	game.InitGame(os.Args[1:])
	game.StartGame()
}
