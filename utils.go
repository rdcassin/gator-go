package main

import (
	"fmt"
	"reflect"

	"github.com/rdcassin/gator-go/internal/database"
	"github.com/rdcassin/gator-go/internal/rss"
)

type dataItems interface {
	database.Feed | database.GetFeedsRow | database.CreateFeedFollowRow | database.GetFeedFollowsForUserRow | rss.RSSItem
}

func printResults[T dataItems](items []T, verbose bool) {
    for _, item := range items {
        if verbose {
            printAllFields(item)
        } else {
            printCustom(item)
        }
    }
}

func printAllFields(item interface{}) {
    v := reflect.ValueOf(item)
    t := reflect.TypeOf(item)
    
    // Define field order for consistent output
    fieldOrder := []string{"ID", "CreatedAt", "UpdatedAt", "Url", "Name", "UserID", "FeedID", "FeedName", "UserName", "Title", "Link", "Description", "PubDate"}
    
    // Print fields in order if they exist
    for _, fieldName := range fieldOrder {
        if field, ok := t.FieldByName(fieldName); ok && field.IsExported() {
            value := v.FieldByName(fieldName)
            fmt.Printf("* %-22s %v\n", fieldName+":", value.Interface())
        }
    }
	fmt.Println("===============================================================")
}

func printCustom(item interface{}) {
    switch v := item.(type) {
    case database.Feed:
        fmt.Printf("* %-4s %s\n", "Name:", v.Name)
        fmt.Printf("* %-4s %s\n", "URL:", v.Url)
    case database.GetFeedsRow:
        fmt.Printf("* %-4s %s\n", "Name:", v.Name)
        fmt.Printf("* %-4s %s\n", "URL:", v.Url)
        fmt.Printf("* %-4s %s\n", "User:", v.UserName)
    case database.CreateFeedFollowRow:
        fmt.Printf("* %-4s %s\n", "User:", v.UserName)
        fmt.Printf("* %-4s %s\n", "Feed:", v.FeedName)
    case database.GetFeedFollowsForUserRow:
        fmt.Printf("* %-4s %s\n", "Feed:", v.FeedName)
    case rss.RSSItem:
        fmt.Printf("* %-4s %s\n", "Title:", v.Title)
    default:
        fmt.Printf("Unknown type: %T\n", v)
    }
}