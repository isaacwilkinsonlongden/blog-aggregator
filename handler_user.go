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

func handerlerReset(s *state, cmd command) error {
	err := s.db.DeleteUsers(context.Background())
	if err != nil {
		return err
	}

	fmt.Println("Reset table: 'users'")
	return nil
}

func handlerUsers(s *state, cmd command) error {
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return err
	}

	for _, user := range users {
		if user.Name == s.cfg.CurrentUserName {
			fmt.Printf("* %s (current)\n", user.Name)
		} else {
			fmt.Printf("* %s\n", user.Name)
		}
	}

	return nil
}
