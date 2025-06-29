package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/rdcassin/gator-go/internal/database"
)

func handlerBrowse(s *state, cmd command, user database.User) error {
	if len(cmd.Args) > 1 {
		return fmt.Errorf("usage: <%s> [PostLimit] (PostLimit must be a positive integer and defaults to 2 if no value is entered)", cmd.Name)
	}

	limitStr := cmd.Args[0]
	limit64, err := strconv.Atoi(limitStr)
	limit := int32(limit64)
	if err != nil {
		fmt.Println("error converting [PostLimit] to integer... using default value of 2")
		limit = 2
	}

	if limit <= 0 {
		fmt.Println("[PostLimit] must be greater than 0... using default value of 2")
		limit = 2
	}

	userID := user.ID
	limit = int32(limit)

	getPostParams := database.GetPostsForUserParams{
		UserID: userID,
		Limit: limit,
	}

	posts, err := s.db.GetPostsForUser(context.Background(), getPostParams)
	if err != nil {
		return fmt.Errorf("error fetching posts for user: %w", err)
	}

	printResults(posts, true)
	return nil
}