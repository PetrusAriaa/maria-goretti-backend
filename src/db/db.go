package db

import (
	"context"
	"log"
	"os"

	"cloud.google.com/go/storage"
)

type DBConnections struct {
	CloudStorageClient *storage.BucketHandle
}

func NewDBConnections(ctx context.Context) *DBConnections {

	s, err := storage.NewClient(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}
	bkt := s.Bucket(os.Getenv("GCLOUD_BUCKET"))
	return &DBConnections{
		CloudStorageClient: bkt,
	}
}
