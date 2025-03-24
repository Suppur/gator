package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/Suppur/gator/internal/database"
)

func handlerBrowse(s *state, cmd command, usr database.User) (err error) {
	limit := int32(2)

	if len(cmd.args) > 0 {
		parsedLimit, err := strconv.ParseInt(cmd.args[0], 10, 32)
		if err != nil {
			return fmt.Errorf("error converting arg string to limit int: %w", err)
		}
		limit = int32(parsedLimit)
	}

	dbPostsArgs := database.GetPostsForUserParams{
		ID:    usr.ID,
		Limit: int32(limit),
	}

	posts, err := s.db.GetPostsForUser(context.Background(), dbPostsArgs)
	if err != nil {
		return fmt.Errorf("failed to get followed feeds")
	}

	fmt.Println("========== Posts in followed feeds: ==========")

	for i, post := range posts {
		fmt.Printf("* Post #%v: %v\n", i+1, post.Title)
		fmt.Printf("* Published %v at %v\n", post.PublishedAt.Format("Mon Jan 2"), post.Url)
		fmt.Printf("* Description: %+v\n", post.Description.String)
		fmt.Println("----------------------------------------------")
	}

	return nil
}
