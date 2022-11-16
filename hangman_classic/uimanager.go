package hangman_classic

import (
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"time"
)

var informationHeadMessages = []string{}

func (game *HangmanGame) DisplayBody() {
	_, clearscreen, _ := game.Config.GetConfigItem(ConfigAutoClear)
	if clearscreen {
		ClearScreen()
	}
	println(informationHeadMessages[len(informationHeadMessages)-1])
	_, useAscii, _ := game.Config.GetConfigItem(ConfigUseAscii)
	if useAscii {
		for _, line := range game.cache.BuildASCIIWord(game.GetGameWord()) {
			println(line)
		}
	} else {
		println(game.GetGameWord())
	}

	for _, v := range game.cache.HangmanByStatus[game.GetGameTries()] {
		println(v)
	}

	gameMode, _, _ := game.Config.GetConfigItem(ConfigGameMode)
	maxTries, _, _ := game.Config.GetConfigItem(ConfigMaxTries)

	if gameMode != HARD {
		println("Used: " + game.GetGameUsed())
	}
	println("You have " + strconv.Itoa(maxTries-game.GetGameTries()) + " mistakes left.")
	print("Choose: ")
}

func AddInformationHeadMessage(message string) {
	informationHeadMessages = append(informationHeadMessages, message)
}

func (game *HangmanGame) DisplayLooseLogo() {
	speed := time.Second / 5
	gameover := ""
	for _, c := range "OH SNAP !" {
		ClearScreen()
		gameover += string(c)
		display := game.cache.BuildASCIIWord(gameover)
		for _, line := range display {
			println(line)
		}
		time.Sleep(speed)
	}

	println("The word was " + game.GetGameToFind())
}

func (game *HangmanGame) DisplayWinLogo() {
	speed := time.Second / 5
	gameover := ""
	for _, c := range "YOU WIN !" {
		ClearScreen()
		gameover += string(c)
		display := game.cache.BuildASCIIWord(gameover)
		for _, line := range display {
			println(line)
		}
		time.Sleep(speed)
	}

	println("The word was " + game.GetGameToFind())
}

func ClearScreen() {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}
