package handler

import (
	"github.com/L04DB4L4NC3R/jokes-rss-bot/src/feed"
	repo "github.com/L04DB4L4NC3R/jokes-rss-bot/src/transit/repositorie"
	"github.com/L04DB4L4NC3R/jokes-rss-bot/src/transit/service"
	"github.com/yanzay/tbot/v2"
)

type Handler interface {
	Greet(*tbot.Message)
	Apologize(*tbot.Message)
	CaughtUp(*tbot.Message)
	Register(*tbot.Message)
	UnRegister(*tbot.Message)
	GetMeta(*tbot.Message)
	MainFunc(*tbot.Message)
	HandleBot()
}

type Handle struct {
	bot       *tbot.Server
	messenger service.Messenger
	feed      feed.Feeder
	repo      repo.Repository
}
