package main

import (
	"context"
	"fmt"
	"strconv"
)

func handlerBrowse(s *state, cmd command) error {
	var limitString string
	if len(cmd.Args) != 1 {
		limitString = "2"
	} else {
		limitString = cmd.Args[0]
	}

	ctx := context.Background()

	limit, err := strconv.Atoi(limitString)
	if err != nil {
		return fmt.Errorf("you have not typed a correct int limit value: %w", err)
	}

	posts, err := s.db.GetPostsForUser(ctx, int32(limit))
	if err != nil {
		return fmt.Errorf("something went wrong when getting posts for user: %w", err)
	}

	for _, post := range posts {
		fmt.Printf("- %+v", post)
	}
	return nil
}
