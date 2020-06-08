package transit

import (
	"fmt"
	"os"

	"github.com/L04DB4L4NC3R/jokes-rss-bot/src/feed"
	"github.com/yanzay/tbot/v2"
)

func NewTelegramServer() *tbot.Server {
	bot := tbot.New(os.Getenv("BOT_TOKEN"))
	return bot
}

func HandleBot(bot *tbot.Server, jokesMessenger Messenger, jokesFeed feed.Feeder) {
	bot.HandleMessage(".*joke.*", func(m *tbot.Message) {
		jokes, err := jokesFeed.FetchFeed()
		if err != nil {
			jokesMessenger.Apologize(m.Chat.ID)
			return
		}
		for _, joke := range jokes {
			err := jokesMessenger.Send(m.Chat.ID, joke)
			if err != nil {
				fmt.Println(err)
			}
		}
	})
}
