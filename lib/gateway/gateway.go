package gateway

import (
	"context"
	"errors"

	"github.com/PetrusAriaa/web-margot-backend/lib/db"
	"github.com/PetrusAriaa/web-margot-backend/lib/service"
)

type Gateway struct {
	event string
}

func NewGateway(e string) *Gateway {
	return &Gateway{
		event: e,
	}
}

func (g *Gateway) Call() (interface{}, error) {
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
	default:
		return nil, errors.New("Unknown event: " + ev)
	}
}
