package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/rdcassin/gator-go/internal/database"
)

func handlerListFeedFollows(s *state, cmd command) error {
	if len(cmd.Args) != 0 {
		return fmt.Errorf("usage: %s", cmd.Name)
	}

	username := s.cfg.CurrentUsername
	user, err := s.db.GetUser(context.Background(), username)
	if err != nil {
		return fmt.Errorf("error fetching feed: %w", err)
	}
	userID := user.ID

	feeds, err := s.db.GetFeedFollowsForUser(context.Background(), userID)
	if err != nil {
		return fmt.Errorf("error fetching feeds for %s", err)
	}

	for _, feed := range feeds {
		printFeedFollows(feed)
		fmt.Println("===============================================================")
	}

	return nil
}

func handlerFollow(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <FeedURL>", cmd.Name)
	}

	feedURL := cmd.Args[0]
	feed, err := s.db.GetFeedByURL(context.Background(), feedURL)
	if err != nil {
		return fmt.Errorf("error fetching feed: %w", err)
	}
	feedID := feed.ID
	username := s.cfg.CurrentUsername
	user, err := s.db.GetUser(context.Background(), username)
	if err != nil {
		return fmt.Errorf("error fetching user info: %w", err)
	}
	userID := user.ID

	newFeedFollow := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    userID,
		FeedID:    feedID,
	}

	_, err = s.db.CreateFeedFollow(context.Background(), newFeedFollow)
	if err != nil {
		return fmt.Errorf("error following feed: %w", err)
	}

	fmt.Printf("%s is now following %s\n", username, feedURL)

	return nil
}

func handlerListFeeds(s *state, cmd command) error {
	if len(cmd.Args) != 0 {
		return fmt.Errorf("usage: %s", cmd.Name)
	}

	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("error fetching all feeds: %w", err)
	}

	for _, feed := range feeds {
		user, err := s.db.GetUserById(context.Background(), feed.UserID)
		if err != nil {
			fmt.Errorf("error fetching user info: %w", err)
		}
		username := user.Name

		printFeed(feed)
		fmt.Printf("* %-22s %s\n", "Username", username)
		fmt.Println("===============================================================")
	}

	return nil
}

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.Args) != 2 {
		return fmt.Errorf("usage: %s <FeedName> <FeedURL>", cmd.Name)
	}

	name := cmd.Args[0]
	feedURL := cmd.Args[1]

	username := s.cfg.CurrentUsername
	user, err := s.db.GetUser(context.Background(), username)
	if err != nil {
		return fmt.Errorf("error fetching user info: %w", err)
	}

	userID := user.ID

	feedParams := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
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
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    userID,
		FeedID:    feed.ID,
	}

	_, err = s.db.CreateFeedFollow(context.Background(), newFeedFollow)
	if err != nil {
		return fmt.Errorf("error following feed: %w", err)
	}

	printFeed(feed)
	fmt.Printf("%s feed was added to the database and %s is now following this feed", feed.Name, username)
	return nil
}

func printFeed(f database.Feed) {
	fmt.Printf("* %-22s %s\n", "ID:", f.ID)
	fmt.Printf("* %-22s %s\n", "Created At:", f.CreatedAt)
	fmt.Printf("* %-22s %s\n", "Updated At:", f.UpdatedAt)
	fmt.Printf("* %-22s %s\n", "Name:", f.Name)
	fmt.Printf("* %-22s %s\n", "URL:", f.Url)
	fmt.Printf("* %-22s %s\n", "User ID:", f.UserID)
}

func printFeedFollows(f database.GetFeedFollowsForUserRow) {
	fmt.Printf("* %-22s %s\n", "ID:", f.ID)
	fmt.Printf("* %-22s %s\n", "Created At:", f.CreatedAt)
	fmt.Printf("* %-22s %s\n", "Updated At:", f.UpdatedAt)
	fmt.Printf("* %-22s %s\n", "Username:", f.UserName)
	fmt.Printf("* %-22s %s\n", "User ID:", f.UserID)
	fmt.Printf("* %-22s %s\n", "Feed Name:", f.FeedName)
	fmt.Printf("* %-22s %s\n", "Feed ID:", f.FeedID)
}