package main

import (
	"log"
	"os"
	"time"

	"github.com/L04DB4L4NC3R/jokes-rss-bot/src/feed"
	"github.com/L04DB4L4NC3R/jokes-rss-bot/src/transit"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Error loading .env file")
	}

	bot := transit.NewTelegramServer()

	greeting := "Hey, There!"
	apology := "Sorry, Could not be completed!"
	botname := "chutkulabot"

	jokesMessenger := transit.NewJokesMessenger(greeting, apology, botname, os.Getenv("TELEGRAM_GROUPID"), bot.Client())

	jokesFeed := feed.NewJokesFeed(os.Getenv("JOKES_RSS"), botname, time.Second*60)

	transit.HandleBot(bot, jokesMessenger, jokesFeed)

	log.Println("Starting Bot")
	bot.Start()
}
