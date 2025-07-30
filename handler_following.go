package main

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

func handlerFollowing(s *state, cmd command) error {
	usr, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if usr.ID == uuid.Nil {
		return fmt.Errorf("User doesnt exist: %w", s.cfg.CurrentUserName)
	}

	feedfollows, err := s.db.GetFeedFollowsForUser(context.Background(), usr.Name)
	if err != nil {
		return fmt.Errorf("couldn't fetch follow feed: %w", err)
	}

	for _, v := range feedfollows {
		fmt.Println(v.FeedName + "\n")
	}

	return nil
}
