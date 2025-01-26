package azstorage

import (
	"context"
	"io"

	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
)

type BlobService struct {
	client *azblob.Client
}

func NewBlobService(connectionString string) (*BlobService, error) {
	client, err := azblob.NewClientFromConnectionString(connectionString, nil)
	if err != nil {
		return nil, err
	}
	return &BlobService{
		client: client,
	}, nil
}

func (bs BlobService) CreateBlobContainer(ctx context.Context, containerName string) (*Blob, error) {
	_, err := bs.client.CreateContainer(ctx, containerName, nil)
	if err != nil {
		return nil, err
	}

	return &Blob{
		containerName: containerName,
		client:        bs.client,
	}, nil

}

type Blob struct {
	containerName string
	client        *azblob.Client
}

func (b Blob) Upload(ctx context.Context, fileName string, data io.Reader) error {
	_, err := b.client.UploadStream(ctx, b.containerName, fileName, data, nil)
	return err
}
