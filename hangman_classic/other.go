package hangman_classic

import (
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"time"
)

func InitEnvironement() {
	rand.Seed(time.Now().UnixMicro())
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
