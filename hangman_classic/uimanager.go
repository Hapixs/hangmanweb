package hangman_classic

import (
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"time"

	"github.com/rivo/tview"
)

var informationHeadMessages = []string{}
var debugMessages = []string{}
var inforamtionFooterMessages = []string{}

var useBetterTerminal = false

var termApplication = tview.NewApplication()
var termRoot = tview.NewFlex()

func InitUI() {
	_, useBetterTerminal, _ = GetConfigItem(ConfigBetterTerminal)
	if useBetterTerminal {
		err := termApplication.SetRoot(termRoot, true).Run()
		if err != nil {
			useBetterTerminal = false
			println("Error: unable to use better terminal feature")
		}
	}
}

func DisplayBody() {
	if useBetterTerminal {
		termRoot.Clear()
		hangmanLogo := ""
		for _, line := range BuildASCIIWord("HANGMAN") {
			hangmanLogo += line + "\n"
		}
		termRoot.SetDirection(tview.FlexColumn)
		termRoot.AddItem(tview.NewTextView().SetText(hangmanLogo), 0, 1, false)
		termRoot.AddItem(tview.NewFlex().
			AddItem(tview.NewTextView().SetText("GAME"), 0, 1, false).
			AddItem(tview.NewTextView().SetText("HELP"), 0, 1, false), 0, 1, false)
		termApplication.SetRoot(termRoot, true)
	} else {
		_, clearscreen, _ := GetConfigItem(ConfigAutoClear)
		if clearscreen {
			ClearScreen()
		}
		println(informationHeadMessages[len(informationHeadMessages)-1])
		_, useAscii, _ := GetConfigItem(ConfigUseAscii)
		if useAscii {
			for _, line := range BuildASCIIWord(GetGameWord()) {
				println(line)
			}
		} else {
			println(GetGameWord())
		}

		for _, v := range GetCacheHangmanByIndex(GetGameTries()) {
			println(v)
		}

		if gameMode != HARD {
			println("Used: " + GetGameUsed())
		}
		println("You have " + strconv.Itoa(maxTries-GetGameTries()) + " mistakes left.")
		print("Choose: ")
	}
}

func addInformationHeadMessage(message string) {
	informationHeadMessages = append(informationHeadMessages, message)
}

func DisplayLooseLogo() {
	speed := time.Second / 5
	gameover := ""
	for _, c := range "OH SNAP !" {
		ClearScreen()
		gameover += string(c)
		display := BuildASCIIWord(gameover)
		for _, line := range display {
			println(line)
		}
		time.Sleep(speed)
	}

	println("The word was " + GetGameToFind())
}

func DisplayWinLogo() {
	speed := time.Second / 5
	gameover := ""
	for _, c := range "YOU WIN !" {
		ClearScreen()
		gameover += string(c)
		display := BuildASCIIWord(gameover)
		for _, line := range display {
			println(line)
		}
		time.Sleep(speed)
	}

	println("The word was " + GetGameToFind())
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
