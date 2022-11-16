package hangman_classic

import (
	"os"
	"strings"
)

type Gamecache struct {
	HangmanByStatus map[int]([]string)
	Words           []string
	AsciiByChar     map[rune]([]string)
}

var FromSave = false

func (c *Gamecache) InitGameCache(game *HangmanGame) {
	_, _, fileName := game.Config.GetConfigItem(ConfigWordsList)
	if len(fileName) <= 0 {
		println("Please specify a word file")
		os.Exit(1)
	}
	content, err := os.ReadFile(fileName)
	if err != nil {
		println("Error! Unable to load word list '" + fileName + "'")
		os.Exit(1)
	}
	c.Words = strings.Split(string(content), "\n")

	_, _, hangmanFileName := game.Config.GetConfigItem(ConfigHangmanFile)
	content, err = os.ReadFile(hangmanFileName)
	if err != nil {
		println("Error! Unable to load hangman list '" + fileName + "'")
		os.Exit(1)
	}
	hangmanStatContentSplited := strings.Split(string(content), "\n")

	maxTries, _, _ := game.Config.GetConfigItem(ConfigMaxTries)
	c.HangmanByStatus = make(map[int][]string, maxTries)
	for i := 0; i < maxTries; i++ {
		hangmanHeight, _, _ := game.Config.GetConfigItem(ConfigHangmanHeight)
		currentMin := i * hangmanHeight
		currentMax := currentMin + hangmanHeight
		c.HangmanByStatus[i+1] = hangmanStatContentSplited[currentMin:currentMax]
	}

	_, _, asciiFileName := game.Config.GetConfigItem(ConfigASCIIFile)

	content, err = os.ReadFile(asciiFileName)
	if err != nil {
		println("Error! Unable to load ascii file '" + fileName + "'")
		os.Exit(1)
	}

	asciiCharacterContentSplited := strings.Split(string(content), "\n")
	c.AsciiByChar = make(map[rune][]string)
	for i := 0; i < 127-32; i++ {
		asciiHeight, _, _ := game.Config.GetConfigItem(ConfigASCIIHeight)
		currentMin := i * asciiHeight
		currentMax := currentMin + asciiHeight
		c.AsciiByChar[rune(i+32)] = asciiCharacterContentSplited[currentMin:currentMax]
	}
	if !FromSave {
		wordToFind := strings.ReplaceAll(GetRandomWord(c.Words), "\r", "")
		wordToFind = strings.ReplaceAll(wordToFind, "\n", "")
		game.SetGameToFindEncoded(EncodeStrInBase64(string(GetEncodedStringInSha256(wordToFind))) + " # Lol tu le trouvra pas !")
		game.SetGameToFind(wordToFind)
		game.SetGameWord(game.SetupGameWord(wordToFind))
	} else {
		for _, word := range c.Words {
			word = strings.ReplaceAll(word, "\r", "")
			word = strings.ReplaceAll(word, "\n", "")
			if EncodeStrInBase64(string(GetEncodedStringInSha256(word))) == strings.Split(game.GetGameToFindEncoded(), " # Lol tu le trouvra pas !")[0] {
				game.SetGameToFind(word)
				return
			}
		}
		println("Unable to load, save corrupted..")
		os.Exit(1)
	}
}
