package main

import (
	"context"
	"fmt"

	"github.com/rdcassin/gator-go/internal/database"
)

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	return func(s *state, cmd command) error {
		username := s.cfg.CurrentUsername
		user, err := s.db.GetUser(context.Background(), username)
		if err != nil {
			return fmt.Errorf("error... please login/register before proceeding")
		}
		return handler(s, cmd, user)
	}

}
