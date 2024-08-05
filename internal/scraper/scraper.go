package scraper

import (
	"context"
	"database/sql"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/tariqs26/rss-aggregator/internal/database"
)

func StartScraping(db *database.Queries, concurrency int, timeBetweenRequest time.Duration) {

	log.Printf("Scarping on %v goroutines every %s duration", concurrency, timeBetweenRequest)
	ticker := time.NewTicker(timeBetweenRequest)

	for ; ; <-ticker.C {
		feeds, err := db.GetNextFeedToFetch(
			context.Background(),
			int32(concurrency),
		)

		if err != nil {
			log.Println("Error fetching feeds: ", err)
			continue
		}

		wg := sync.WaitGroup{}

		for _, feed := range feeds {
			wg.Add(1)
			go scrapeFeed(db, &wg, feed)
		}
		wg.Wait()
	}

}

func scrapeFeed(db *database.Queries, wg *sync.WaitGroup, feed database.Feed) {
	defer wg.Done()
	_, err := db.MarkFeedAsFetched(context.Background(), feed.ID)
	if err != nil {
		log.Println("Error marking feed as fetched:", err)
	}

	rssFeed, err := urlToFeed(feed.Url)

	if err != nil {
		log.Println("Error fetching feed:", err)
		return
	}

	for _, item := range rssFeed.Channel.Item {
		description := sql.NullString{}

		if item.Description != "" {
			description = sql.NullString{String: item.Description, Valid: true}
		}

		publishedAt, err := time.Parse(time.RFC1123Z, item.PubDate)

		if err != nil {
			log.Println("Error parsing time:", err)
			continue
		}

		_, err = db.CreatePost(context.Background(), database.CreatePostParams{
			Title:       item.Title,
			Description: description,
			Url:         feed.Url,
			PublishedAt: publishedAt,
			FeedID:      feed.ID,
		})

		if err != nil {
			if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
				continue

			}
			log.Println("Error creating post:", err)
		}
	}
	log.Printf("Feed %s collected, %v posts found", feed.Name, len(rssFeed.Channel.Item))
}
