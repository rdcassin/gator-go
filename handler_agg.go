package main

import (
	"fmt"
	"time"
)

/* const (
	feedURL1 string = "https://www.wagslane.dev/index.xml"
	feedURL2 string = "https://techcrunch.com/feed/"
	feedURL3 string = "https://news.ycombinator.com/rss"
	feedURL4 string = "https://blog.boot.dev/index.xml"
)
*/

func handlerAggregate(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: <%s> <TimeBetweenRequests> (e.g. 1h30m5s; must be greater than 15s)", cmd.Name)
	}

	rawTime := cmd.Args[0]
	parsedTime, err := time.ParseDuration(rawTime)
	if err != nil {
		fmt.Println("error parsing time... using default time of 1m")
		parsedTime = time.Minute
	}

	if parsedTime < 15 * time.Second {
		fmt.Println("time between request cannot be less than 15 seconds... using default time of 1m")
		parsedTime = time.Minute
	}

	fmt.Printf("Collecting feeds every %v\n", parsedTime)
	fmt.Println("Press Ctrl+C to stop the program")

	scrapeFeeds(s)
	go func() {
		ticker := time.NewTicker(parsedTime)
		for ; ; <-ticker.C {
			scrapeFeeds(s)
		}
	}()

	return nil
}
