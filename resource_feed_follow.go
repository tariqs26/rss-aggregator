package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/tariqs26/rss-aggregator/internal/database"
	"github.com/tariqs26/rss-aggregator/internal/util"
)

type FeedFollowResource struct{}

func (FeedFollowResource) Routes() chi.Router {
	router := chi.NewRouter()

	router.Get("/feed-follows", middlewareAuth(handleGetFeedFollows))
	router.Delete("/feed-follows/{id}", middlewareAuth(handleDeleteFeedFollow))

	return router
}

func handleCreateFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
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

func handleGetFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollows, err := DB.GetFeedFollows(r.Context(), user.ID)

	if err != nil {
		util.RespondWithError(w, http.StatusInternalServerError,
			fmt.Sprintf("Error getting feed follows: %v", err),
		)
		return
	}

	util.RespondWithJSON(w, http.StatusOK, feedFollows)
}

func handleDeleteFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	id, err := util.GetIntId(chi.URLParam(r, "id"))

	if err != nil {
		util.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	err = DB.DeleteFeedFollow(r.Context(), database.DeleteFeedFollowParams{
		ID:     id,
		UserID: user.ID,
	})

	if err != nil {
		util.RespondWithError(w, http.StatusInternalServerError,
			fmt.Sprintf("Error deleting feed follow: %v", err),
		)
		return
	}

	util.RespondWithJSON(w, http.StatusOK, "Feed un-followed successfully")
}
