package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/Suppur/gator/internal/database"
	"github.com/google/uuid"
)

func handlerFollow(s *state, cmd command, usr database.User) (err error) {
	if len(cmd.args) != 1 {
		return errors.New("please enter a URL")
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

func handlerListFollows(s *state, cmd command, usr database.User) (err error) {
	if len(cmd.args) != 0 {
		return errors.New("error, the following command takes no arguments")
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

func handlerUnfollow(s *state, cmd command, usr database.User) (err error) {
	if len(cmd.args) != 1 {
		return errors.New("error, the unfollow command requires the URL of the feed to unfollow")
	}

	feed, err := s.db.GetFeedByURL(context.Background(), cmd.args[0])
	if err != nil {
		return fmt.Errorf("failed to retrieve feed with URL provided: %w", err)
	}

	dbDeleteFeedFollowParams := database.DeleteFeedFollowByUserFeedParams{
		UserID: uuid.NullUUID{
			UUID:  usr.ID,
			Valid: true,
		},
		FeedID: uuid.NullUUID{
			UUID:  feed.ID,
			Valid: true,
		},
	}

	if err := s.db.DeleteFeedFollowByUserFeed(context.Background(), dbDeleteFeedFollowParams); err != nil {
		return fmt.Errorf("failed to delete feed from followed feeds: %w", err)
	}

	fmt.Println("Feed successfully deleted from followed feeds!")

	return nil
}
