package cronjob

import (
	"context"

	"github.com/L04DB4L4NC3R/jokes-rss-bot/src/feed"
	repo "github.com/L04DB4L4NC3R/jokes-rss-bot/src/transit/repository"
	"github.com/L04DB4L4NC3R/jokes-rss-bot/src/transit/service"
	log "github.com/sirupsen/logrus"
	"gopkg.in/robfig/cron.v2"
)

func FeedUpdate(messenger service.Messenger, feed feed.Feeder, repo repo.Repository) *cron.Cron {
	c := cron.New()

	c.AddFunc("@every 0h45m0s", func() {

		log.Infof("Running CronJob")
		// get all registered chats

		chats, err := repo.GetUpdatedStates(context.Background())
		if err != nil {
			log.Infof("Error fetching chats...%s", err.Error())
			return
		}

		// fetch raw feed
		raw, err := feed.FetchRawFeed()
		if err != nil {
			log.Errorf("Handle failed with error %t", err)
			return
		}
		for _, chat := range chats {

			// parse raw feed and determine which ones to show
			var jokes []string
			var content string
			var uptodatecount = 0
			for _, i := range raw.Items {
				if !feed.IsSyncedTime(i.UpdatedParsed, chat.UpdatedAt) {
					uptodatecount++
					continue
				}
				content = feed.ParseContent(i.Content, i.Title)
				jokes = append(jokes, content)
			}

			log.Infof("FOR CHATID: %s ::: Succeeded fetching feed. Items: %d. Updated: %s. Up to date count: %d. New feed count: %d", chat.ChatID, len(raw.Items), raw.Updated, uptodatecount, len(jokes))

			if len(jokes) == 0 {
				log.Infof("All caught up")
				/* err := messenger.CaughtUp(chat.ChatID)
				if err != nil {
					log.Errorf("Handle failed while sending affirmation, error %t", err)
				} */
				continue
			}

			var (
				errcount     = 0
				successcount = 0
			)
			for _, joke := range jokes {
				err := messenger.Send(chat.ChatID, joke)
				if err != nil {
					log.Errorf("Handle failed while sending feed: %t, error %t", joke, err)
					errcount++
				} else {
					successcount++
				}
			}
			log.Infof("Total Feed: %d Feed Sent: %d Feed Failed: %d", errcount+successcount, successcount, errcount)

			// update the new time
			if err = repo.UpdateTimeStamp(context.Background(), raw.UpdatedParsed, chat.ChatID, feed.GetFeedName()); err != nil {
				log.Errorf("Update timestamp Handle failed with error %t", err)
				messenger.Send(chat.ChatID, "Timestamp Could Not be Updated")
				return
			}
		}
	})

	return c
}
