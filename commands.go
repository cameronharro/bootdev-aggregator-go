package main

import "fmt"

type Command struct {
	name string
	args []string
}

type Commands struct {
	registry map[string]func(*State, Command) error
}

func (c *Commands) Run(s *State, cmd Command) error {
	callback, ok := c.registry[cmd.name]
	if !ok {
		return fmt.Errorf("%s is not a registered command", cmd.name)
	}
	return callback(s, cmd)
}

func (c *Commands) register(name string, f func(*State, Command) error) {
	c.registry[name] = f
}

func GetCommandRegistry() Commands {
	c := Commands{
		registry: make(map[string]func(*State, Command) error),
	}
	c.register("login", handlerLogin)
	c.register("register", handlerRegister)
	c.register("reset", handlerReset)
	c.register("users", handlerUsers)
	c.register("agg", handlerAgg)
	c.register("addfeed", middlewareLoggedIn(handlerAddFeed))
	c.register("feeds", handlerFeeds)
	c.register("follow", middlewareLoggedIn(handlerFollow))
	c.register("following", middlewareLoggedIn(handlerFollowing))
	c.register("unfollow", middlewareLoggedIn(handlerUnFollow))
	c.register("browse", middlewareLoggedIn(handlerBrowse))
	return c
}
