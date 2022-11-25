package main

import (
	"fmt"
	"handlers"
	"net/http"
	"objects"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	StartServer()
}

func StartServer() {
	handlers.InitWebServer()
	objects.LoadUserCSV()

	go autoSaveWorker()
	sigchnl := make(chan os.Signal, 1)
	signal.Notify(sigchnl)

	go func() {
		for {
			s := <-sigchnl
			osSignalHandler(s)
		}
	}()

	http.ListenAndServe(":8080", nil)
}

func osSignalHandler(signal os.Signal) {
	if signal == syscall.SIGTERM || signal == syscall.SIGINT {
		fmt.Println("Program will terminate now.")
		objects.SaveUserCSV()
		os.Exit(0)
	}
}
func autoSaveWorker() {
	for {
		time.Sleep(time.Second * 10)
		objects.SaveUserCSV()
	}
}
