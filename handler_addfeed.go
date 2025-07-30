package main

import (
	"context"
	"fmt"
	"time"

	"github.com/Ejdyfon/genA/internal/database"
	"github.com/google/uuid"
)

func handlerAddfeed(s *state, cmd command) error {
	if len(cmd.Args) != 2 {
		return fmt.Errorf("usage: %s <name> <url>", cmd.Name)
	}
	name := cmd.Args[0]
	url := cmd.Args[1]

	usr, _ := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if usr.ID == uuid.Nil {
		return fmt.Errorf("User doesnt exist: %w", s.cfg.CurrentUserName)
	}

	param := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      name,
		Url:       url,
		UserID:    usr.ID,
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
		UserID:    usr.ID,
		FeedID:    feed.ID,
	}

	feedfollow, err := s.db.CreateFeedFollow(context.Background(), param1)
	if err != nil {
		return fmt.Errorf("couldn't follow feed: %w", err)
	}

	fmt.Printf("Feed %v successfully followed by user %v !", feedfollow.FeedName, feedfollow.UserName)
	return nil
}
