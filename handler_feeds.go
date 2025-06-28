package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/rdcassin/gator-go/internal/database"
)

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

	feedParams := database.CreateFeedParams {
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name: name,
		Url: feedURL,
		UserID: userID,
	}

	feed, err := s.db.CreateFeed(context.Background(), feedParams)
	if err != nil {
		return fmt.Errorf("error creating new feed: %w", err)
	}

	printFeed(feed)
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