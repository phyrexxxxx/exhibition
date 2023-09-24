package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

	"github.com/phyrexxxxx/exhibition/auth"
	"github.com/phyrexxxxx/exhibition/config"
	db "github.com/phyrexxxxx/exhibition/database"
	"github.com/phyrexxxxx/exhibition/handlers"
	"github.com/phyrexxxxx/exhibition/internal/database"
	"github.com/phyrexxxxx/exhibition/router"
	"github.com/phyrexxxxx/exhibition/src"
)

func main() {
	godotenv.Load(".env")

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT environment variable not set")
	}

	// initializes a database connection
	db := db.InitDB()

	// Embedding
	dbQueries := database.New(db)
	apiCfg := config.ApiConfig{
		DB: dbQueries,
	}
	handlerApiCfg := handlers.HandlerApiConfig{
		ApiConfig: &apiCfg,
	}
	authApiCfg := auth.AuthApiConfig{
		ApiConfig: &apiCfg,
	}

	// starts a scraping process in a separate goroutine
	go src.StartScraping(dbQueries, 10, time.Minute)

	// initializes a router
	router := router.InitRouter(handlerApiCfg, authApiCfg)

	server := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}

	log.Printf("Server serving on port %v", portString)
	log.Fatal(server.ListenAndServe())
}
