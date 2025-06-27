package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/rdcassin/gator-go/internal/database"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <Username>", cmd.Name)
	}

	username := cmd.Args[0]

	existingUser, err := s.db.GetUser(context.Background(), username)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("user does not exist... please register user first before logging in")
		}
		return fmt.Errorf("internal database error: %s", err)
	}

	err = s.cfg.SetUser(existingUser.Name)
	if err != nil {
		return err
	}

	fmt.Printf("user has been set to %s\n", existingUser.Name)
	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <Username>", cmd.Name)
	}

	username := cmd.Args[0]

	// Checking if user is already registered
	existingUser, err := s.db.GetUser(context.Background(), username)
	if err == nil {
		return fmt.Errorf("%s is already a registered user... please proceed by logging in", existingUser.Name)
	}
	if !errors.Is(err, sql.ErrNoRows) {
		return fmt.Errorf("internal database error: %s", err)
	}

	// Creating new user
	newUserParams := database.CreateUserParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name: username,
	}
	newUser, err := s.db.CreateUser(context.Background(), newUserParams)
	if err != nil {
		return fmt.Errorf("error creating user: %s", err)
	}

	err = s.cfg.SetUser(newUser.Name)
	if err != nil {
		return err
	}

	fmt.Printf("user has been added and set to %s\n", newUser.Name)
	return nil
}

func handlerListUsers(s *state, cmd command) error {
	if len(cmd.Args) != 0 {
		return fmt.Errorf("usage: %s", cmd.Name)
	}

	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("error fetching all users: %s", err)
	}

	currentUser := s.cfg.CurrentUsername
	
	for _, user := range users {
		currentTag := "(current)"
		if user.Name != currentUser {
			currentTag = ""
		}
		fmt.Println(user.Name, currentTag)
	}

	return nil
}