package main

import (
	"context"
	"fmt"
	"time"

	"github.com/Ejdyfon/genA/internal/database"
	"github.com/google/uuid"
)

func handlerRegister(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}
	name := cmd.Args[0]

	usr, err := s.db.GetUser(context.Background(), name)
	if usr.ID != uuid.Nil {
		return fmt.Errorf("User already exist: %w", name)
	}

	param := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      name,
	}

	newUsr, err := s.db.CreateUser(context.Background(), param)
	if err != nil {
		return fmt.Errorf("couldn't create user: %w", err)
	}
	s.cfg.SetUser(newUsr.Name)
	fmt.Println("User was created successfully!")
	return nil
}
