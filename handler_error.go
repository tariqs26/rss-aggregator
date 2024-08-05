package main

import (
	"net/http"
	"github.com/tariqs26/rss-aggregator/internal/util"
)

func handleError(w http.ResponseWriter, _ *http.Request) {
	util.RespondWithError(w, http.StatusBadRequest, "Something went wrong")
}
