package main

import (
	"context"
	"fmt"
)

func handlerLogin(s *State, cmd Command) error {
	if len(cmd.args) < 1 {
		return fmt.Errorf("No login username provided")
	}
	dbUser, err := s.db.GetUser(context.Background(), cmd.args[0])
	if err != nil {
		return err
	}

	err = s.config.SetUser(dbUser.Name)
	if err != nil {
		return fmt.Errorf("Failed to update current user")
	}
	fmt.Printf("Updated current user to %s\n", cmd.args[0])
	return nil
}
