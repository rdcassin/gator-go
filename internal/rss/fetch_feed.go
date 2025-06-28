package rss

import (
	"context"
	"encoding/xml"
	"fmt"
	"html"
	"net/http"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func FetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	feeds := RSSFeed{}
	
	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
		if err != nil {
			return &feeds, fmt.Errorf("error generating request: %w", err)
		}
	req.Header.Set("User-Agent", "gator-go")
	
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return &feeds, fmt.Errorf("error fetching RSSFeed: %w", err)
	}
	defer res.Body.Close()

	decoder := xml.NewDecoder(res.Body)
	err = decoder.Decode(&feeds)
	if err != nil {
		return &feeds, fmt.Errorf("error decoding RSSFeed: %w", err)
	}

	feeds.Channel.Title = html.UnescapeString(feeds.Channel.Title)
	feeds.Channel.Description = html.UnescapeString(feeds.Channel.Description)

	for i, item := range feeds.Channel.Item {
		item.Title = html.UnescapeString(item.Title)
		item.Description = html.UnescapeString(item.Description)
		feeds.Channel.Item[i] = item
	}

	return &feeds, nil
}