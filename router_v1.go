package main

import (
	"github.com/go-chi/chi/v5"
)

func v1Router() *chi.Mux {
	router := chi.NewRouter()

	router.Get("/healthz", handleReady)
	router.Get("/err", handleError)

	router.Mount("/users", userResource{}.Routes())
	router.Mount("/feeds", feedResource{}.Routes())
	router.Mount("/feed-follows", feedFollowResource{}.Routes())

	return router
}
