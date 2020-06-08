package transit

import "time"

type Messenger interface {
	Send(chatID string, message string) error
	Greet(chatID string) error
	SendGroup(message string) error
	Apologize(chatID string) error
	SyncTime(newFetchedAt *time.Time) error
	GetLastSyncTime() (*time.Time, error)
}
