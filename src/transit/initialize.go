package transit

import (
	"os"

	"github.com/yanzay/tbot/v2"
)

func NewTelegramServer() *tbot.Server {
	bot := tbot.New(os.Getenv("BOT_TOKEN"))
	return bot
}
