package main

import (
	"errors"
	"log"
	"os"
)

import _ "github.com/lib/pq"

func getCommand() (Command, error) {
	args := os.Args
	if len(args) < 2 {
		return Command{}, errors.New("No command provided")
	}
	command := Command{
		name: args[1],
		args: args[2:],
	}
	return command, nil
}

func printError(e error) {
	log.Fatalf("Error: %v\n", e)
}

func main() {
	state, err := NewState()
	if err != nil {
		printError(err)
	}

	command, err := getCommand()
	if err != nil {
		printError(err)
	}

	registry := GetCommandRegistry()
	err = registry.Run(state, command)
	if err != nil {
		printError(err)
	}
}
