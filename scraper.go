package main

import (
	"context"
	"database/sql"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/shaksiper/gorss/internal/database"
)

func startScraping(db *database.Queries, concurrency int, timeBetweenReq time.Duration) {
	log.Print("Scraping started")
	log.Printf("Scraping on %v goroutines in intervals of %s", concurrency, timeBetweenReq)

	ticker := time.NewTicker(timeBetweenReq)
	for ; ; <-ticker.C {
		feeds, err := db.GetNextFeedsToFetch(context.Background(), int32(concurrency))
		if err != nil {
			log.Print("Could not fetch feeds from DB")
			continue // scrape must go on
		}

		wg := &sync.WaitGroup{}

		for _, feed := range feeds {
			wg.Add(1)
			go scrapeFeed(db, wg, feed)
		}
		wg.Wait()

	}
}

func scrapeFeed(db *database.Queries, wg *sync.WaitGroup, feed database.Feed) RSSFeed {
	defer wg.Done()

	_, err := db.MarkFeedAsFetched(context.Background(), feed.ID)
	if err != nil {
		log.Printf("Error: could not mark feed [%v] as fetched: %v", feed.ID, err)
		return RSSFeed{}
	}

	rssFeed, err := urlToFeed(feed.Url)
	if err != nil {
		log.Printf("Error: fetching feed [%v] failed: %v", feed.ID, err)
		return RSSFeed{}
	}

	for _, item := range rssFeed.Channel.Item {
		description := sql.NullString{}
		if item.Description != "" {
			description.String = item.Description
			description.Valid = true
		}

		// User a more robust parsing
		pubDate, err := time.Parse(time.RFC1123Z, item.PubDate)
		if err != nil {
			log.Printf("Could not parse the publication date %v: %v", item.PubDate, err)
			pubDate = time.Now()
		}

		_, err = db.CreatePost(context.Background(), database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now().UTC(),
			UpdatedAt:   time.Now().UTC(),
			Title:       item.Title,
			Description: description,
			PublishedAt: pubDate,
			Url:         item.Link,
			FeedID:      feed.ID,
		})
		if err != nil {
			// do not log duplicate key errors
			if strings.Contains(err.Error(), "duplicate key") {
				continue
			}
			log.Printf("Could not create post for %v: %v", feed.ID, err)
		}
	}

	return rssFeed
}
