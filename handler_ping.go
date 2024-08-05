package main

import (
	"github.com/tariqs26/rss-aggregator/internal/util"
	"net/http"
)

func handlerPing(w http.ResponseWriter, _ *http.Request) {
	util.RespondWithJSON(w, http.StatusOK, map[string]string{"message": "pong"})
}
