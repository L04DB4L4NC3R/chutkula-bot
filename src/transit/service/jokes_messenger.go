package service

import (
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/yanzay/tbot/v2"
)

type JokesMessenger struct {
	Greeting       string
	Apology        string
	Name           string
	GroupID        string
	LastSentAt     *time.Time
	TelegramClient *tbot.Client
}

func NewJokesMessenger(greeting, apology, name, groupID string, tbc *tbot.Client) Messenger {
	return &JokesMessenger{
		Greeting:       greeting,
		Name:           name,
		GroupID:        groupID,
		LastSentAt:     nil,
		TelegramClient: tbc,
	}
}

func (j *JokesMessenger) Send(chatID string, message string) error {
	_, err := j.TelegramClient.SendMessage(chatID, message)

	return err
}

func (j *JokesMessenger) Greet(chatID string) error {
	log.Infof("Message body: %s", j.Greeting)
	_, err := j.TelegramClient.SendMessage(chatID, j.Greeting)

	return err
}

func (j *JokesMessenger) Apologize(chatID string) error {
	log.Infof("Message body: %s", j.Apology)
	_, err := j.TelegramClient.SendMessage(chatID, j.Apology)

	return err
}

func (j *JokesMessenger) SendGroup(message string) error {
	_, err := j.TelegramClient.SendMessage(j.GroupID, message)

	return err
}

func (j *JokesMessenger) CaughtUp(chatID string) error {
	_, err := j.TelegramClient.SendMessage(chatID, "You are all caught up!")

	return err

}
