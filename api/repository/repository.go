package repository

import (
	"context"

	"github.com/PetrusAriaa/web-margot-backend/api/db"
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
