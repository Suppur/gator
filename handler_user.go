package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/Suppur/gator/internal/database"
	"github.com/google/uuid"
)

func handlerLogin(s *state, cmd command) (err error) {
	if len(cmd.args) != 1 {
		return errors.New("error, please enter a username")
	}

	_, ok := s.db.GetUser(context.Background(), cmd.args[0])
	if ok != nil {
		log.Fatal(err)
	}

	if err := s.conf.SetUser(cmd.args[0]); err != nil {
		return err
	}

	fmt.Printf("User %v has been set \n", s.conf.CurrentUserName)

	return nil
}

func handlerRegister(s *state, cmd command) (err error) {
	if len(cmd.args) != 1 {
		return errors.New("error, please enter a user to register")
	}
	dbCreateUserArgs := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().Local(),
		UpdatedAt: time.Now().Local(),
		Name:      cmd.args[0],
	}
	_, ok := s.db.GetUser(context.Background(), dbCreateUserArgs.Name)
	if ok == nil {
		log.Fatal(err)
	}

	user, err := s.db.CreateUser(context.Background(), dbCreateUserArgs)
	if err != nil {
		return errors.New("error, user creation failed")
	}

	if err := s.conf.SetUser(user.Name); err != nil {
		return errors.New("error setting user")
	}

	fmt.Printf("User %v was created!\n", user.Name)
	fmt.Printf("User ID: %v \nUser created at: %v\nUser updated at: %v\nUser name: %v\n", user.ID, user.CreatedAt, user.UpdatedAt, user.Name)

	return nil
}

func handlerListUsers(s *state, cmd command) (err error) {
	usrs, err := s.db.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("error listing users: %w", err)
	}

	if len(usrs) < 1 {
		return errors.New("no users to list. register users first with the <register> command")
	}

	for _, v := range usrs {
		if v.Name == s.conf.CurrentUserName {
			fmt.Println("* " + v.Name + " (current)")
		} else {
			fmt.Println("* " + v.Name)
		}
	}

	return nil
}
