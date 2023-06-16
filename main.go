package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/landrunner/go-rss-line/line"
	"github.com/mmcdole/gofeed"
)

func main() {
	token := os.Getenv("LINE_TOKEN")
	if token == "" {
		fmt.Println("No environment variable LINE_TOKEN")
		return
	}
	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL("http://feeds.bbci.co.uk/news/rss.xml")
	titles := []string{}
	for i := 0; i < 10; i++ {
		titles = append(titles, feed.Items[i].Title)
	}
	l := line.New(token)
	l.SendMessage(strings.Join(titles, "\n\n"))
}
