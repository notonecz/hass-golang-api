package main

import (
	"hass-golang-api/rest"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load()
	token := os.Getenv("TOKEN")

	home1 := rest.Init("HOME1", token, "homeassistant.local:8123")

	state, _ := rest.GetState(home1, "sensor.alarm_clock_24f6_temperature")

	print(string(state))

	garage, _ := rest.PostService(home1, "input_button", "press", "{\"entity_id\": \"input_button.garaz\"}")

	print(string(garage))
}
