package hangman_classic

import (
	"strings"
)

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

func BuildASCIIWord(word string) []string {
	words := make([]string, 9)
	for _, runes := range word {
		for i, line := range GetASCIIArtFromRune(runes) {
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
