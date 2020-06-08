package transit

import (
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/yanzay/tbot/v2"
)

func NewTelegramServer() *tbot.Server {
	bot := tbot.New(os.Getenv("BOT_TOKEN"))
	log.Infof("Created new bot server")
	return bot
}
