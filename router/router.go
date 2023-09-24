package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/phyrexxxxx/exhibition/auth"
	"github.com/phyrexxxxx/exhibition/handlers"
)

func InitRouter(handlerApiCfg handlers.HandlerApiConfig, authApiCfg auth.AuthApiConfig) *chi.Mux {
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
	v1Router.Delete("/feed_follows/{feedFollowID}", authApiCfg.MiddlewareAuth(handlerApiCfg.HandlerDeleteFeedFollow))
	v1Router.Get("/posts", authApiCfg.MiddlewareAuth(handlerApiCfg.HandlerGetPostsByUser))

	router.Mount("/v1", v1Router)

	return router
}
