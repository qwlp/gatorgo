package main

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"time"

	"github.com/qwlp/gatorgo/internal/database"
)

func handlerRegister(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}

	name := cmd.Args[0]
	ctx := context.Background()
	user, err := s.db.CreateUser(ctx, database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      name,
	},
	)

	if err != nil {
		return fmt.Errorf("couldn't create user: %w", err)
	}

	err = s.cfg.SetUser(user.Name)
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}

	fmt.Printf("User %s was created", user.Name)
	return nil
}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}

	username := cmd.Args[0]
	ctx := context.Background()

	_, err := s.db.GetUser(ctx, username)
	if err != nil {
		return fmt.Errorf("something went wrong when getting user: %w", err)
	}

	err = s.cfg.SetUser(username)
	if err != nil {
		return fmt.Errorf("could't set username: %w", err)
	}

	fmt.Printf("You have logged in as %s", username)

	return nil
}

func handlerUsers(s *state, cmd command) error {
	ctx := context.Background()
	users, err := s.db.GetUsers(ctx)

	if err != nil {
		fmt.Errorf("something went wrong when getting users: %w", err)
	}

	for _, u := range users {
		if s.cfg.CurrentUsername == u.Name {
			fmt.Printf("* %s (current)", u.Name)
		} else {
			fmt.Printf("* %s", u.Name)
		}
	}
	return nil
}

func printUser(user database.User) {
	fmt.Printf("  * ID:		 	%v\n", user.ID)
	fmt.Printf("  * Name:		%v\n", user.Name)
}
