package main

import (
	"log"
	"time"

	"github.com/phyrexxxxx/exhibition/internal/database"
)

func startScraping(db *database.Queries, concurrency int, timeBetweenRequest time.Duration) {
	log.Printf("Scraping on %v goroutines every %s duration", concurrency, timeBetweenRequest)

	// creates a new ticker that will fire every `timeBetweenRequest` duration
	ticker := time.NewTicker(timeBetweenRequest)
	for ; ; <-ticker.C {

	}
}
