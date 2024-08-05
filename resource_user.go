package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/tariqs26/rss-aggregator/internal/database"
	"github.com/tariqs26/rss-aggregator/internal/util"
)

type userResource struct{}

func (userResource) Routes() chi.Router {
	router := chi.NewRouter()

	router.Post("/", handleCreateUser)
	router.Get("/", middlewareAuth(handleGetUserByApiKey))
	router.Delete("/{id}", handleDeleteUser)
	router.Get("/posts", middlewareAuth(handleGetUserPosts))

	return router
}

func handleCreateUser(w http.ResponseWriter, r *http.Request) {
	type Params struct {
		Name string `json:"name"`
	}

	params, err := util.ValidateJSONBody(r, Params{})

	if err != nil {
		util.RespondWithError(w, http.StatusBadRequest,
			fmt.Sprintf("Error parsing JSON: %v", err),
		)
		return
	}

	user, err := apiConfig.DB.CreateUser(r.Context(), params.Name)

	if err != nil {
		util.RespondWithError(w, http.StatusInternalServerError,
			fmt.Sprintf("Error creating user: %v", err),
		)
		return
	}

	util.RespondWithJSON(w, http.StatusCreated, user)
}

func handleGetUserByApiKey(w http.ResponseWriter, r *http.Request, user database.User) {
	util.RespondWithJSON(w, http.StatusOK, user)
}

func handleDeleteUser(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")

	id, err := uuid.Parse(idParam)

	if err != nil {
		util.RespondWithError(w, http.StatusBadRequest,
			fmt.Sprintf("Error parsing UUID: %v", err),
		)
		return
	}

	err = apiConfig.DB.DeleteUser(r.Context(), id)

	if err != nil {
		util.RespondWithError(w, http.StatusInternalServerError,
			fmt.Sprintf("User could not be deleted: %v", err),
		)
		return
	}

	util.RespondWithJSON(w, http.StatusOK, "User deleted successfully")
}

func handleGetUserPosts(w http.ResponseWriter, r *http.Request, user database.User) {
	posts, err := apiConfig.DB.GetUserPosts(r.Context(), database.GetUserPostsParams{
		UserID: user.ID,
		Limit:  10,
	})

	if err != nil {
		util.RespondWithError(w, http.StatusInternalServerError,
			fmt.Sprintf("Error getting user posts: %v", err),
		)
		return
	}

	util.RespondWithJSON(w, http.StatusOK, posts)
}