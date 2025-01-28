package azstorage

import (
	"context"
	"errors"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
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
		var responseError *azcore.ResponseError
		if errors.As(err, &responseError) {
			if responseError.StatusCode != 409 {
				return nil, errors.New(fmt.Sprintf("could not create queue %s: %v", queueName, responseError.Error()))
			}else{
				return &Queue{
					client: qc,
				}, nil
			}
		}
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
