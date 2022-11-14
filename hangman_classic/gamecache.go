package hangman_classic

import (
	"os"
	"strings"
)

var hangmanByStatus = map[int]([]string){}
var words []string
var asciiByChar = map[rune]([]string){}

var FromSave = false

func InitGameCache() {

	_, _, fileName := GetConfigItem(ConfigWordsList)
	if len(fileName) <= 0 {
		println("Please specify a word file")
		os.Exit(1)
	}
	content, err := os.ReadFile(fileName)
	if err != nil {
		println("Error! Unable to load word list '" + fileName + "'")
		os.Exit(1)
	}
	words = strings.Split(string(content), "\n")

	_, _, hangmanFileName := GetConfigItem(ConfigHangmanFile)
	content, err = os.ReadFile(hangmanFileName)
	if err != nil {
		println("Error! Unable to load hangman list '" + fileName + "'")
		os.Exit(1)
	}
	hangmanStatContentSplited := strings.Split(string(content), "\n")

	maxTires, _, _ := GetConfigItem(ConfigMaxTries)
	for i := 0; i < maxTires; i++ {
		hangmanHeight, _, _ := GetConfigItem(ConfigHangmanHeight)
		currentMin := i * hangmanHeight
		currentMax := currentMin + hangmanHeight
		hangmanByStatus[i+1] = hangmanStatContentSplited[currentMin:currentMax]
	}

	_, _, asciiFileName := GetConfigItem(ConfigASCIIFile)

	content, err = os.ReadFile(asciiFileName)
	if err != nil {
		println("Error! Unable to load ascii file '" + fileName + "'")
		os.Exit(1)
	}

	asciiCharacterContentSplited := strings.Split(string(content), "\n")
	for i := 0; i < 127-32; i++ {
		asciiHeight, _, _ := GetConfigItem(ConfigASCIIHeight)
		currentMin := i * asciiHeight
		currentMax := currentMin + asciiHeight
		asciiByChar[rune(i+32)] = asciiCharacterContentSplited[currentMin:currentMax]
	}
	if !FromSave {
		wordToFind := strings.ReplaceAll(GetRandomWord(words), "\r", "")
		wordToFind = strings.ReplaceAll(wordToFind, "\n", "")
		SetGameToFindEncoded(EncodeStrInBase64(string(GetEncodedStringInSha256(wordToFind))) + " # Lol tu le trouvra pas !")
		SetGameToFind(wordToFind)
		SetGameWord(SetupGameWord(wordToFind))
	} else {
		for _, word := range GetCacheWordList() {
			word = strings.ReplaceAll(word, "\r", "")
			word = strings.ReplaceAll(word, "\n", "")
			if EncodeStrInBase64(string(GetEncodedStringInSha256(word))) == strings.Split(GetGameToFindEncoded(), " # Lol tu le trouvra pas !")[0] {
				SetGameToFind(word)
				return
			}
		}
		println("Unable to load, save corrupted..")
		os.Exit(1)
	}
}

func GetCacheHangmanByIndex(index int) []string {
	return hangmanByStatus[index]
}

func GetCacheWordList() []string {
	return words
}

func GetASCIIArtFromRune(r rune) []string {
	return asciiByChar[r]
}
