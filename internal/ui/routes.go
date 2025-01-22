package ui

import "github.com/gorilla/mux"

func RegisterRoutes(router *mux.Router) {
	spa := spaHandler{staticPath: "frontend/dist", indexPath: "index.html"}
	router.PathPrefix("/").Handler(spa)
}
