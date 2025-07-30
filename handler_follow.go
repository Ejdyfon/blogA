package main

import (
	"context"
	"fmt"
	"time"

	"github.com/Ejdyfon/genA/internal/database"
	"github.com/google/uuid"
)

func handlerFollow(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <url>", cmd.Name)
	}
	url := cmd.Args[0]

	feed, _ := s.db.GetFeedByUrl(context.Background(), url)
	if feed.ID == uuid.Nil {
		return fmt.Errorf("Feed doesnt exist: %w", url)
	}

	usr, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if usr.ID == uuid.Nil {
		return fmt.Errorf("User doesnt exist: %w", s.cfg.CurrentUserName)
	}

	param := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    usr.ID,
		FeedID:    feed.ID,
	}

	feedfollow, err := s.db.CreateFeedFollow(context.Background(), param)
	if err != nil {
		return fmt.Errorf("couldn't follow feed: %w", err)
	}

	fmt.Printf("Feed %v successfully followed by user %v !", feedfollow.FeedName, feedfollow.UserName)
	return nil
}
