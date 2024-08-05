package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/tariqs26/rss-aggregator/internal/database"
	"github.com/tariqs26/rss-aggregator/internal/util"
)

func feedRoutes() chi.Router {
	router := chi.NewRouter()

	router.Post("/", middlewareAuth(createFeed))
	router.Get("/", middlewareAuth(getFeeds))
	router.Delete("/{id}", middlewareAuth(deleteFeed))
	router.Post("/{id}/follow", middlewareAuth(createFeedFollow))

	return router
}

func createFeed(w http.ResponseWriter, r *http.Request, user database.User) {
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

func getFeeds(w http.ResponseWriter, r *http.Request, user database.User) {
	feeds, err := DB.GetFeeds(r.Context(), user.ID)

	if err != nil {
		util.RespondWithError(w, http.StatusInternalServerError,
			fmt.Sprintf("Error getting feeds: %v", err),
		)
		return
	}

	util.RespondWithJSON(w, http.StatusOK, feeds)
}

func deleteFeed(w http.ResponseWriter, r *http.Request, user database.User) {
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

func createFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	feedIdParam := chi.URLParam(r, "id")

	feedId, err := strconv.Atoi(feedIdParam)

	if err != nil {
		util.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Invalid feed ID: %v", feedIdParam))
		return
	}

	feed, err := DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		FeedID: int32(feedId),
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
