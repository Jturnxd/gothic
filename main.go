package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
	"github.com/pquerna/otp/totp"
)

type Token struct {
	Name   string
	Secret string
	URI    string
}

func main() {
	a := app.New()
	w := a.NewWindow("gothic")

	config, err := os.Open("config.json")
	if err != nil {
		log.Fatal(err)

	}
	defer config.Close()

	configBytes, err := ioutil.ReadAll(config)
	if err != nil {
		log.Fatal(err)
	}

	var tokens []Token

	err = json.Unmarshal(configBytes, &tokens)
	if err != nil {
		log.Fatal(err)
	}

	go timer(w, tokens)

	w.Resize(fyne.NewSize(400, 400))

	w.ShowAndRun()
}

func refresh(w fyne.Window, tokens []Token) {
	codes := widget.NewLabel("")
	for _, token := range tokens {
		code, err := totp.GenerateCode(token.Secret, time.Now())
		if err != nil {
			log.Fatal(err)
		}
		codes.SetText(codes.Text + "\n" + token.Name + ": " + code)
	}
	w.SetContent(codes)
}

func timer(w fyne.Window, tokens []Token) {
	refresh(w, tokens)
	for {
		if time.Now().Second() == 0 || time.Now().Second() == 30 {
			refresh(w, tokens)
			log.Println("refreshed codes")
		}
		time.Sleep(1 * time.Second)
	}
}
