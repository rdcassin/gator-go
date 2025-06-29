package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"github.com/rdcassin/gator-go/internal/database"
	"github.com/rdcassin/gator-go/internal/rss"
)

const (
	uniqueViolation string = "23505" // PostgreSQL driver pq error code for UNIQUE violation
)

var acceptedErrorCodes = map[string]struct{}{
	uniqueViolation: {},
}

func scrapeFeeds(s *state) error {
	feed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return fmt.Errorf("error fetching next feed to update: %w", err)
	}

	feedID := feed.ID

	markFeed := database.MarkFeedFetchedParams{
		LastFetchedAt: sql.NullTime{
			Time:  time.Now().UTC(),
			Valid: true,
		},
		ID: feedID,
	}

	err = s.db.MarkFeedFetched(context.Background(), markFeed)
	if err != nil {
		return fmt.Errorf("error marking feed as fetched: %w", err)
	}

	feedURL := feed.Url

	rssFeed, err := rss.FetchFeed(context.Background(), feedURL)
	if err != nil {
		return fmt.Errorf("error fetching RSS Feed: %w", err)
	}

	items := rssFeed.Channel.Item
	for _, item := range items {
		pubDate := parsePubDate(item.PubDate)

		newPost := database.CreatePostParams{
			ID:        uuid.New(),
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
			Title:     item.Title,
			Url:       item.Link,
			Description: sql.NullString{
				String: item.Description,
				Valid:  true,
			},
			PublishedAt: pubDate,
			FeedID:      feedID,
		}
		_, err = s.db.CreatePost(context.Background(), newPost)

		if err != nil {
			if pqErr, ok := err.(*pq.Error); ok {
				if _, exists := acceptedErrorCodes[string(pqErr.Code)]; exists {
					continue
				}
			}
			return fmt.Errorf("error creating post: %w", err)
		}
	}
	return nil
}

func parsePubDate(pubdateStr string) sql.NullTime {
	pubDate, err := time.Parse(time.RFC1123Z, pubdateStr)
	isValidPubDate := true
	if err != nil {
		fmt.Println("error parsing publication date... storing with time of 0")
		isValidPubDate = false
	}
	postPubDate := sql.NullTime{
		Time:  pubDate,
		Valid: isValidPubDate,
	}
	return postPubDate
}
