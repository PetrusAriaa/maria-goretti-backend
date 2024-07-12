package gateway

import (
	"context"
	"errors"

	"github.com/PetrusAriaa/web-margot-backend/lib/db"
	"github.com/PetrusAriaa/web-margot-backend/lib/service"
)

type Gateway struct {
	event   string
	Service *service.Service
}

func NewGateway(e string) (*Gateway, error) {
	ctx := context.Background()
	client, err := db.NewDBConnections(ctx)
	if err != nil {
		return nil, err
	}
	defer client.CloudStorageClient.Close()

	srv := service.NewService(client, ctx)

	return &Gateway{
		event:   e,
		Service: srv,
	}, nil
}

func (g *Gateway) Call(data ...string) (interface{}, error) {
	ctx := context.Background()
	client, err := db.NewDBConnections(ctx)
	if err != nil {
		return nil, err
	}
	defer client.CloudStorageClient.Close()

	srv := service.NewService(client, ctx)

	ev := g.event
	switch ev {
	case "images":
		imgs, err := srv.GetImageList()
		return imgs, err
	case "images:get":
		img, attr, err := srv.GetImage(data[0])
		if err != nil {
			return nil, err
		}
		type ImgData struct {
			content     []byte
			contentType string
		}
		imd := &ImgData{
			img,
			attr,
		}
		return imd, nil
	default:
		return nil, errors.New("Unknown event: " + ev)
	}
}
