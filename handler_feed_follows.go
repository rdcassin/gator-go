package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/rdcassin/gator-go/internal/database"
)

func handlerListFeedFollows(s *state, cmd command) error {
	if len(cmd.Args) > 1 {
		return fmt.Errorf("usage: <%s> [-v]", cmd.Name)
	}

	if len(cmd.Args) == 1 && cmd.Args[0] != "-v" {
		return fmt.Errorf("usage: <%s> [-v]", cmd.Name)
	}

	/* Checker for -v flag... Since I only have the option of one type of flag,
	I'm not storing the flag in a dictionary and verifying that its a key */
	verbose := false
	if len(cmd.Args) == 1 {
		verbose = true
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

	printResults(feeds, verbose)
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
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
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