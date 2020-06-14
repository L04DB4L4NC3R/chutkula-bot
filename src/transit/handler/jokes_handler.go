package handler

import (
	"context"

	"github.com/L04DB4L4NC3R/jokes-rss-bot/src/feed"
	repo "github.com/L04DB4L4NC3R/jokes-rss-bot/src/transit/repository"
	"github.com/L04DB4L4NC3R/jokes-rss-bot/src/transit/service"
	log "github.com/sirupsen/logrus"
	"github.com/yanzay/tbot/v2"
)

func NewJokesHandler(bot *tbot.Server, jokesMessenger service.Messenger, jokesFeed feed.Feeder, jokesRepo repo.Repository) Handler {
	return &Handle{
		bot:       bot,
		messenger: jokesMessenger,
		feed:      jokesFeed,
		repo:      jokesRepo,
	}
}

func (h *Handle) HandleBot() {
	h.bot.HandleMessage("/hi", h.Greet)
	h.bot.HandleMessage("/sorry", h.Apologize)
	h.bot.HandleMessage("/caughtup", h.CaughtUp)
	h.bot.HandleMessage("/register", h.Register)
	h.bot.HandleMessage("/unregister", h.UnRegister)
	h.bot.HandleMessage("/time", h.GetMeta)
	h.bot.HandleMessage("/jokes", h.MainFunc)
	h.bot.HandleMessage("/lol", h.Lol)
}
func (h *Handle) Greet(m *tbot.Message) {
	if err := h.messenger.Greet(m.Chat.ID); err != nil {
		log.Errorf("Handle failed while sending affirmation, error %s", err.Error())
	} else {
		log.Infof("Sent greeting")
	}
}

func (h *Handle) Apologize(m *tbot.Message) {
	if err := h.messenger.Apologize(m.Chat.ID); err != nil {
		log.Errorf("Handle failed while sending affirmation, error %s", err.Error())
	}
	log.Infof("Sent Apology")
}

func (h *Handle) CaughtUp(m *tbot.Message) {
	if err := h.messenger.Send(m.Chat.ID, "All Caught Up!"); err != nil {
		log.Errorf("Handle failed while sending affirmation, error %s", err.Error())
	}
	log.Infof("Sent Apology")
}

func (h *Handle) Register(m *tbot.Message) {
	_, updatedAt, err := h.feed.FetchFeedUnSync()
	if err != nil {
		log.Errorf("Handle failed with error %s", err.Error())
		h.messenger.Apologize(m.Chat.ID)
		return
	}
	ctx := context.Background()
	if err = h.repo.Register(ctx, m.Chat.ID, h.feed.GetFeedName(), updatedAt); err != nil {
		if err.Error() == "Document Already Exists" {
			log.Info("Already registered")
			h.messenger.Send(m.Chat.ID, "Already Registered!")
			return
		} else {
			log.Errorf("Handle failed with error %s", err.Error())
			h.messenger.Apologize(m.Chat.ID)
			return
		}
	}
	if err := h.messenger.Send(m.Chat.ID, "Registered!"); err != nil {
		log.Errorf("Handle failed while sending affirmation, error %s", err.Error())
	} else {
		log.Infof("Sent Message")
	}

}

func (h *Handle) UnRegister(m *tbot.Message) {
	if err := h.repo.UnRegister(context.Background(), m.Chat.ID, h.feed.GetFeedName()); err != nil {
		if err.Error() == "No Chat Found" {
			log.Info("Already unregistered")
			h.messenger.Send(m.Chat.ID, "Already Un-Registered!")
			return
		} else {
			log.Errorf("Handle failed with error %s", err.Error())
			h.messenger.Apologize(m.Chat.ID)
			return
		}
	}
	if err := h.messenger.Send(m.Chat.ID, "Un-Registered!"); err != nil {
		log.Errorf("Handle failed while sending affirmation, error %s", err.Error())
	} else {
		log.Infof("Sent Message")
	}
}

func (h *Handle) GetMeta(m *tbot.Message) {
	ctx := context.Background()
	t, err := h.repo.GetUpdatedAt(ctx, m.Chat.ID, h.feed.GetFeedName())

	if err != nil {
		log.Errorf("Handle failed while sending affirmation, error %s", err.Error())
	}
	if err = h.messenger.Send(m.Chat.ID, t.String()); err != nil {
		log.Errorf("Handle failed while sending affirmation, error %s", err.Error())
	}
}
func (h *Handle) MainFunc(m *tbot.Message) {
	log.Infof("Recieved From: %s At %t ChatID: %s", m.From, m.Date, m.Chat.ID)
	// get last updated timestamp
	ts, err := h.repo.GetUpdatedAt(context.Background(), m.Chat.ID, h.feed.GetFeedName())
	if err != nil {
		log.Errorf("Handle failed with error %s", err.Error())
		h.messenger.Apologize(m.Chat.ID)
		return
	}
	jokes, newtime, err := h.feed.FetchFeed(ts)
	if err != nil {
		log.Errorf("Handle failed with error %s", err.Error())
		h.messenger.Apologize(m.Chat.ID)
		return
	}
	if len(jokes) == 0 {
		log.Infof("All caught up")
		err := h.messenger.CaughtUp(m.Chat.ID)
		if err != nil {
			log.Errorf("Handle failed while sending affirmation, error %s", err.Error())
		}
	}

	var (
		errcount     = 0
		successcount = 0
	)
	for _, joke := range jokes {
		err := h.messenger.Send(m.Chat.ID, joke)
		if err != nil {
			log.Errorf("Handle failed while sending feed: %t, error %s", joke, err.Error())
			errcount++
		} else {
			successcount++
		}
	}
	log.Infof("Total Feed: %d Feed Sent: %d Feed Failed: %d", errcount+successcount, successcount, errcount)

	// update the new time
	if err = h.repo.UpdateTimeStamp(context.Background(), newtime, m.Chat.ID, h.feed.GetFeedName()); err != nil {
		log.Errorf("Update timestamp Handle failed with error %s", err.Error())
		h.messenger.Send(m.Chat.ID, "Timestamp Could Not be Updated")
		return
	}
}

func (h *Handle) Lol(m *tbot.Message) {
	log.Infof("Recieved From: %s At %t ChatID: %s", m.From, m.Date, m.Chat.ID)
	jokes, _, err := h.feed.FetchFeedUnSync()
	if err != nil {
		log.Errorf("Handle failed with error %s", err.Error())
		h.messenger.Apologize(m.Chat.ID)
		return
	}
	if len(jokes) == 0 {
		log.Infof("All caught up")
		err := h.messenger.CaughtUp(m.Chat.ID)
		if err != nil {
			log.Errorf("Handle failed while sending affirmation, error %s", err.Error())
		}
	}

	var (
		errcount     = 0
		successcount = 0
	)
	for _, joke := range jokes {
		err := h.messenger.Send(m.Chat.ID, joke)
		if err != nil {
			log.Errorf("Handle failed while sending feed: %t, error %s", joke, err.Error())
			errcount++
		} else {
			successcount++
		}
	}
	log.Infof("Total Feed: %d Feed Sent: %d Feed Failed: %d", errcount+successcount, successcount, errcount)
}
