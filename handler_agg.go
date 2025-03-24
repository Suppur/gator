package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/Suppur/gator/internal/database"
	"github.com/google/uuid"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return errors.New("please enter the time delay between requests")
	}

	timeBetweenReqs, err := time.ParseDuration(cmd.args[0])
	if err != nil {
		return fmt.Errorf("failed to parse time delay: %w, please try again", err)
	}

	fmt.Printf("Collecting feeds every %v\n", timeBetweenReqs)

	ticker := time.NewTicker(timeBetweenReqs)
	for ; ; <-ticker.C {
		if err = scrapeFeeds(s); err != nil {
			log.Fatal(err)
		}
	}

}

func scrapeFeeds(s *state) error {
	next, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return fmt.Errorf("unable to fetch next feed to scrape from DB: %w", err)
	}

	dbMarkFetchedArgs := database.MarkFeedFetchedParams{
		ID: next.ID,
		LastFetchedAt: sql.NullTime{
			Time:  time.Now().Local(),
			Valid: true,
		},
		UpdatedAt: time.Now().Local(),
	}

	if err := s.db.MarkFeedFetched(context.Background(), dbMarkFetchedArgs); err != nil {
		return fmt.Errorf("failed to mark feed as fetched: %w", err)
	}

	feed, err := fetchFeed(context.Background(), next.Url)
	if err != nil {
		return fmt.Errorf("failed to fetch feed: %w", err)
	}

	for i, post := range feed.Channel.Item {
		pubTime, err := time.Parse(time.RFC1123, post.PubDate)
		if err != nil {
			return fmt.Errorf("pubdate parsing failed: %w", err)
		}
		dbCreatePostArgs := database.CreatePostParams{
			ID:        uuid.New(),
			CreatedAt: time.Now().Local(),
			UpdatedAt: time.Now().Local(),
			Title:     post.Title,
			Url:       post.Link,
			Description: sql.NullString{
				String: post.Description,
				Valid:  true,
			},
			PublishedAt: pubTime,
			FeedID: uuid.NullUUID{
				UUID:  next.ID,
				Valid: true,
			},
		}
		p, err := s.db.CreatePost(context.Background(), dbCreatePostArgs)
		if err != nil {
			if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
				continue
			}
			log.Printf("creating post failed: %v", err)
			continue
		}
		fmt.Printf("Post %v: %v inserted in the database!\n", i+1, p.Title)
	}

	/* fmt.Println("========== Printing feed titles: ==========")

	for i, v := range feed.Channel.Item {
		fmt.Printf("* #%v: %v\n", i+1, v.Title)
		fmt.Println("-------------------------------------------")
	} */

	return nil
}
