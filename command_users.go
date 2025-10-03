package main

import (
	"context"
	"fmt"
)

func handlerUsers(state *State, cmd Command) error {
	users, err := state.db.GetUsers(context.Background())
	if err != nil {
		return err
	}
	for _, user := range users {
		message := "* " + user.Name
		if state.config.CurrentUserName == user.Name {
			message = message + " (current)"
		}
		fmt.Println(message)
	}
	return nil
}
