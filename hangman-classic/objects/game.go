package hangmanclassicobjects

import (
	"encoding/json"
	"math/rand"
	"os"
	"strconv"
	"time"
)

type HangmanGame struct {
	// PUBLIC
	Tries         int        // Nb of user's mistakes
	Word          string     // The displayed word
	EncodedToFind string     // The encoded word (Only used for the save)
	Used          string     // This string contains all used letter in this game
	Config        GameConfig // The actual game config
	Gamestatus    int        // The actual game status
	PublicId      string

	// PRIVATE
	tofind     string
	cache      Gamecache
	executions []GameExecution
}

type GameMode int
type Gamestatus int

const (
	NORMAL = 0
	HARD   = 1

	PLAYING Gamestatus = 0
	ENDED   Gamestatus = 1
)

func (game *HangmanGame) InitGame(args []string) {
	InitEnvironement()
	game.Config.InitConfig()

	game.GameProcessArguments(args)

	gc := Gamecache{}
	gc.InitGameCache(game)
	game.cache = gc

	game.InitGameExecutions()

	maxTries, _, _ := game.Config.GetConfigItem(ConfigMaxTries)
	AddInformationHeadMessage("Good Luck, you have " + strconv.Itoa(maxTries-game.Tries) + "  attempts.")
	game.Config.SetConfigItemValue(ConfigIsInit, true)
	game.GeneratePublicId()
	game.Gamestatus = int(PLAYING)
}

func (game *HangmanGame) GeneratePublicId() {
	game.PublicId = strconv.Itoa(rand.Intn(150000000))
}

func (game *HangmanGame) StartGame() {
	_, isInit, _ := game.Config.GetConfigItem(ConfigIsInit)
	if !isInit {
		game.InitGame([]string{})
	}
	_, b, _ := game.Config.GetConfigItem(ConfigMultipleWorkers)
	if b {
		go game.processExecutionsFunc()
	} else {
		game.processExecutionsFunc()
	}
}

func (game *HangmanGame) Kill() {
	if game == nil {
		return
	}
	game.Gamestatus = int(ENDED)
}

func (game *HangmanGame) processExecutionsFunc() {
	for game.Gamestatus == int(PLAYING) {
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
}

func (game *HangmanGame) AddGameTry() int {
	game.Tries++
	maxtries, _, _ := game.Config.GetConfigItem(ConfigMaxTries)
	if game.Tries > maxtries {
		game.Tries = maxtries
	}
	return game.Tries
}

func (game *HangmanGame) GetGameTries() int {
	return game.Tries
}

func (game *HangmanGame) SetGameWord(word string) string {
	game.Word = word
	return game.Word
}

func (game *HangmanGame) GetGameWord() string {
	return game.Word
}

func (game *HangmanGame) GetGameToFindEncoded() string {
	return game.EncodedToFind
}

func (game *HangmanGame) SetGameToFindEncoded(tofind string) string {
	game.EncodedToFind = tofind
	return game.tofind
}

func (game *HangmanGame) SetGameToFind(tofind string) string {
	game.tofind = tofind
	return game.tofind
}

func (game *HangmanGame) GetGameToFind() string {
	_, b, _ := game.Config.GetConfigItem(ConfigUseAscii)
	if b {
		return ConvertToUnicode(game.tofind)
	}
	return game.tofind
}

func (game *HangmanGame) GetGameUsed() string {
	return game.Used
}

func (game *HangmanGame) AddGameUsed(r rune) string {
	game.Used += string(r)
	return game.Used
}

func (game *HangmanGame) SaveGame() {

	_, _, fileName := game.Config.GetConfigItem("saves/" + ConfigSaveFile)
	saveEnc, err := json.Marshal(game)
	if err != nil {
		println("Error: JSON error")
		return
	}

	file, err := os.Create(fileName)
	if err != nil {
		println("Error: Unable to save 'save.txt'")
		return
	}

	file.Write([]byte(EncodeStrInBase64(string(saveEnc))))
	file.Close()
}

func (g *HangmanGame) GetConfig() *GameConfig {
	return &g.Config
}

func (game *HangmanGame) UpdateGameWord(toFind string, word string, letterToCheck rune) string {
	wordR := []rune(word)
	indexs := GetOccurenceLetter(toFind, letterToCheck)
	for _, index := range indexs {
		wordR[index] = letterToCheck
	}
	word = string(wordR)
	game.Word = word
	return word
}

func (game *HangmanGame) SetupGameWord(startupword string) string {
	size := len([]rune(startupword))
	runeTableWord := make([]rune, size)
	for i := 0; i < len(runeTableWord); i++ {
		runeTableWord[i] = '_'
	}
	for nbLettersToShow := len([]rune(runeTableWord))/2 - 1; nbLettersToShow > 0; nbLettersToShow-- {
		randomTableI := rand.Intn(len([]rune(runeTableWord)))
		if runeTableWord[randomTableI] != '_' {
			nbLettersToShow++
		} else {
			runeTableWord[randomTableI] = []rune(startupword)[randomTableI]
		}
	}
	listOfLettersGiven := make([]rune, len([]rune(runeTableWord)))
	for _, letter := range runeTableWord {
		if letter != '_' {
			listOfLettersGiven = append(listOfLettersGiven, letter)
		}
	}
	for _, letter := range listOfLettersGiven {
		runeTableWord = []rune(game.UpdateGameWord(startupword, string(runeTableWord), letter))
		game.AddGameUsed(letter)
	}
	game.Word = string(runeTableWord)
	return string(runeTableWord)
}

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

	if gameMode != int(HARD) {
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

func (game *HangmanGame) GameProcessArguments(args []string) {
	for len(args) > 0 {
		arg := args[0]
		if arg[0] == "-"[0] {
			if arg == "--help" || arg == "-h" {
				flagShowHelpMessage(game, args)
				return
			}
			if val, ok := flagExecutors[arg]; ok {
				cmdFlag := val
				if cmdFlag.IsAliase {
					cmdFlag = flagExecutors[cmdFlag.AliaseOf]
				}
				args = cmdFlag.FlagExecutor(game, args)
			} else {
				args = append(args[:0], args[1:]...)
				println("Can't find argument " + arg)
			}
		} else {
			game.Config.SetConfigItemValue(ConfigWordsList, arg)
			args = append(args[:0], args[1:]...)
		}
	}
}
