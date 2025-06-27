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

func (c *commands) register(name string, f func(*state, command) error) {
	c.registeredCommands[name] = f
}

func (c *commands) run(s *state, cmd command) error {
	if s.cfg == nil {
		return fmt.Errorf("configuration invalid... can't run command")
	}

	chosenCommand, exists := c.registeredCommands[cmd.Name]
	if !exists {
		return fmt.Errorf("invalid command")
	}

	return chosenCommand(s, cmd)
}
