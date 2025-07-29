package main

import (
	"context"
	"fmt"
)

func handlerUsers(s *state, cmd command) error {
	usr, err := s.db.GetAllUsers(context.Background())
	if err != nil {
		return fmt.Errorf("Error while fetching users: %w", err)
	}

	for _, v := range usr {
		if v.Name != s.cfg.CurrentUserName {
			fmt.Println("* " + v.Name)
		} else {
			fmt.Println("* " + v.Name + " (current)")
		}
	}

	return nil
}
