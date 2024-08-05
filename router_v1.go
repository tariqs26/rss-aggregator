package main

import (
	"github.com/go-chi/chi/v5"
)

func v1Router() *chi.Mux {
	router := chi.NewRouter()

	router.Get("/healthz", handleReady)
	router.Get("/err", handleError)

	router.Mount("/users", UserResource{}.Routes())
	router.Mount("/feeds", FeedResource{}.Routes())
	router.Mount("/feed-follows", FeedFollowResource{}.Routes())

	return router
}
