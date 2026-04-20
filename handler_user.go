package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/isaacwilkinsonlongden/blog-aggregator/internal/database"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("no username provided")
	}

	_, err := s.db.GetUser(context.Background(), cmd.args[0])
	if err != nil {
		return err
	}

	if err = s.cfg.SetUser(cmd.args[0]); err != nil {
		return err
	}

	fmt.Printf("Username set to %s\n", cmd.args[0])
	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("no username provided")
	}

	params := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.args[0],
	}

	user, err := s.db.CreateUser(context.Background(), params)
	if err != nil {
		return err
	}

	if err = s.cfg.SetUser(user.Name); err != nil {
		return err
	}
	fmt.Printf("Registered user %s\n", user.Name)
	fmt.Printf("%+v\n", user)

	return nil
}
