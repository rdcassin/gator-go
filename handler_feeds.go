package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/rdcassin/gator-go/internal/database"
	"github.com/rdcassin/gator-go/internal/rss"
)

func handlerListFeeds(s *state, cmd command) error {
	if len(cmd.Args) != 0 {
		return fmt.Errorf("usage: <%s>", cmd.Name)
	}

	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("error fetching all feeds: %w", err)
	}

	printResults(feeds, true)
	return nil
}

func handlerAddFeed(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 2 {
		return fmt.Errorf("usage: <%s> <FeedName> <FeedURL>", cmd.Name)
	}

	name := cmd.Args[0]
	feedURL := cmd.Args[1]

	username := user.Name
	userID := user.ID

	feedParams := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      name,
		Url:       feedURL,
		UserID:    userID,
	}

	feed, err := s.db.CreateFeed(context.Background(), feedParams)
	if err != nil {
		return fmt.Errorf("error creating new feed: %w", err)
	}

	newFeedFollow := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    userID,
		FeedID:    feed.ID,
	}

	_, err = s.db.CreateFeedFollow(context.Background(), newFeedFollow)
	if err != nil {
		return fmt.Errorf("error following feed: %w", err)
	}

	printResults([]database.Feed{feed}, true)
	fmt.Printf("%s feed was added to the database and %s is now following this feed", feed.Name, username)
	return nil
}

func scrapeFeeds(s* state) error {
	feed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return fmt.Errorf("error fetching next feed to update")
	}

	feedID := feed.ID

	markFeed := database.MarkFeedFetchedParams {
		LastFetchedAt: sql.NullTime{
			Time: time.Now().UTC(),
			Valid: true,
		},
		ID: feedID,
	}

	err = s.db.MarkFeedFetched(context.Background(), markFeed)

	feedURL := feed.Url

	rssFeed, err := rss.FetchFeed(context.Background(), feedURL)
	if err != nil {
		return fmt.Errorf("error fetching RSS Feed: %w", err)
	}

	items := rssFeed.Channel.Item
	printResults(items, false)
	return nil
}