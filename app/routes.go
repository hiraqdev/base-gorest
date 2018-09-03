package app

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hiraqdev/base-gorest/app/modules/ping"
)

func routes(h *mux.Router) http.Handler {
	h.HandleFunc("/ping", ping.Handler).Methods("GET")

	return h
}
