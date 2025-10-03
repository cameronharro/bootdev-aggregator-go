package main

import "context"

func handlerReset(state *State, _ Command) error {
	err := state.db.ResetUsers(context.Background())
	if err != nil {
		return err
	}
	return nil
}

