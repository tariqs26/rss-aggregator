package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/tariqs26/rss-aggregator/internal/database"
	"github.com/tariqs26/rss-aggregator/internal/util"
)

type FeedResource struct{}

func (FeedResource) Routes() chi.Router {
	router := chi.NewRouter()

	router.Post("/", middlewareAuth(handleCreateFeed))
	router.Get("/", middlewareAuth(handleGetFeeds))
	router.Delete("/{id}", middlewareAuth(handleDeleteFeed))
	router.Post("/{id}/follow", middlewareAuth(handleCreateFeedFollow))

	return router
}

func handleCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type Params struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}

	params, err := util.ValidateJSONBody(r, Params{})

	if err != nil {
		util.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	feed, err := DB.CreateFeed(r.Context(), database.CreateFeedParams{
		Url:    params.Url,
		Name:   params.Name,
		UserID: user.ID,
	})

	if err != nil {
		util.RespondWithError(w, http.StatusInternalServerError,
			fmt.Sprintf("Error creating feed: %v", err),
		)
		return
	}

	util.RespondWithJSON(w, http.StatusCreated, feed)
}

func handleGetFeeds(w http.ResponseWriter, r *http.Request, user database.User) {
	feeds, err := DB.GetFeeds(r.Context(), user.ID)

	if err != nil {
		util.RespondWithError(w, http.StatusInternalServerError,
			fmt.Sprintf("Error getting feeds: %v", err),
		)
		return
	}

	util.RespondWithJSON(w, http.StatusOK, feeds)
}

func handleDeleteFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	id, err := util.GetIntId(chi.URLParam(r, "id"))

	if err != nil {
		util.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	err = DB.DeleteFeed(r.Context(),
		database.DeleteFeedParams{
			ID:     id,
			UserID: user.ID,
		})

	if err != nil {
		util.RespondWithError(w, http.StatusInternalServerError,
			fmt.Sprintf("Feed could not be deleted: %v", err),
		)
		return
	}

	util.RespondWithJSON(w, http.StatusOK, "Feed deleted successfully")
}
