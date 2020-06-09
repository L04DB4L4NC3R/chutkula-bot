package transit

import (
	"github.com/L04DB4L4NC3R/jokes-rss-bot/src/feed"
	log "github.com/sirupsen/logrus"
	"github.com/yanzay/tbot/v2"
)

func HandleBot(bot *tbot.Server, jokesMessenger Messenger, jokesFeed feed.Feeder) {

	// handle jokes
	bot.HandleMessage(".*joke.*", func(m *tbot.Message) {
		log.Infof("Recieved From: %s At %t ChatID: %s", m.From, m.Date, m.Chat.ID)
		jokes, err := jokesFeed.FetchFeed()
		if err != nil {
			log.Errorf("Handle failed with error %t", err)
			jokesMessenger.Apologize(m.Chat.ID)
			return
		}
		if len(jokes) == 0 {
			log.Infof("All caught up")
			err := jokesMessenger.CaughtUp(m.Chat.ID)
			if err != nil {
				log.Errorf("Handle failed while sending affirmation, error %t", err)
			}
		}

		var (
			errcount     = 0
			successcount = 0
		)
		for _, joke := range jokes {
			err := jokesMessenger.Send(m.Chat.ID, joke)
			if err != nil {
				log.Errorf("Handle failed while sending feed: %t, error %t", joke, err)
				errcount++
			} else {
				successcount++
			}
		}
		log.Infof("Total Feed: %d Feed Sent: %d Feed Failed: %d", errcount+successcount, successcount, errcount)
	})

	// handle greeting
	bot.HandleMessage("/hi", func(m *tbot.Message) {
		if err := jokesMessenger.Greet(m.Chat.ID); err != nil {
			log.Errorf("Handle failed while sending affirmation, error %t", err)
		}
		log.Infof("Sent greeting")
	})

	// handle apology
	bot.HandleMessage("/sorry", func(m *tbot.Message) {
		if err := jokesMessenger.Apologize(m.Chat.ID); err != nil {
			log.Errorf("Handle failed while sending affirmation, error %t", err)
		}
		log.Infof("Sent Apology")
	})

	// get all jokes and don't sync time
	bot.HandleMessage("/lol", func(m *tbot.Message) {
		log.Infof("Recieved From: %s At %t ChatID: %s", m.From, m.Date, m.Chat.ID)
		jokes, err := jokesFeed.FetchFeedUnSync()
		if err != nil {
			log.Errorf("Handle failed with error %t", err)
			jokesMessenger.Apologize(m.Chat.ID)
			return
		}
		if len(jokes) == 0 {
			log.Infof("All caught up")
			err := jokesMessenger.CaughtUp(m.Chat.ID)
			if err != nil {
				log.Errorf("Handle failed while sending affirmation, error %t", err)
			}
		}

		var (
			errcount     = 0
			successcount = 0
		)
		for _, joke := range jokes {
			err := jokesMessenger.Send(m.Chat.ID, joke)
			if err != nil {
				log.Errorf("Handle failed while sending feed: %t, error %t", joke, err)
				errcount++
			} else {
				successcount++
			}
		}
		log.Infof("Total Feed: %d Feed Sent: %d Feed Failed: %d", errcount+successcount, successcount, errcount)
	})

}
