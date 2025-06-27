package main

import (
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: login <Username>")
	}

	username := cmd.Args[0]

	err := s.cfg.SetUser(username)
	if err != nil {
		return err
	}

	fmt.Printf("user has been set to %s\n", username)
	return nil
}