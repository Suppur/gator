package main

import (
	"context"
	"fmt"

	"github.com/Suppur/gator/internal/database"
)

func middlewareLoggedIn(handler func(s *state, cmd command, usr database.User) error) func(*state, command) error {
	return func(s *state, cmd command) error {
		usr, err := s.db.GetUser(context.Background(), s.conf.CurrentUserName)
		if err != nil {
			return fmt.Errorf("fetching user from DB failed %w", err)
		}
		err = handler(s, cmd, usr)
		return err
	}
}
