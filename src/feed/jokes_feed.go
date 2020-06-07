package feed

import (
	"context"
	"fmt"
	"strings"
	"time"

	strip "github.com/grokify/html-strip-tags-go"
	"github.com/mmcdole/gofeed"
)

type JokesFeed struct {
	Url          string
	EmojiList    []string
	FeedName     string
	FetchTimeout time.Duration
}

func NewJokesFeed(url, feedname string, fetchTimeout time.Duration) Feeder {
	return &JokesFeed{
		Url:          url,
		FeedName:     feedname,
		EmojiList:    []string{},
		FetchTimeout: fetchTimeout,
	}
}

func (j *JokesFeed) ParseContent(content string, title string) (parsedItem string) {
	replacer := strings.NewReplacer("&quot", "", "&#32", "", ";", "", "[link]", "", "[comments", "", "submitted by", "\n\nsubmitted by: ", "]", "", "&#39", "")
	content = replacer.Replace(strip.StripTags(content))

	// TODO: use random emoji
	return fmt.Sprintf("👽🤣😋😏 \n\n%s\n\n%s\n\n😂😳😛😵", title, content)
}

func (j *JokesFeed) FetchFeed() (items []string, err error) {
	// set 60 second timeout
	ctx, cancel := context.WithTimeout(context.Background(), j.FetchTimeout)
	defer cancel()

	// parse the reddit jokes feed
	fp := gofeed.NewParser()
	feed, err := fp.ParseURLWithContext(j.Url, ctx)

	if err != nil {
		return nil, err
	}

	// TODO: check last updated
	// if time after last updated then run the following
	// else return both nil
	var reply []string
	var content string
	for _, i := range feed.Items {
		// grab items in p tag
		content = j.ParseContent(i.Content, i.Title)
		reply = append(reply, content)
	}
	return reply, nil
}

func (j *JokesFeed) IsSyncedTime(updatedTime *time.Time) (bool, error) {
	panic("not implemented") // TODO: Implement
}

func (j *JokesFeed) EmojiInjector() (emojis []string) {
	panic("not implemented") // TODO: Implement
}
