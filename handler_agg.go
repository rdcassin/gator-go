package main

import (
	"context"
	"fmt"

	"github.com/rdcassin/gator-go/internal/rss"
)

const feedURL = "https://www.wagslane.dev/index.xml"

func handlerAggregate(s *state, cmd command) error {
	if len(cmd.Args) != 0 {
		return fmt.Errorf("usage: %s", cmd.Name)
	}

	feed, err := rss.FetchFeed(context.Background(), feedURL)
	if err != nil {
		return err
	}

	fmt.Printf("Feed: %+v\n", *feed)

	return nil
}
