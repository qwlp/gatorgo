package main

import (
	"context"
	"fmt"
)

func handlerFeeds(s *state, cmd command) error {
	if len(cmd.Args) != 0 {
		return fmt.Errorf("usage: %s", cmd.Name)
	}

	ctx := context.Background()
	
	feeds, err := s.db.GetFeeds(ctx)
	if err != nil {
		return fmt.Errorf("couldn't create feed: %w", err)
	}

	fmt.Printf("all feeds: \n")
	for _, f := range feeds {
		user, err := s.db.GetUserWithId(ctx, f.UserID)
		if err != nil {
			return fmt.Errorf("something went wrong when searching for user: %w", err)
		}
		fmt.Printf("- feed name:      %s\n", f.Name)
		fmt.Printf("- feed url:       %s\n", f.Url)
		fmt.Printf("- feed's user:    %s\n", user.Name)
	}
	return nil
}
