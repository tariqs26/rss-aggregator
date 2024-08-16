package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/tariqs26/rss-aggregator/internal/database"
	"github.com/tariqs26/rss-aggregator/internal/util"
)

func userRoutes() *chi.Mux {
	router := chi.NewRouter()

	router.Post("/", createUser)
	router.Get("/", authMiddleware(getUser))
	router.Delete("/{id}", authMiddleware((deleteUser)))
	router.Get("/posts", authMiddleware(getUserPosts))

	return router
}

func createUser(w http.ResponseWriter, r *http.Request) {
	type Body struct {
		Name string `json:"name"`
	}

	params, err := util.ValidateJSONBody(r, Body{})

	if err != nil {
		util.RespondWithError(w, http.StatusBadRequest,
			fmt.Sprintf("Error parsing JSON: %v", err),
		)
		return
	}

	user, err := DB.CreateUser(r.Context(), params.Name)

	if err != nil {
		util.RespondWithError(w, http.StatusInternalServerError,
			fmt.Sprintf("Error creating user: %v", err),
		)
		return
	}

	util.RespondWithJSON(w, http.StatusCreated, user)
}

func getUser(w http.ResponseWriter, r *http.Request, user database.User) {
	util.RespondWithJSON(w, http.StatusOK, user)
}

func deleteUser(w http.ResponseWriter, r *http.Request, user database.User) {
	err := DB.DeleteUser(r.Context(), user.ID)

	if err != nil {
		util.RespondWithError(w, http.StatusInternalServerError,
			fmt.Sprintf("User could not be deleted: %v", err),
		)
		return
	}

	util.RespondWithJSON(w, http.StatusOK, "User deleted successfully")
}

func getUserPosts(w http.ResponseWriter, r *http.Request, user database.User) {
	posts, err := DB.GetUserPosts(r.Context(), database.GetUserPostsParams{
		UserID: user.ID,
		Limit:  10,
	})

	if err != nil {
		util.RespondWithError(w, http.StatusInternalServerError,
			fmt.Sprintf("Error getting posts: %v", err),
		)
		return
	}

	util.RespondWithJSON(w, http.StatusOK, posts)
}
