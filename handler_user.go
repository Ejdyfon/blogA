package main

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %v <name>", cmd.Name)
	}
	name := cmd.Args[0]

	usr, _ := s.db.GetUser(context.Background(), name)
	if usr.ID == uuid.Nil {
		return fmt.Errorf("User doesnt exist: %w", name)
	}

	err := s.cfg.SetUser(name)
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}

	fmt.Println("User switched successfully!")
	return nil
}
