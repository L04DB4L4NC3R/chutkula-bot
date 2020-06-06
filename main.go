package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"net/http"

	strip "github.com/grokify/html-strip-tags-go"
	"github.com/joho/godotenv"
	"github.com/mmcdole/gofeed"
	tbot "github.com/yanzay/tbot/v2"
)

func parseContent(content string, title string) string {
	replacer := strings.NewReplacer("&quot", "", "&#32", "", ";", "", "[link]", "", "[comments", "", "submitted by", "\n\nsubmitted by: ", "]", "", "&#39", "")
	content = replacer.Replace(strip.StripTags(content))
	return fmt.Sprintf("👽🤣😋😏 \n\n%s\n\n%s\n\n😂😳😛😵", title, content)
}

func fetchPosts() []string {
	// set 60 second timeout
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	// parse the reddit jokes feed
	fp := gofeed.NewParser()
	feed, _ := fp.ParseURLWithContext("https://www.reddit.com/r/Jokes/.rss", ctx)

	// open a file in read write mode
	f, _ := os.OpenFile("./last_updated.txt", os.O_RDWR, 0644)
	defer f.Close()

	// scan the last known timestamp
	var line string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line = scanner.Text()
		// fmt.Println(line)
	}

	layout := "2020-06-06 15:24:15 +0000 UTC"
	lastKnownTime, _ := time.Parse(line, layout)

	// check if we are caught up or not
	// [1] write latest time if we aren't
	// [2] return nothing if we are
	if feed.UpdatedParsed.After(lastKnownTime) {
		fmt.Println("Grabbed Feed")
		f.WriteString(feed.UpdatedParsed.String())
		var reply []string
		var content string
		for _, i := range feed.Items {
			// grab items in p tag
			content = parseContent(i.Content, i.Title)
			reply = append(reply, content)
		}
		fmt.Println("Giving reply")
		return reply
	}
	fmt.Println("You are all caught up")
	return nil
}

func sendJokes(c *tbot.Client, chatID string) {
	jokes := fetchPosts()
	if jokes == nil {
		c.SendMessage("-495493407", "You are all caught up")
	} else {
		for _, v := range jokes {
			c.SendMessage(chatID, v)
		}
	}
}
func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	bot := tbot.New(os.Getenv("BOT_TOKEN"))
	c := bot.Client()
	bot.HandleMessage(".*joke.*", func(m *tbot.Message) {
		c.SendChatAction(m.Chat.ID, tbot.ActionTyping)
		c.SendMessage(m.Chat.ID, "Hello! I am the HumourBaba")
		sendJokes(c, m.Chat.ID)
	})

	// sendJokes(c, os.Getenv("GROUP_ID")
	log.Println("Starting bot....")
	go log.Fatal(bot.Start())
	http.ListenAndServe(os.Getenv("PORT"), nil)
}
