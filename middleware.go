package main

import (
	"context"
	"fmt"
	"github.com/qwlp/gatorgo/internal/database"
)

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	return func(s *state, cmd command) error {
		ctx := context.Background()
		user, err := s.db.GetUser(ctx, s.cfg.CurrentUsername)

		if err != nil {
			return fmt.Errorf("Something went wrong when searching for user: %w", err)
		}
		return handler(s, cmd, user)
	}
}
