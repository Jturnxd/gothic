package main

import (
	"log"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/pquerna/otp/totp"
)

func main() {
	a := app.New()
	w := a.NewWindow("gothic")

	hello := widget.NewLabel("Welcome to gothic, enter your secret key and press the button below to generate the code.")
	secret := widget.NewEntry()

	w.SetContent(container.NewVBox(
		hello,
		secret,
		widget.NewButton("Generate code", func() {
			code, err := totp.GenerateCode(secret.Text, time.Now())
			if err != nil {
				log.Println(err)
			} else {
				hello.SetText(code)
			}
		}),
	))
	w.Resize(fyne.NewSize(400, 400))

	w.ShowAndRun()
}
