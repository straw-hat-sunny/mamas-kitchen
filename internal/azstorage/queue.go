package azstorage

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/storage/azqueue"
)

type QueueService struct {
	client *azqueue.ServiceClient
}

func NewQueueService(connectionString string) (*QueueService, error) {
	queueServiceClient, err := azqueue.NewServiceClientFromConnectionString(connectionString, nil)
	if err != nil {
		return nil, err
	}

	return &QueueService{
		client: queueServiceClient,
	}, nil
}

func (qs QueueService) CreateQueue(ctx context.Context, queueName string) (*Queue, error) {
	qc := qs.client.NewQueueClient(queueName)

	_, err := qc.Create(ctx, nil)
	if err != nil {
		return nil, err
	}

	return &Queue{
		client: qc,
	}, nil
}

type Queue struct {
	client *azqueue.QueueClient
}

func (q Queue) EnqueueMessage(ctx context.Context, msg string) error {
	_, err := q.client.EnqueueMessage(ctx, msg, nil)
	return err
}
