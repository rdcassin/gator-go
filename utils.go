package main

import (
	"fmt"
	"reflect"
	"time"

	"github.com/rdcassin/gator-go/internal/database"
)

type dataItems interface {
	database.Feed | database.GetFeedsRow | database.CreateFeedFollowRow | database.GetFeedFollowsForUserRow | database.GetPostsForUserRow
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
	fieldOrder := []string{"ID", "CreatedAt", "UpdatedAt", "Title", "Url", "Description", "PublishedAt", "Name", "UserID", "FeedID", "FeedName", "UserName"}

	// Print fields in order if they exist
	for _, fieldName := range fieldOrder {
		if field, ok := t.FieldByName(fieldName); ok && field.IsExported() {
			value := v.FieldByName(fieldName)

            switch fieldName {
            case "Description":
                if value.IsValid() && value.Kind() == reflect.Struct {
                    strField := value.FieldByName("String")
                    fmt.Printf("* %-22s %v\n", fieldName+":", strField.String())
                }
            case "PublishedAt":
				if value.IsValid() && value.Kind() == reflect.Struct {
					timeField := value.FieldByName("Time")
					if timeField.IsValid() {
						t, ok := timeField.Interface().(time.Time)
						if ok {
							fmt.Printf("* %-22s %v\n", fieldName+":", t)
						} else {
							fmt.Printf("* %-22s %v\n", fieldName+":", timeField.Interface())
						}
					}
				}
			default:
				fmt.Printf("* %-22s %v\n", fieldName+":", value.Interface())
			}
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
	default:
		fmt.Printf("Unknown type: %T\n", v)
	}
}
