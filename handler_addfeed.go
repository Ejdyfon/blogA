package main

import (
	"context"
	"fmt"
	"time"

	"github.com/Ejdyfon/genA/internal/database"
	"github.com/google/uuid"
)

func handlerAddFeed(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 2 {
		return fmt.Errorf("usage: %s <name> <url>", cmd.Name)
	}
	name := cmd.Args[0]
	url := cmd.Args[1]

	param := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      name,
		Url:       url,
		UserID:    user.ID,
	}

	feed, err := s.db.CreateFeed(context.Background(), param)
	if err != nil {
		return fmt.Errorf("couldn't set current feed: %w", err)
	}

	fmt.Println("Feed added successfully!")

	param1 := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	}

	feedfollow, err := s.db.CreateFeedFollow(context.Background(), param1)
	if err != nil {
		return fmt.Errorf("couldn't follow feed: %w", err)
	}

	fmt.Printf("Feed %v successfully followed by user %v !", feedfollow.FeedName, feedfollow.UserName)
	return nil
}
