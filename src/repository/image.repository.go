package repository

import (
	"context"
	"log"
	"os"

	"github.com/PetrusAriaa/web-margot-backend/src/db"
	"google.golang.org/api/iterator"
)

type Repository struct {
	db  *db.DBConnections
	ctx context.Context
}

func NewRepository(db *db.DBConnections, ctx context.Context) *Repository {
	return &Repository{
		db,
		ctx,
	}
}

func (r *Repository) GetImageList() []string {
	bkt := r.db.CloudStorageClient.Bucket(os.Getenv("GCLOUD_BUCKET"))
	obj := bkt.Objects(r.ctx, nil)

	var blobList []string

	for {
		attr, err := obj.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatal(err.Error())
		}
		blobList = append(blobList, attr.Name)
	}
	return blobList
}
