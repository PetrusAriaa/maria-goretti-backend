package repository

import (
	"log"
	"os"

	"google.golang.org/api/iterator"
)

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
