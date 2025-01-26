package audio

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type AudioHandler struct {
	blobClient  BlobClient
	queueClient QueueClient
}

type BlobClient interface {
	Upload(ctx context.Context, fileName string, data io.Reader) error
}
type QueueClient interface {
	EnqueueMessage(ctx context.Context, msg string) error
}

func NewAudioHandler(bc BlobClient, qc QueueClient) (*AudioHandler, error) {
	return &AudioHandler{
		blobClient:  bc,
		queueClient: qc,
	}, nil
}

func (ah AudioHandler) HandleUploadToAzure(w http.ResponseWriter, r *http.Request) {
	log.Println("Uploading file to Azure Blob Storage...")
	err := r.ParseMultipartForm(10 << 20) // limit upload size to 10MB
	if err != nil {
		log.Println("Error parsing form")
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}

	file, handler, err := r.FormFile("file")
	if err != nil {
		log.Println("Error retrieving the file")
		http.Error(w, "Error retrieving the file", http.StatusBadRequest)
		return
	}

	defer file.Close()

	// upload file to blob store
	err = ah.blobClient.Upload(context.Background(), handler.Filename, file)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Unable to upload the file", http.StatusInternalServerError)
		return
	}
	msg := &AudioMessage{
		FileName: handler.Filename,
	}

	msgBytes, err := json.Marshal(msg)
	if err != nil {
		log.Println("Error marshalling the message")
		http.Error(w, "Error marshalling the message", http.StatusInternalServerError)
		return
	}

	err = ah.queueClient.EnqueueMessage(context.Background(), string(msgBytes))
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Unable to enqueue the message", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	log.Println("File uploaded successfully to Azure Blob Storage")
	w.Write([]byte("File uploaded successfully to Azure Blob Storage"))
}
