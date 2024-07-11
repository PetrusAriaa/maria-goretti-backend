package service

import (
	"context"

	"github.com/PetrusAriaa/web-margot-backend/lib/db"
	"github.com/PetrusAriaa/web-margot-backend/lib/repository"
)

type Service struct {
	repository *repository.Repository
}

func NewService(db *db.DBConnections, ctx context.Context) *Service {
	r := repository.NewRepository(db, ctx)
	return &Service{
		r,
	}
}
