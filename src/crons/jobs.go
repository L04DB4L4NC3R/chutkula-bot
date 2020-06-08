package cronjob

import (
	"fmt"

	"github.com/L04DB4L4NC3R/jokes-rss-bot/src/feed"
	"github.com/L04DB4L4NC3R/jokes-rss-bot/src/transit"
	log "github.com/sirupsen/logrus"
	"gopkg.in/robfig/cron.v2"
)

func FeedUpdate(chatID string, messenger transit.Messenger, feed feed.Feeder) *cron.Cron {
	c := cron.New()
	c.AddFunc("@every 4h0m0s", func() {

		log.Infof("Running Daily CronJob")
		jokes, err := feed.FetchFeed()
		if err != nil {
			log.Errorf("Handle failed with error %t", err)
			messenger.Apologize(chatID)
			return
		}
		if len(jokes) == 0 {
			fmt.Println("Caught Up")
			log.Infof("All caught up")
			err := messenger.CaughtUp(chatID)
			if err != nil {
				log.Errorf("Handle failed while sending affirmation, error %t", err)
			}
		}

		var (
			errcount     = 0
			successcount = 0
		)
		for _, joke := range jokes {
			err := messenger.Send(chatID, joke)
			if err != nil {
				log.Errorf("Handle failed while sending feed: %t, error %t", joke, err)
				errcount++
			} else {
				successcount++
			}
		}
		log.Infof("Total Feed: %d Feed Sent: %d Feed Failed: %d", errcount+successcount, successcount, errcount)
	})

	return c
}
