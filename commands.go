package main

import (
	"fmt"
)

type command struct {
	Name string
	Args []string
}

type commands struct {
	registeredCommands map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	cmdHandler, exists := c.registeredCommands[cmd.Name]
	if !exists {
		return fmt.Errorf("Command does not exists by name: %s", cmd.Name)
	}

	return cmdHandler(s, cmd)
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.registeredCommands[name] = f
}
