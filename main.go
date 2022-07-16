package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
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

	hello := widget.NewLabel("Welcome to gothic, enter your secret key and press the button below to generate the code.")

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

	w.SetContent(container.NewVBox(
		hello,
		widget.NewButton("Refresh", func() {
			hello.Text = ""
			for _, token := range tokens {
				code, err := totp.GenerateCode(token.Secret, time.Now())
				if err != nil {
					log.Fatal(err)
				}
				hello.SetText(hello.Text + "\n" + token.Name + ": " + code)
			}

		}),
	))
	w.Resize(fyne.NewSize(400, 400))

	w.ShowAndRun()
}
