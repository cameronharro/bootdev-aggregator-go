package main

import (
	"bootdev-aggregate-go/internal/database"
	"context"
	"fmt"
)

func handlerFollowing(s *State, cmd Command, user database.User) error {
	feedFollows, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return err
	}

	for _, feedFollow := range feedFollows {
		fmt.Println(feedFollow.UserName, feedFollow.FeedName)
	}
	return nil
}
