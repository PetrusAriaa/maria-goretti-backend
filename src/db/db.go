package db

import (
	"context"
	"encoding/base64"
	"log"
	"os"

	"cloud.google.com/go/storage"
	"google.golang.org/api/option"
)

type DBConnections struct {
	CloudStorageClient *storage.Client
	StorageBucket      *storage.BucketHandle
}

func NewDBConnections(ctx context.Context) *DBConnections {
	b64, err := base64.StdEncoding.DecodeString(os.Getenv("GOOGLE_CREDENTIALS_BASE64"))
	if err != nil {
		panic(err.Error())
	}

	opt := option.WithCredentialsJSON(b64)
	s, err := storage.NewClient(ctx, opt)
	if err != nil {
		log.Fatal(err.Error())
	}
	bkt := s.Bucket(os.Getenv("GCLOUD_BUCKET"))
	return &DBConnections{
		CloudStorageClient: s,
		StorageBucket:      bkt,
	}
}
