package main

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"time"

	"github.com/qwlp/gatorgo/internal/database"
)

func handlerAgg(s *state, cmd command) error {
	timeBetweenRequestsString := cmd.Args[0]
	timeBetweenRequests, err := time.ParseDuration(timeBetweenRequestsString)

	if err != nil {
		return fmt.Errorf("something went wrong when parsing time into string: %w", err)
	}

	fmt.Printf("collecting feed every %s\n", timeBetweenRequests)

	ticker := time.NewTicker(timeBetweenRequests)
	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}

	return nil
}

func scrapeFeeds(s *state) error {
	ctx := context.Background()
	nextFeed, err := s.db.GetNextFeedToFetch(ctx)
	if err != nil {
		return fmt.Errorf("something went wrong when fetching next feed: %v", err)
	}

	_, err = s.db.MarkFeedFetched(ctx, nextFeed.ID)
	if err != nil {
		return fmt.Errorf("something went wrong when marking feed as fetched: %w", err)
	}

	feed, err := fetchFeed(ctx, nextFeed.Url)

	if err != nil {
		return fmt.Errorf("something went wrong while fetching feed via url: %w", err)
	}
	for _, i := range feed.Channel.Item {

		publishedAt, err := time.Parse(time.RFC1123, i.PubDate)
		if err != nil {
			return fmt.Errorf("something went wrong when parsing time: %w", err)
		}

		post := database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now().UTC(),
			UpdatedAt:   time.Now().UTC(),
			Title:       i.Title,
			Url:         i.Link,
			Description: i.Description,
			PublishedAt: publishedAt,
			FeedID:      nextFeed.ID,
		}

		fmt.Println(post)
		_, err = s.db.CreatePost(ctx, post)
		if err != nil {
			fmt.Printf("Something went wrong when creating post: %v", err)
		} else {
			fmt.Printf("Created post %+v", post)
		}
	}

	return nil
}
