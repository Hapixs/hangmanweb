package hangmanclassicobjects

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
	wordcontent, worderr := os.ReadFile(fileName)
	if worderr != nil {
		println("Error! Unable to load word list '" + fileName + "'")
		os.Exit(1)
	}
	c.Words = strings.Split(string(wordcontent), "\n")

	_, _, hangmanFileName := game.Config.GetConfigItem(ConfigHangmanFile)
	hangmancontent, hangmanerr := os.ReadFile(hangmanFileName)
	if hangmanerr == nil {
		hangmanStatContentSplited := strings.Split(string(hangmancontent), "\n")

		maxTries, _, _ := game.Config.GetConfigItem(ConfigMaxTries)
		c.HangmanByStatus = make(map[int][]string, maxTries)
		for i := 0; i < maxTries; i++ {
			hangmanHeight, _, _ := game.Config.GetConfigItem(ConfigHangmanHeight)
			currentMin := i * hangmanHeight
			currentMax := currentMin + hangmanHeight
			c.HangmanByStatus[i+1] = hangmanStatContentSplited[currentMin:currentMax]
		}
	} else {
		println("Error! Unable to load hangman list '" + fileName + "'")
	}

	_, _, asciiFileName := game.Config.GetConfigItem(ConfigASCIIFile)
	asciicontent, asciierr := os.ReadFile(asciiFileName)
	if asciierr == nil {
		asciiCharacterContentSplited := strings.Split(string(asciicontent), "\n")
		c.AsciiByChar = make(map[rune][]string)
		for i := 0; i < 127-32; i++ {
			asciiHeight, _, _ := game.Config.GetConfigItem(ConfigASCIIHeight)
			currentMin := i * asciiHeight
			currentMax := currentMin + asciiHeight
			c.AsciiByChar[rune(i+32)] = asciiCharacterContentSplited[currentMin:currentMax]
		}
	} else {
		println("Error! Unable to load ascii file '" + asciiFileName + "'")
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
