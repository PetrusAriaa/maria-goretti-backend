package repository

import (
	"io"
	"log"
	"os"

	"cloud.google.com/go/storage"
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

func (r *Repository) GetImage(img string) ([]byte, string) {
	bkt := r.db.CloudStorageClient.Bucket(os.Getenv("GCLOUD_BUCKET"))

	obj := bkt.Object(img)
	attr, err := obj.Attrs(r.ctx)
	if err != nil {
		log.Default().Fatal(err.Error())
	}

	reader, err := obj.NewReader(r.ctx)
	if err == storage.ErrObjectNotExist {
		return nil, ""
	}
	if err != nil {
		log.Default().Fatal(err.Error())
	}
	defer reader.Close()

	p, err := io.ReadAll(reader)
	if err != nil {
		log.Default().Fatal(err.Error())
	}

	return p, attr.ContentType
}
