package main

import (
	"bootdev-aggregate-go/internal/database"
	"bootdev-aggregate-go/internal/rss_client"
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"strings"
	"time"

	"github.com/google/uuid"
)

func handlerAgg(s *State, cmd Command) error {
	ticker, err := getTicker(cmd.args)
	if err != nil {
		return err
	}

	scrapeFeeds(s, ticker)
	return nil
}

func getTicker(args []string) (*time.Ticker, error) {
	var timeBetweenReqs time.Duration
	if len(args) > 0 {
		duration, err := time.ParseDuration(args[0])
		if err != nil {
			return nil, err
		}
		timeBetweenReqs = duration
	} else {
		timeBetweenReqs = time.Minute
	}
	fmt.Println("Fetching feeds every ", timeBetweenReqs)
	return time.NewTicker(timeBetweenReqs), nil
}

func scrapeFeeds(s *State, ticker *time.Ticker) {
	for ; ; <-ticker.C {
		dbFeed, err := pickFeed(s)
		if err != nil {
			continue
		}
		fetchedFeed, err := fetchFeed(dbFeed.Url)
		if err != nil {
			continue
		}
		saveItems(s, dbFeed.ID, fetchedFeed)
	}
}

func pickFeed(s *State) (database.Feed, error) {
	feed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return database.Feed{}, err
	}
	err = s.db.MarkFeedFetched(context.Background(), database.MarkFeedFetchedParams{
		ID: feed.ID,
	})
	if err != nil {
		return database.Feed{}, err
	}
	return feed, nil
}

func fetchFeed(url string) (*rss_client.RSSFeed, error) {
	fmt.Println("Fetching feed from: ", url)
	return rss_client.FetchFeed(context.Background(), url)
}

func saveItems(s *State, feedId uuid.UUID, feed *rss_client.RSSFeed) {
	for _, item := range feed.Channel.Items {
		publishedAt, err := time.Parse("Mon, 02 Jan 2006 15:04:05 -0700", item.PubDate)
		if err != nil {
			continue
		}
		_, err = s.db.CreatePost(context.Background(), database.CreatePostParams{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Title:     item.Title,
			Url:       item.Link,
			Description: sql.NullString{
				String: item.Description,
				Valid:  true,
			},
			PublishedAt: sql.NullTime{
				Time:  publishedAt,
				Valid: true,
			},
			FeedID: feedId,
		})
		if err != nil {
			if !strings.Contains(err.Error(), "posts_url_key") {
				slog.Error(err.Error())
			}
		}
	}
}
