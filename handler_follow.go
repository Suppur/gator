package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/Suppur/gator/internal/database"
	"github.com/google/uuid"
)

func handlerFollow(s *state, cmd command) (err error) {
	if len(cmd.args) != 1 {
		return errors.New("please enter a URL")
	}

	usr, err := s.db.GetUser(context.Background(), s.conf.CurrentUserName)
	if err != nil {
		return fmt.Errorf("fetching user from DB failed %w", err)
	}

	feed, err := s.db.GetFeedByURL(context.Background(), cmd.args[0])
	if err != nil {
		return fmt.Errorf("feed retrieval failed: %w", err)
	}

	dbFeedFollowParams := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().Local(),
		UpdatedAt: time.Now().Local(),
		UserID: uuid.NullUUID{
			UUID:  usr.ID,
			Valid: true,
		},
		FeedID: uuid.NullUUID{
			UUID:  feed.ID,
			Valid: true,
		},
	}

	feedFollow, err := s.db.CreateFeedFollow(context.Background(), dbFeedFollowParams)
	if err != nil {
		return fmt.Errorf("feed follow failed: %w", err)
	}

	fmt.Println("========== New feed followed: ==========")
	fmt.Printf("* New feed followed: %v\n", feedFollow.FeedName)
	fmt.Printf("* By user:           %v\n", feedFollow.UserName)
	fmt.Println("----------------------------------------")

	return nil
}

func handlerListFollows(s *state, cmd command) (err error) {
	if len(cmd.args) != 0 {
		return errors.New("error, the following command takes no arguments")
	}

	usr, err := s.db.GetUser(context.Background(), s.conf.CurrentUserName)
	if err != nil {
		return fmt.Errorf("fetching user from DB failed %w", err)
	}

	usrFeed := uuid.NullUUID{
		UUID:  usr.ID,
		Valid: true,
	}

	follows, err := s.db.GetFeedFollowsForUser(context.Background(), usrFeed)
	if err != nil {
		return fmt.Errorf("fetching followed feeds failed: %w", err)
	}

	fmt.Println("========== You are currently following: ==========")

	for _, follow := range follows {
		fmt.Printf("* %v\n", follow.FeedName)
	}

	fmt.Println("--------------------------------------------------")

	return nil
}
