package main

import (
	"bootdev-aggregate-go/internal/config"
	"bootdev-aggregate-go/internal/database"
	"database/sql"
)

type State struct {
	db     *database.Queries
	config *config.Config
}

func NewState() (*State, error) {
	config, err := config.Read()
	if err != nil {
		return &State{}, err
	}

	dbConnection, err := sql.Open("postgres", config.DbUrl)
	if err != nil {
		return &State{}, err
	}

	dbQueries := database.New(dbConnection)
	return &State{
		db:     dbQueries,
		config: &config,
	}, nil
}
