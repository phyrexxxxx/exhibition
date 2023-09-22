package main

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/phyrexxxxx/exhibition/internal/database"
)

// `startScraping“ starts the scraping process using the provided database connection,
// concurrency level, and time duration between each request.
//
// Parameters:
// - `db`: a pointer to the `database.Queries` struct representing the database connection.
// - `concurrency`: an integer indicating the number of goroutines to use for scraping.
// - `timeBetweenRequest`: a `time.Duration` specifying the duration between each scraping request.
func startScraping(db *database.Queries, concurrency int, timeBetweenRequest time.Duration) {
	log.Printf("Scraping on %v goroutines every %s duration", concurrency, timeBetweenRequest)

	// creates a new ticker that will fire every `timeBetweenRequest` duration
	ticker := time.NewTicker(timeBetweenRequest)
	for ; ; <-ticker.C {
		feeds, err := db.GetNextFeedsToFetch(
			context.Background(),
			int32(concurrency),
		)
		if err != nil {
			log.Println("error fetching feeds:", err)
			continue
		}

		wg := &sync.WaitGroup{}
		for _, feed := range feeds {
			wg.Add(1)
			go scrapeFeed(db, wg, feed)
		}
		// waits for all the goroutines to finish before repeating the process
		wg.Wait()
	}
}

func scrapeFeed(db *database.Queries, wg *sync.WaitGroup, feed database.Feed) {
	defer wg.Done()

	_, err := db.MarkFeedAsFetched(context.Background(), feed.ID)
	if err != nil {
		log.Println("error marking feed as fetched:", err)
		return
	}

	rssFeed, err := urlToFeed(feed.Url)
	if err != nil {
		log.Println("error fetching feed:", err)
		return
	}

	for _, item := range rssFeed.Channel.Item {
		log.Println("Found post:", item.Title, "on feed:", feed.Name)
	}
	log.Printf("Feed %s fetched, %v posts found", feed.Name, len(rssFeed.Channel.Item))
}
