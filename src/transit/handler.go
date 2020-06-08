package transit

import (
	"fmt"

	"github.com/L04DB4L4NC3R/jokes-rss-bot/src/feed"
	"github.com/yanzay/tbot/v2"
)

func HandleBot(bot *tbot.Server, jokesMessenger Messenger, jokesFeed feed.Feeder) {

	// handle jokes
	bot.HandleMessage(".*joke.*", func(m *tbot.Message) {
		jokes, err := jokesFeed.FetchFeed()
		if err != nil {
			jokesMessenger.Apologize(m.Chat.ID)
			return
		}
		if len(jokes) == 0 {
			fmt.Println("Caught Up")
			err := jokesMessenger.CaughtUp(m.Chat.ID)
			if err != nil {
				fmt.Println(err)
			}
		}
		for _, joke := range jokes {
			err := jokesMessenger.Send(m.Chat.ID, joke)
			if err != nil {
				fmt.Println(err)
			}
		}
	})

	// handle greeting
	bot.HandleMessage("/hi", func(m *tbot.Message) {
		if err := jokesMessenger.Greet(m.Chat.ID); err != nil {
			fmt.Println(err)
		}
	})

	// handle apology
	bot.HandleMessage("/sorry", func(m *tbot.Message) {
		if err := jokesMessenger.Apologize(m.Chat.ID); err != nil {
			fmt.Println(err)
		}
	})
}
