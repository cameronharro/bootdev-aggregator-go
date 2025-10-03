package main

import (
	"bootdev-aggregate-go/internal/database"
	"context"
	"errors"
	"fmt"
	"log/slog"
	"time"

	"github.com/google/uuid"
)

func handlerRegister(state *State, cmd Command) error {
	if len(cmd.args) == 0 {
		return errors.New("No login username provided")
	}
	result, err := state.db.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.args[0],
	})
	if err != nil {
		return err
	}
	slog.Info(fmt.Sprintf("Registered %s and set as active user", result.Name))
	state.config.SetUser(result.Name)
	return nil
}
