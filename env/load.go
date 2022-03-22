package env

import (
	"os"

	"github.com/joho/godotenv"
)

var Homeserver string
var Username string
var Password string

var RoomId string

var WebserverHost string
var WebserverPort string

var PresharedKey string

func LoadEnv() {
	// Load env variables
	godotenv.Load(".env")

	Homeserver = getenv("MATRIX_HOMESERVER", "https://matrix.org")
	Username = getenv("MATRIX_USERNAME", "")
	Password = getenv("MATRIX_PASSWORD", "")

	RoomId = getenv("MATRIX_ROOM_ID", "")

	WebserverHost = getenv("WEBSERVER_HOST", "localhost")
	WebserverPort = getenv("WEBSERVER_PORT", "8080")

	PresharedKey = getenv("WEBSERVER_PRESHARED_KEY", "none")
}

func getenv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 && fallback == "" {
		panic("Environment variable " + key + " is not set")
	} else if len(value) == 0 {
		return fallback
	}
	return value
}
