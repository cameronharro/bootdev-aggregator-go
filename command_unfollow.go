package main

import (
	"bootdev-aggregate-go/internal/database"
	"context"
	"fmt"
)

func handlerUnFollow(s *State, cmd Command, user database.User) error {
	if len(cmd.args) < 1 {
		return fmt.Errorf("usage: unfollow <url>")
	}

	feed, err := s.db.GetFeed(context.Background(), cmd.args[0])
	if err != nil {
		return err
	}

	err = s.db.DeleteFeedFollowForUserAndFeed(context.Background(), database.DeleteFeedFollowForUserAndFeedParams{
		UserID: user.ID,
		FeedID: feed.ID,
	})
	if err != nil {
		return err
	}

	return nil
}
