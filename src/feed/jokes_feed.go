package feed

import (
	"context"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/L04DB4L4NC3R/jokes-rss-bot/src/static"
	strip "github.com/grokify/html-strip-tags-go"
	log "github.com/sirupsen/logrus"

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

func (j *JokesFeed) ParseContent(content string, title string, link string) (parsedItem string) {

	// replace useless text from the subreddit
	replacer := strings.NewReplacer("&quot", "", "&#32", "", ";", "", "[link]", "", "[comments", "", "submitted by", "\n\nsubmitted by:", "]", "", "&#39", "")

	// check if image exists or not
	var imageurl string
	cont := strings.Split(content, " ")
	n := len(cont)

	for i := 0; i < n-1; i++ {
		if cont[i] == "<img" {
			imageurl = cont[i+1][5 : len(cont[i+1])-1]
		}
	}

	// if image exists then exit out with the image url
	if imageurl != "" {
		// this ensures no image is there, comment for image memes
		return ""
		// return fmt.Sprintf("%s$$%s\n\nView full post at %s\n\nby %s", imageurl, title, link, j.BotName)
	}
	content = replacer.Replace(strip.StripTags(content))

	// inject random emoji
	emj := j.EmojiInjector(8)
	return fmt.Sprintf("%s %s %s %s\n\n%s\n\n%s\n\nView full post at %s\n\n%s %s %s %s\n\nby %s", emj[0], emj[1], emj[2], emj[3], title, content, link, emj[4], emj[5], emj[6], emj[7], "http://kutt.loadbalancer.tech/chutkulabot")
}

func (j *JokesFeed) FetchFeedUnSync() (items []string, updatedAt *time.Time, err error) {

	// set 60 second timeout
	ctx, cancel := context.WithTimeout(context.Background(), j.FetchTimeout)
	defer cancel()

	// parse the reddit jokes feed
	fp := gofeed.NewParser()
	feed, err := fp.ParseURLWithContext(j.Url, ctx)

	if err != nil {
		log.Errorf("Error fetching feed: %t", err.Error())
		return nil, nil, err
	}

	// if time after last updated then run the following
	// else return both nil
	var reply []string
	var content string
	for _, i := range feed.Items {
		content = j.ParseContent(i.Content, i.Title, i.Link)
		if reply != "" {
			reply = append(reply, content)
		}
	}
	log.Infof("Succeeded fetching feed. Items: %d. Updated: %s. New feed count: %d", len(feed.Items), feed.Updated, len(reply))
	return reply, feed.UpdatedParsed, nil
}

func (j *JokesFeed) FetchFeed(lastUpdatedAt *time.Time) (items []string, newtime *time.Time, err error) {
	// set 60 second timeout
	ctx, cancel := context.WithTimeout(context.Background(), j.FetchTimeout)
	defer cancel()

	// parse the reddit jokes feed
	fp := gofeed.NewParser()
	feed, err := fp.ParseURLWithContext(j.Url, ctx)

	if err != nil {
		log.Errorf("Error fetching feed: %t", err.Error())
		return nil, nil, err
	}

	// if time after last updated then run the following
	// else return both nil
	var reply []string
	var content string
	var uptodatecount = 0
	for _, i := range feed.Items {
		if !j.IsSyncedTime(i.UpdatedParsed, lastUpdatedAt) {
			uptodatecount++
			continue
		}
		content = j.ParseContent(i.Content, i.Title, i.Link)
		reply = append(reply, content)
	}
	log.Infof("Succeeded fetching feed. Items: %d. Updated: %s. Up to date count: %d. New feed count: %d", len(feed.Items), feed.Updated, uptodatecount, len(reply))
	return reply, feed.UpdatedParsed, nil
}

func (j *JokesFeed) IsSyncedTime(updatedTime *time.Time, lastUpdatedAt *time.Time) bool {
	if lastUpdatedAt == nil || updatedTime.After(*lastUpdatedAt) {
		return true
	}
	return false
}

func (j *JokesFeed) EmojiInjector(num int) (emojis []string) {
	var index int
	for i := 0; i < num; i++ {
		index = rand.Intn(len(static.EmojiList))
		emojis = append(emojis, static.EmojiList[index])
	}
	return emojis
}

func (j *JokesFeed) GetFeedName() string {
	return j.BotName
}

func (j *JokesFeed) FetchRawFeed() (*gofeed.Feed, error) {

	// set 60 second timeout
	ctx, cancel := context.WithTimeout(context.Background(), j.FetchTimeout)
	defer cancel()

	// parse the reddit jokes feed
	fp := gofeed.NewParser()
	feed, err := fp.ParseURLWithContext(j.Url, ctx)

	if err != nil {
		log.Errorf("Error fetching feed: %t", err.Error())
		return nil, err
	}
	return feed, nil
}
