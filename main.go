package main

import (
	"net/http"
	"os"
	"time"

	cronjob "github.com/L04DB4L4NC3R/jokes-rss-bot/src/crons"
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
		os.Getenv("BOTNAME"), os.Getenv("GROUP_ID"), bot.Client())

	// create functional layer
	jokesFeed := feed.NewJokesFeed(os.Getenv("RSS"), os.Getenv("BOTNAME"), time.Second*60)

	// connect to the Database
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGODB_URI")))

	if err != nil {
		log.Fatalf("Error connecting to the database: %t", err)
	}

	jokesRepo := transit.NewMongoRepo(client.Database("chutkula").Collection("feed_state"))

	// handle transit
	transit.HandleBot(ctx, bot, jokesMessenger, jokesFeed, &jokesRepo)

	// start CRON Jobs
	dailcron := cronjob.FeedUpdate(os.Getenv("GROUP_ID"), jokesMessenger, jokesFeed)
	dailcron.Start()

	// start the worker
	log.Infoln("Starting Bot")
	go bot.Start()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Pong"))
	})
	if err := http.ListenAndServe(":"+os.Getenv("PORT"), nil); err != nil {
		dailcron.Stop()
		cancel()
		log.Fatalln(err)
	}
}
