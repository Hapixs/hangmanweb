package hangman_classic

import (
	"encoding/json"
	"os"
)

var game *HangmanGame = &HangmanGame{}

func GetGame() *HangmanGame {
	return game
}

func SetGame(HangmanGame *HangmanGame) {
	game = HangmanGame
}

func AddGameTry() int {
	game.Tries++
	if game.Tries > maxTries {
		game.Tries = maxTries
	}
	return game.Tries
}

func GetGameTries() int {
	return game.Tries
}

func SetGameWord(word string) string {
	game.Word = word
	return game.Word
}

func GetGameToFindEncoded() string {
	return game.EncodedToFind
}

func SetGameToFindEncoded(tofind string) string {
	game.EncodedToFind = tofind
	return game.tofind
}

func GetGameWord() string {
	return game.Word
}

func SetGameToFind(tofind string) string {
	game.tofind = tofind
	return game.tofind
}

func GetGameToFind() string {
	_, b, _ := GetConfigItem(ConfigUseAscii)
	if b {
		return ConvertToUnicode(game.tofind)
	}
	return game.tofind
}

func GetGameUsed() string {
	return game.Used
}

func AddGameUsed(r rune) string {
	game.Used += string(r)
	return game.Used
}

func GetGameSave() GameSave {
	return GameSave{*game, GetGameConfig()}
}

func SaveGame() {
	_, _, fileName := GetConfigItem(ConfigSaveFile)
	saveEnc, err := json.Marshal(GetGameSave())
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

func LoadSave(fileName string) error {
	content, err := os.ReadFile(fileName)
	if err != nil {
		println("Error: File " + fileName + " doesn't exist ! please specify a existing one.")
		return err
	}

	gameSave := &GameSave{}

	er := json.Unmarshal([]byte(DecodeStrInBase64(string(content))), gameSave)

	if er != nil {
		println("Error: File " + fileName + " doesn't exist ! please specify a existing one.")
		return err
	}

	SetGame(&gameSave.Game)
	SetGameConfig(gameSave.Config)
	FromSave = true
	return nil
}

func ResetTempData() {
	game.Used = ""
}
