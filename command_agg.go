package main

import (
	"bootdev-aggregate-go/internal/rss_client"
	"context"
	"fmt"
)

func handlerAgg(s *State, cmd Command) error {
	url := "https://wagslane.dev/index.xml"
	fmt.Println("Fetching feed from: ", url)
	feed, err := rss_client.FetchFeed(context.Background(), url)
	if err != nil {
		return err
	}
	fmt.Println(feed)
	return nil
}
