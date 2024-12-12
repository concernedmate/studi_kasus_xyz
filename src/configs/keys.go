package configs

import (
	"log"
	"os"
)

var ACCESS_SECRET string

func Init() {
	priv, err := os.ReadFile(os.Getenv("PRIVATE_KEY_LOCATION"))
	if err != nil {
		log.Fatal("Error reading auth private key!", err)
	}
	ACCESS_SECRET = string(priv)
}

func GetAccessKey() string {
	return ACCESS_SECRET
}
