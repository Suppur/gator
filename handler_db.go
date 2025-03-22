package main

import (
	"context"
	"errors"
	"fmt"
)

func handlerReset(s *state, cmd command) error {
	if len(cmd.args) > 0 {
		return errors.New("error, command reset takes no args")
	}
	err := s.db.DeleteUsers(context.Background())
	if err != nil {
		return fmt.Errorf("error wiping database: %w", err)
	}

	fmt.Println("Success! Database has been wiped")

	return nil
}
