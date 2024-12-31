package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/tariqs26/rss-aggregator/internal/database"
	"github.com/tariqs26/rss-aggregator/internal/util"
)

func feedFollowRoutes() *chi.Mux {
	router := chi.NewRouter()

	router.Get("/", authMiddleware(getFeedFollows))
	router.Delete("/{id}", authMiddleware(deleteFeedFollow))

	return router
}

func getFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollows, err := DB.GetFeedFollows(r.Context(), user.ID)

	if err != nil {
		util.RespondWithError(w, http.StatusInternalServerError,
			fmt.Sprintf("Error getting feed follows: %v", err),
		)
		return
	}

	util.RespondWithJSON(w, http.StatusOK, feedFollows)
}

func deleteFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
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
