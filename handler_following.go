package main

import (
	"context"
	"fmt"

	"github.com/Ejdyfon/genA/internal/database"
)

func handlerFollowing(s *state, cmd command, user database.User) error {
	feedfollows, err := s.db.GetFeedFollowsForUser(context.Background(), user.Name)
	if err != nil {
		return fmt.Errorf("couldn't fetch follow feed: %w", err)
	}

	for _, v := range feedfollows {
		fmt.Println(v.FeedName + "\n")
	}

	return nil
}
