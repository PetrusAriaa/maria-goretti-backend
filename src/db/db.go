package db

import (
	"context"
	"log"

	"cloud.google.com/go/storage"
)

type DBConnections struct {
	CloudStorageClient *storage.Client
}

func NewDBConnections(ctx context.Context) *DBConnections {

	s, err := storage.NewClient(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}
	return &DBConnections{
		CloudStorageClient: s,
	}
}
