package main

import (
	"log"
	"net/http"
	"time"

	"mamas-kitchen/internal/audio"
	"mamas-kitchen/internal/recipe"
	"mamas-kitchen/internal/ui"

	"github.com/gorilla/mux"
)

func main() {
	log.Println("Starting Mama's Kitchen...")
	router := mux.NewRouter()

	// Recipe Data API
	recipeRouter := router.PathPrefix("/api/recipes").Subrouter()
	err := recipe.RegisterRoutes(recipeRouter)
	if err != nil {
		log.Fatalf("could not register recipe routes: %v", err)
	}

	// Audio Upload API
	audioRouter := router.PathPrefix("/api/audio").Subrouter()
	err = audio.RegisterRoutes(audioRouter)
	if err != nil {
		log.Fatalf("could not register audio routes: %v", err)
	}

	// Serve the frontend
	ui.RegisterRoutes(router)

	srv := &http.Server{
		Handler: router,
		Addr:    "0.0.0.0:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
