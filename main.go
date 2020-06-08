package main

import (
	"os"
	"time"

	"github.com/L04DB4L4NC3R/jokes-rss-bot/src/feed"
	"github.com/L04DB4L4NC3R/jokes-rss-bot/src/transit"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

func initialize() {

	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)

	if err := godotenv.Load(); err != nil {
		log.Warnf("Error loading .env file: %t", err)
	}

}

func main() {

	// initialize logger and environment variables
	initialize()

	// create telegram bot
	bot := transit.NewTelegramServer()

	// create transit layer
	jokesMessenger := transit.NewJokesMessenger(os.Getenv("GREETING"), os.Getenv("APOLOGY"),
		os.Getenv("BOTNAME"), os.Getenv("GROUPID"), bot.Client())

	// create functional layer
	jokesFeed := feed.NewJokesFeed(os.Getenv("RSS"), os.Getenv("BOTNAME"), time.Second*60)

	// handle transit
	transit.HandleBot(bot, jokesMessenger, jokesFeed)

	// start the worker
	log.Infoln("Starting Bot")
	bot.Start()
}
