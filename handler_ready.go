package main

import (
	"net/http"
	"github.com/tariqs26/rss-aggregator/internal/util"
)

func handleReady(w http.ResponseWriter, _ *http.Request) {
	util.RespondWithJSON(w, http.StatusOK, map[string]bool{"ok": true})
}
