package main

import (
	"context"
	"fmt"

	"github.com/Ejdyfon/genA/internal/database"
	"github.com/google/uuid"
)

func handlerUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <url>", cmd.Name)
	}
	url := cmd.Args[0]

	feed, _ := s.db.GetFeedByUrl(context.Background(), url)
	if feed.ID == uuid.Nil {
		return fmt.Errorf("Feed doesnt exist: %w", url)
	}

	param := database.DeleteFeedFollowByUserAndFeedParams{
		UserID: user.ID,
		FeedID: feed.ID,
	}

	err := s.db.DeleteFeedFollowByUserAndFeed(context.Background(), param)
	if err != nil {
		return fmt.Errorf("couldn't unfollow: %w", err)
	}

	fmt.Printf("Feed %v successfully unfollowed by user %v !", feed.Name, user.Name)
	return nil
}
