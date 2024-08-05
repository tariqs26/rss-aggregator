package main

import (
	"fmt"
	"net/http"

	"github.com/tariqs26/rss-aggregator/internal/auth"
	"github.com/tariqs26/rss-aggregator/internal/database"
	"github.com/tariqs26/rss-aggregator/internal/util"
)

type AuthHandler func(http.ResponseWriter, *http.Request, database.User)

func middlewareAuth(handler AuthHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetApiKey(r.Header)

		if err != nil {
			util.RespondWithError(w, http.StatusBadRequest,
				fmt.Sprintf("auth error: %v", err),
			)
			return
		}

		user, err := DB.GetUserByApiKey(r.Context(), apiKey)

		if err != nil {
			util.RespondWithError(w, http.StatusInternalServerError,
				"Error getting user",
			)
			return
		}

		handler(w, r, user)
	}
}
