package feed

import (
	"time"

	"github.com/mmcdole/gofeed"
)

type Feeder interface {
	ParseContent(content string, title string, link string) (parsedItem string)

	FetchFeed(lastUpdatedAt *time.Time) (items []string, newtime *time.Time, err error)

	FetchFeedUnSync() (items []string, updatedAt *time.Time, err error)

	IsSyncedTime(updatedTime *time.Time, lastUpdatedAt *time.Time) bool

	EmojiInjector(num int) (emojis []string)

	GetFeedName() string

	FetchRawFeed() (*gofeed.Feed, error)
}
