package audio

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azqueue"
)

type AudioHandler struct{}

func NewAudioHandler() (*AudioHandler, error) {
	return &AudioHandler{}, nil
}

func (ah AudioHandler) HandleUploadToAzure(w http.ResponseWriter, r *http.Request) {

	err := r.ParseMultipartForm(10 << 20) // limit upload size to 10MB
	if err != nil {
		println("Error parsing form")
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}

	file, handler, err := r.FormFile("file")
	if err != nil {
		println("Error retrieving the file")
		http.Error(w, "Error retrieving the file", http.StatusBadRequest)
		return
	}

	defer file.Close()

	connectionString := "AccountName=devstoreaccount1;AccountKey=Eby8vdM02xNOcqFlqUwJPLlmEtlCDXJ1OUzFT50uSRZ6IFsuFq2UVErCz4I6tq/K1SZFPTOtr/KBHBeksoGMGw==;DefaultEndpointsProtocol=http;BlobEndpoint=http://127.0.0.1:10000/devstoreaccount1;QueueEndpoint=http://127.0.0.1:10001/devstoreaccount1;TableEndpoint=http://127.0.0.1:10002/devstoreaccount1;"
	client, err := azblob.NewClientFromConnectionString(connectionString, nil)
	if err != nil {
		println("Unable to create a new client")
		http.Error(w, "Unable to create a new client", http.StatusInternalServerError)
		return
	}

	queueClient, err := azqueue.NewQueueClientFromConnectionString(connectionString, "audio-files", nil)
	if err != nil {
		println("Unable to create a new queue client")
		http.Error(w, "Unable to create a new queue client", http.StatusInternalServerError)
		return
	}

	// Create a container client
	_, err = client.UploadStream(context.Background(), "audio-files", handler.Filename, file, nil)
	if err != nil {
		println(err.Error())
		http.Error(w, "Unable to upload the file", http.StatusInternalServerError)
		return
	}
	msg := &AudioMessage{
		FileName: handler.Filename,
	}

	msgBytes, err := json.Marshal(msg)
	if err != nil {
		println("Error marshalling the message")
		http.Error(w, "Error marshalling the message", http.StatusInternalServerError)
		return
	}

	_, err = queueClient.EnqueueMessage(context.Background(), string(msgBytes), nil)
	if err != nil {
		println(err.Error())
		http.Error(w, "Unable to enqueue the message", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	println("File uploaded successfully to Azure Blob Storage")
	w.Write([]byte("File uploaded successfully to Azure Blob Storage"))
}
