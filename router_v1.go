package main

import (
	"github.com/go-chi/chi/v5"
)

func v1Router() *chi.Mux {
	router := chi.NewRouter()

	router.Get("/ping", handlerPing)
	router.Mount("/users", userRoutes())
	router.Mount("/feeds", feedRoutes())
	router.Mount("/feed-follows", feedFollowRoutes())

	return router
}
