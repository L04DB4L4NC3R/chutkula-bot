package main

import (
	"log"
	"net/http"
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
	botname := "ChutkulaBot"

	jokesMessenger := transit.NewJokesMessenger(greeting, apology, botname, os.Getenv("TELEGRAM_GROUPID"), bot.Client())

	jokesFeed := feed.NewJokesFeed(os.Getenv("JOKES_RSS"), botname, time.Second*60)

	transit.HandleBot(bot, jokesMessenger, jokesFeed)

	log.Println("Starting Bot")
	go bot.Start()

	// due to heroku build process not being
	// able to startup worker properly
	// this is needed
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}
