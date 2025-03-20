package main

import (
	"errors"
	"fmt"
	"log"

	"github.com/Suppur/gator/internal/config"
	"github.com/Suppur/gator/internal/database"
)

type state struct {
	db   *database.Queries
	conf *config.Config
}

type command struct {
	name string
	args []string
}

type commands struct {
	cmdList map[string]func(*state, command) error
}

func handlerLogin(s *state, cmd command) (err error) {
	if len(cmd.args) != 1 {
		return errors.New("error, please enter a username")
	}

	if err := s.conf.SetUser(cmd.args[0]); err != nil {
		return err
	}

	fmt.Printf("User %v has been set \n", s.conf.CurrentUserName)

	return nil
}

func (c *commands) register(name string, f func(*state, command) error) {
	if name == "" {
		log.Fatal("error, enter a command name")
	}

	c.cmdList[name] = f
}

func (c *commands) run(s *state, cmd command) error {
	if cmd.name == "" || len(cmd.args) != 1 {
		return errors.New("error, please enter a valid cmd name or valid cmd argument")
	}

	_, found := c.cmdList[cmd.name]
	if !found {
		return errors.New("error, command not found")
	} else {
		handler := c.cmdList[cmd.name]
		if err := handler(s, cmd); err != nil {
			return fmt.Errorf("error! %v", err)
		}
	}

	return nil
}
