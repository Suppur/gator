package main

import (
	"errors"
	"fmt"
	"log"
)

type command struct {
	name string
	args []string
}

type commands struct {
	cmdList map[string]func(*state, command) error
}

func (c *commands) register(name string, f func(*state, command) error) {
	if name == "" {
		log.Fatal("error, enter a command name")
	}

	c.cmdList[name] = f
}

func (c *commands) run(s *state, cmd command) error {
	if cmd.name == "" || len(cmd.args) > 1 {
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
