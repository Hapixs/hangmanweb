package hangman_classic

import (
	"os"
	"strconv"
)

var isInit = false

func (game *HangmanGame) InitGame() {
	args := os.Args[1:]

	game.Config.InitConfig()

	game.GameProcessArguments(args)

	gc := Gamecache{}

	InitEnvironement()
	gc.InitGameCache(game)
	game.cache = gc

	game.InitGameExecutions()

	maxTries, _, _ := game.Config.GetConfigItem(ConfigMaxTries)
	AddInformationHeadMessage("Good Luck, you have " + strconv.Itoa(maxTries-game.Tries) + "  attempts.")
	isInit = true
	game.Gamestatus = PLAYING
}

func (game *HangmanGame) StartGame() {
	if !isInit {
		game.InitGame()
	}
	go game.processExecutionsFunc()
}

func (game *HangmanGame) Kill() {
	game.Gamestatus = ENDED
}

func (game *HangmanGame) processExecutionsFunc() {
	for game.Gamestatus == PLAYING {
		userInput := ""
		for _, execution := range game.executions {
			if execution.Func(&userInput, game) {
				break
			}
		}
	}
}

func (game *HangmanGame) WinGame() {
	game.DisplayWinLogo()
	QuitGame()
}

func QuitGame() {
	os.Exit(0)
}
