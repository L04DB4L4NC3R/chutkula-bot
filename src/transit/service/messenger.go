package service

type Messenger interface {
	Send(chatID string, message string) error
	Greet(chatID string) error
	SendGroup(message string) error
	Apologize(chatID string) error
	CaughtUp(chatID string) error
}
