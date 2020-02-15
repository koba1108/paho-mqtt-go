package main

import (
	"github.com/joho/godotenv"
	"github.com/koba1108/paho-mqtt-go/internal/subscriber"
	"os"
)

func main() {
	if os.Getenv("APP_ENV") != "production" {
		if err := godotenv.Load(".env"); err != nil {
			panic(err)
		}
	}
	subscriber.Run()
}
