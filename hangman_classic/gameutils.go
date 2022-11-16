package hangman_classic

import (
	"crypto/sha256"
	"encoding/base64"
	"math/rand"
	"strings"
)

func GetRandomWord(liste []string) string {
	return liste[rand.Intn(len(liste))]
}

func HasOccurenceLetter(word string, letterToCheck rune) bool {
	for _, letter := range word {
		if letter == letterToCheck {
			return true
		}
	}
	return false
}

func GetOccurenceLetter(word string, letterToCheck rune) []int {
	var occ []int
	wordR := []rune(word)
	for i, letter := range wordR {
		if letter == letterToCheck {
			occ = append(occ, i)
		}
	}
	return occ
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

func isVowel(r rune) bool {
	return strings.ContainsRune("aeiouy", r)
}
func VowelCount(str string) int {
	count := 0
	for _, r := range str {
		if isVowel(r) {
			count++
		}
	}
	return count
}

func (c *Gamecache) BuildASCIIWord(word string) []string {
	words := make([]string, 9)
	for _, runes := range word {
		for i, line := range c.AsciiByChar[runes] {
			for _, r := range line {
				if r > 31 && r < 126 {
					words[i] = words[i] + string(r)
				}
			}
		}
	}
	return words
}

func ConvertToUnicode(s string) string {
	newWord := ""
	for _, r := range s {
		switch r {
		case 'á', 'à', 'â', 'ä', 'ã', 'å':
			newWord += "a"
		case 'Á', 'À', 'Â', 'Ä', 'Ã', 'Å':
			newWord += "A"
		case 'æ':
			newWord += "ae"
		case 'Æ':
			newWord += "AE"
		case 'ç':
			newWord += "c"
		case 'Ç':
			newWord += "C"
		case 'é', 'è', 'ê', 'ë':
			newWord += "e"
		case 'É', 'È', 'Ê', 'Ë':
			newWord += "E"
		case 'í', 'ì', 'î', 'ï':
			newWord += "i"
		case 'Í', 'Ì', 'Î', 'Ï':
			newWord += "I"
		case 'ñ':
			newWord += "n"
		case 'Ñ':
			newWord += "N"
		case 'ó', 'ò', 'ô', 'ö', 'õ', 'ø':
			newWord += "o"
		case 'Ó', 'Ò', 'Ô', 'Ö', 'Õ', 'Ø':
			newWord += "O"
		case 'œ':
			newWord += "oe"
		case 'Œ':
			newWord += "OE"
		case 'ú', 'ù', 'û', 'ü':
			newWord += "u"
		case 'Ú', 'Ù', 'Û', 'Ü':
			newWord += "U"
		default:
			newWord += string(r)
		}
	}
	return newWord
}

func EncodeStrInBase64(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}

func DecodeStrInBase64(str string) string {
	decoded, _ := base64.StdEncoding.DecodeString(str)
	return string(decoded)
}

func GetEncodedStringInSha256(str string) []byte {
	h := sha256.New()
	h.Write([]byte(str))
	return h.Sum([]byte{})
}
