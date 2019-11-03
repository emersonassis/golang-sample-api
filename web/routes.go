package web

import (
	mux "github.com/gorilla/mux"
	"projects.org/sample/sample-api/core"
)

//Router ...
func Router(h *Handler) *mux.Router {
	router := core.Router()

	router.HandleFunc("/api/v1/version/", h.HandlerVersion).Methods("GET")
	router.HandleFunc("/api/v1/process/", h.HandleProcessObject).Methods("POST")

	return router
}
