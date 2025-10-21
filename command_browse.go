package main

import (
	"bootdev-aggregate-go/internal/database"
	"context"
	"fmt"
	"strconv"
)

func handlerBrowse(s *State, cmd Command, user database.User) error {
	posts, err := s.db.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		ID:    user.ID,
		Limit: int32(getLimit(cmd.args)),
	})
	if err != nil {
		return err
	}

	for _, post := range posts {
		fmt.Println(post.Title, post.Url)
	}
	return nil
}

func getLimit(args []string) int {
	initial := 2
	if len(args) == 0 {
		return initial
	}
	result, err := strconv.Atoi(args[0])
	if err != nil {
		return initial
	}
	return result
}
