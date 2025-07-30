package main

import (
	"context"
	"fmt"
	"time"

	"github.com/Ejdyfon/genA/internal/database"
	"github.com/google/uuid"
)

func handlerFollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <url>", cmd.Name)
	}
	url := cmd.Args[0]

	feed, _ := s.db.GetFeedByUrl(context.Background(), url)
	if feed.ID == uuid.Nil {
		return fmt.Errorf("Feed doesnt exist: %w", url)
	}

	param := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	}

	feedfollow, err := s.db.CreateFeedFollow(context.Background(), param)
	if err != nil {
		return fmt.Errorf("couldn't follow feed: %w", err)
	}

	fmt.Printf("Feed %v successfully followed by user %v !", feedfollow.FeedName, feedfollow.UserName)
	return nil
}
