package feed

import (
	"context"
	"fmt"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"

	strip "github.com/grokify/html-strip-tags-go"
	"github.com/mmcdole/gofeed"
)

type JokesFeed struct {
	Url           string
	EmojiList     []string
	BotName       string
	FetchTimeout  time.Duration
	LastUpdatedAt *time.Time
}

func NewJokesFeed(url, botname string, fetchTimeout time.Duration) Feeder {
	return &JokesFeed{
		Url:           url,
		BotName:       botname,
		EmojiList:     []string{},
		FetchTimeout:  fetchTimeout,
		LastUpdatedAt: nil,
	}
}

func (j *JokesFeed) ParseContent(content string, title string) (parsedItem string) {
	replacer := strings.NewReplacer("&quot", "", "&#32", "", ";", "", "[link]", "", "[comments", "", "submitted by", "\n\nsubmitted by: ", "]", "", "&#39", "")
	content = replacer.Replace(strip.StripTags(content))

	// TODO: use random emoji
	return fmt.Sprintf("👽🤣😋😏 \n\n%s\n\n%s\n\n😂😳😛😵\n\nby @%s", title, content, j.BotName)
}

func (j *JokesFeed) FetchFeed() (items []string, err error) {
	// set 60 second timeout
	ctx, cancel := context.WithTimeout(context.Background(), j.FetchTimeout)
	defer cancel()

	// parse the reddit jokes feed
	fp := gofeed.NewParser()
	feed, err := fp.ParseURLWithContext(j.Url, ctx)

	if err != nil {
		log.Errorf("Error fetching feed: %t", err)
		return nil, err
	}

	// if time after last updated then run the following
	// else return both nil
	var reply []string
	var content string
	var uptodatecount = 0
	for _, i := range feed.Items {
		if !j.IsSyncedTime(i.UpdatedParsed) {
			uptodatecount++
			continue
		}
		content = j.ParseContent(i.Content, i.Title)
		reply = append(reply, content)
	}
	log.Infof("Succeeded fetching feed.\nItems: %d\nUpdated: %s\nUp to date count: %d\nNew feed count: %d", len(feed.Items), feed.Updated, uptodatecount, len(reply))
	j.LastUpdatedAt = feed.UpdatedParsed
	return reply, nil
}

func (j *JokesFeed) IsSyncedTime(updatedTime *time.Time) bool {
	if j.LastUpdatedAt == nil || updatedTime.After(*j.LastUpdatedAt) {
		return true
	}
	return false
}

func (j *JokesFeed) EmojiInjector() (emojis []string) {
	panic("not implemented") // TODO: Implement
}
