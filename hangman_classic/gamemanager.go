package hangman_classic

import (
	"encoding/json"
	"os"
)

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

func (game *HangmanGame) GetGameToFindEncoded() string {
	return game.EncodedToFind
}

func (game *HangmanGame) SetGameToFindEncoded(tofind string) string {
	game.EncodedToFind = tofind
	return game.tofind
}

func (game *HangmanGame) GetGameWord() string {
	return game.Word
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
	_, _, fileName := game.Config.GetConfigItem(ConfigSaveFile)
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

func (game *HangmanGame) LoadSave(fileName string) (HangmanGame, error) {
	content, err := os.ReadFile(fileName)
	if err != nil {
		println("Error: File " + fileName + " doesn't exist ! please specify a existing one.")
		return HangmanGame{}, err
	}

	Game := HangmanGame{}

	er := json.Unmarshal([]byte(DecodeStrInBase64(string(content))), Game)

	if er != nil {
		println("Error: File " + fileName + " doesn't exist ! please specify a existing one.")
		return HangmanGame{}, err
	}

	FromSave = true
	return Game, nil
}

func (g *HangmanGame) GetConfig() *GameConfig {
	return &g.Config
}
