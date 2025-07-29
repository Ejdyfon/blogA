package main

import (
	"context"
	"fmt"
)

func handlerFeeds(s *state, cmd command) error {
	feeds, err := s.db.ListFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("Error while fetching feeds: %w", err)
	}

	for _, v := range feeds {
		usr, _ := s.db.GetUserById(context.Background(), v.UserID)
		fmt.Printf("Feed: %v Url: %v Author: %v\n", v.Name, v.Url, usr.Name)
	}

	return nil
}
