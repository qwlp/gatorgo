package main

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"time"

	"github.com/qwlp/gatorgo/internal/database"
)

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.Args) != 2 {
		return fmt.Errorf("usage: %s <name> <url>", cmd.Name)
	}

	name := cmd.Args[0]
	url := cmd.Args[1]

	ctx := context.Background()

	user, err := s.db.GetUser(ctx, s.cfg.CurrentUsername)
	if err != nil {
		return fmt.Errorf("something went wrong when searching for user: %w", err)
	}
	
		
	feed, err := s.db.CreateFeed(ctx, database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      name,
		Url:       url,
		UserID:    user.ID,
	},
	)

	if err != nil {
		return fmt.Errorf("couldn't create feed: %w", err)
	}
	fmt.Printf("created new feed: %+v", feed)
	return nil
}
