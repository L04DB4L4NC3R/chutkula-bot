package repo

import (
	"context"
	"time"
)

type Repository interface {
	Register(ctx context.Context, chatID string, feed string, updatedAt *time.Time) error

	UnRegister(ctx context.Context, chatID string, feed string) error

	GetUpdatedAt(ctx context.Context, chatID string, feed string) (*time.Time, error)

	UpdateTimeStamp(ctx context.Context, newtime *time.Time, chatID string, feed string) error

	GetUpdatedStates(ctx context.Context) ([]State, error)
}

type State struct {
	ChatID    string     `json:"chat_id" bson:"chat_id"`
	Feed      string     `json:"feed" bson:"feed"`
	UpdatedAt *time.Time `json:"updated_at" bson:"updated_at"`
}
