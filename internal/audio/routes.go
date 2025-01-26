package audio

import "github.com/gorilla/mux"

func RegisterRoutes(router *mux.Router, bc BlobClient, qc QueueClient) error {

	ah, err := NewAudioHandler(bc, qc)
	if err != nil {
		return err
	}

	router.HandleFunc("/upload", ah.HandleUploadToAzure).Methods("POST")
	return nil

}
