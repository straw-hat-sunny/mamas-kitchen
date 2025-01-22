package audio

import "github.com/gorilla/mux"

func RegisterRoutes(router *mux.Router) error {

	ah, err := NewAudioHandler()
	if err != nil {
		return err
	}

	router.HandleFunc("/upload", ah.HandleUploadToAzure).Methods("POST")
	return nil

}
