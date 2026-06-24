package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/qwlp/gatorgo/internal/database"
)

func handlerFollow(s *state, cmd command, user database.User) error {
	ctx := context.Background()

	url := cmd.Args[0]

	feed, err := s.db.GetFeedByUrl(ctx, url)
	if err != nil {
		return fmt.Errorf("something went wrong when fetching feed by url: %w", err)
	}

	feedFollow, err := s.db.CreateFeedFollow(ctx, database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	})

	if err != nil {
		return fmt.Errorf("something went wrong when creating new feed follow: %w", err)
	}

	fmt.Printf("User %s has follow feed %s!", feedFollow.UserName, feedFollow.FeedName)

	return nil
}

func handlerFollowing(s *state, cmd command, user database.User) error {
	ctx := context.Background()
	feedFollows, err := s.db.GetFeedsFollowForUser(ctx, user.ID)
	if err != nil {
		return fmt.Errorf("something went wrong when getting feeds follow for user %s: %w", user.Name, err)
	}

	fmt.Printf("User %s has followed the following feeds:\n", user.Name)
	for _, feedFollow := range feedFollows {
		feed, err := s.db.GetFeedById(ctx, feedFollow.FeedID)
		if err != nil {
			return fmt.Errorf("something went wrong when getting feed: %w", err)
		}
		fmt.Printf("- %s", feed.Name)
	}

	return nil
}

func handlerUnfollow(s *state, cmd command, user database.User) error {
	ctx := context.Background()

	url := cmd.Args[0]
	feed, err := s.db.GetFeedByUrl(ctx, url)

	_, err = s.db.DeleteFeedsFollowWithCombination(ctx, database.DeleteFeedsFollowWithCombinationParams{UserID: user.ID, FeedID: feed.ID})
	if err != nil {
		return fmt.Errorf("something went wrong when deleting feed: %w", err)
	}

	fmt.Printf("successfully deleted feed %+v", feed)
	return nil

}

