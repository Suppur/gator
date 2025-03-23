package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/Suppur/gator/internal/database"
	"github.com/google/uuid"
)

func handlerAddFeed(s *state, cmd command, usr database.User) (err error) {
	if len(cmd.args) != 2 {
		return errors.New("please enter a name and a URL")
	}
	/* usr, err := s.db.GetUser(context.Background(), s.conf.CurrentUserName)
	if err != nil {
		return fmt.Errorf("fetching user from DB failed %w", err)
	} */

	dbAddFeedArgs := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().Local(),
		UpdatedAt: time.Now().Local(),
		Name:      cmd.args[0],
		Url:       cmd.args[1],
		UserID: uuid.NullUUID{
			UUID:  usr.ID,
			Valid: true,
		},
	}

	feed, err := s.db.CreateFeed(context.Background(), dbAddFeedArgs)
	if err != nil {
		return fmt.Errorf("feed creation failed: %w", err)
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

	if _, err := s.db.CreateFeedFollow(context.Background(), dbFeedFollowParams); err != nil {
		return fmt.Errorf("feed follow failed: %w", err)
	}

	fmt.Printf("Feed successfully created!\n\n")
	fmt.Printf("Feed ID: %v\n Feed created at: %v\n Feed updated at: %v\n Feed name: %v\n Feed URL: %v\n Linked user: %v\n", feed.ID, feed.CreatedAt, feed.UpdatedAt, feed.Name, feed.Url, usr.Name)

	return nil
}

func handlerListFeeds(s *state, cmd command) error {
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("fetching feeds from DB failed: %w", err)
	}

	if len(feeds) == 0 {
		fmt.Println("No feeds have been found. Add some feeds first!")
		return nil
	}

	fmt.Println("======== Feeds: ========")
	for i, feed := range feeds {
		usrName, err := s.db.GetUsersByID(context.Background(), feed.UserID.UUID)
		if err != nil {
			return fmt.Errorf("user ID retrieval from DB failed: %w", err)
		}
		fmt.Printf("Feed #: %v\n", i+1)
		fmt.Printf("* Feed Name: %v\n * Feed URL: %v\n * Feed created by user: %v\n", feed.Name, feed.Url, usrName)
		fmt.Println("------------------------")
	}

	return nil
}
