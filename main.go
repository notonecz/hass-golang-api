package main

import (
	"hass-golang-api/rest"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/joho/godotenv"
)

func main() {

	_ = godotenv.Load()
	token := os.Getenv("TOKEN")

	home1 := rest.Init("HOME1", token, "homeassistant.local:8123", false)

	a := app.New()
	w := a.NewWindow("Garaz")

	w.Resize(fyne.NewSize(800, 600))

	w.SetContent(container.NewVBox(
		widget.NewLabel("Ovládání Garáže!"),
		widget.NewButton("Ovladac!", func() {
			_, _ = rest.PostService(home1, "input_button", "press", "{\"entity_id\": \"input_button.garaz\"}")
		}),
		widget.NewLabel("Světlo kuchyň!"),
		widget.NewButton("Zapnout", func() {
			_, _ = rest.PostService(home1, "light", "turn_on", "{\"entity_id\": \"light.mini\"}")
		}),
		widget.NewButton("Vypnout", func() {
			_, _ = rest.PostService(home1, "light", "turn_off", "{\"entity_id\": \"light.mini\"}")
		}),
	))

	w.ShowAndRun()
}
