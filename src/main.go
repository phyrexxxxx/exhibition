package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

	"github.com/phyrexxxxx/exhibition/auth"
	"github.com/phyrexxxxx/exhibition/config"
	"github.com/phyrexxxxx/exhibition/handlers"
	"github.com/phyrexxxxx/exhibition/internal/database"
)

func main() {
	godotenv.Load("../.env")

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT environment variable not set")
	}

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL environment variable not set")
	}

	// opens a connection to a PostgreSQL database
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(err)
	}

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

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1Router := chi.NewRouter()
	v1Router.Get("/healthz", handlers.HandlerReadiness)
	v1Router.Get("/err", handlers.HandlerErr)
	v1Router.Post("/users", handlerApiCfg.HandlerCreateUser)
	v1Router.Get("/users", authApiCfg.MiddlewareAuth(handlerApiCfg.HandlerGetUser))
	v1Router.Post("/feeds", authApiCfg.MiddlewareAuth(handlerApiCfg.HandlerCreateFeed))
	v1Router.Get("/feeds", handlerApiCfg.HandlerGetAllFeeds)
	v1Router.Post("/feed_follows", authApiCfg.MiddlewareAuth(handlerApiCfg.HandlerCreateFeedFollow))
	v1Router.Get("/feed_follows", authApiCfg.MiddlewareAuth(handlerApiCfg.HandlerGetFeedFollows))

	router.Mount("/v1", v1Router)

	server := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}

	log.Printf("Server serving on port %v", portString)
	log.Fatal(server.ListenAndServe())
}
