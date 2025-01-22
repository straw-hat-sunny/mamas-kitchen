package audio

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azqueue"
)

type AudioHandler struct{}

func NewAudioHandler() (*AudioHandler, error) {
	return &AudioHandler{}, nil
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

	connectionString := "AccountName=devstoreaccount1;AccountKey=Eby8vdM02xNOcqFlqUwJPLlmEtlCDXJ1OUzFT50uSRZ6IFsuFq2UVErCz4I6tq/K1SZFPTOtr/KBHBeksoGMGw==;DefaultEndpointsProtocol=http;BlobEndpoint=http://azureite:10000/devstoreaccount1;QueueEndpoint=http://azureite:10001/devstoreaccount1;"
	client, err := azblob.NewClientFromConnectionString(connectionString, nil)
	if err != nil {
		log.Println("Unable to create a new client")
		http.Error(w, "Unable to create a new client", http.StatusInternalServerError)
		return
	}
	log.Println("listing containers")
	// List all containers in the account
	pager := client.NewListContainersPager(nil)
	for pager.More() {
		resp, err := pager.NextPage(context.Background())
		if err != nil {
			log.Println("Error listing containers")
			http.Error(w, "Error listing containers", http.StatusInternalServerError)
			return
		}
		if len(resp.ContainerItems) == 0 {
			log.Println("No containers found")
			_, err = client.CreateContainer(context.Background(), "audio-files", nil)
			if err != nil {
				log.Println("Error creating container")
				http.Error(w, "Error creating container", http.StatusInternalServerError)
				return
			}
		}
	}

	log.Println("Listing queues")
	queueServiceClient, err := azqueue.NewServiceClientFromConnectionString(connectionString, nil)
	if err != nil {
		log.Println("Unable to create a new queue service client")
		http.Error(w, "Unable to create a new queue service client", http.StatusInternalServerError)
		return
	}

	queuePager := queueServiceClient.NewListQueuesPager(nil)
	for queuePager.More() {
		queueResp, err := queuePager.NextPage(context.Background())
		if err != nil {
			log.Println("Error listing queues")
			http.Error(w, "Error listing queues", http.StatusInternalServerError)
			return
		}
		if len(queueResp.Queues) == 0 {
			log.Println("No queues found")
			_, err = queueServiceClient.CreateQueue(context.Background(), "audio-files", nil)
			if err != nil {
				log.Println("Error creating queue")
				http.Error(w, "Error creating queue", http.StatusInternalServerError)
				return
			}
		}
	}
	queueClient, err := azqueue.NewQueueClientFromConnectionString(connectionString, "audio-files", nil)
	if err != nil {
		log.Println("Unable to create a new queue client")

	}

	// Create a container client
	_, err = client.UploadStream(context.Background(), "audio-files", handler.Filename, file, nil)
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

	_, err = queueClient.EnqueueMessage(context.Background(), string(msgBytes), nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Unable to enqueue the message", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	log.Println("File uploaded successfully to Azure Blob Storage")
	w.Write([]byte("File uploaded successfully to Azure Blob Storage"))
}
