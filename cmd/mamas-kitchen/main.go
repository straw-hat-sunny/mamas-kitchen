package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"mamas-kitchen/internal/audio"
	"mamas-kitchen/internal/azstorage"
	"mamas-kitchen/internal/recipe"
	"mamas-kitchen/internal/ui"

	"github.com/gorilla/mux"
)

func main() {
	log.Println("Starting Mama's Kitchen...")
	connectionString := "AccountName=devstoreaccount1;AccountKey=Eby8vdM02xNOcqFlqUwJPLlmEtlCDXJ1OUzFT50uSRZ6IFsuFq2UVErCz4I6tq/K1SZFPTOtr/KBHBeksoGMGw==;DefaultEndpointsProtocol=http;BlobEndpoint=http://azureite:10000/devstoreaccount1;QueueEndpoint=http://azureite:10001/devstoreaccount1;"
	ctx := context.Background()
	// create infrastructure
	queueService, err := azstorage.NewQueueService(connectionString)
	if err != nil {
		log.Fatalf("[%s] could not connect to queue store: %v", err)
	}

	blobService, err := azstorage.NewBlobService(connectionString)
	if err != nil {
		log.Fatalf("could not connect to blob store: %v", err)
	}

	blobClient, err := blobService.CreateBlobContainer(ctx, "audio-files")
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	queueClient, err := queueService.CreateQueue(ctx, "blob-uploaded")
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	_, err = queueService.CreateQueue(ctx, "blob-transcribed")
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	_, err = queueService.CreateQueue(ctx, "blob-translated")
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	_, err = queueService.CreateQueue(ctx, "storage-completed")
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	router := mux.NewRouter()

	// Recipe Data API
	recipeRouter := router.PathPrefix("/api/recipes").Subrouter()
	err = recipe.RegisterRoutes(recipeRouter)
	if err != nil {
		log.Fatalf("could not register recipe routes: %v", err)
	}

	// Audio Upload API
	audioRouter := router.PathPrefix("/api/audio").Subrouter()
	err = audio.RegisterRoutes(audioRouter, blobClient, queueClient)
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
