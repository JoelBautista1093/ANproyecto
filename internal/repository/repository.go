package repository

import (
	"context"

	"github.com/JoelBautista1093/ANproyecto/internal/entity"
	"github.com/jmoiron/sqlx"
)

// Repository is the interface that wraps the basic CRUD operations.
//
// run go generate ./... |run go generate
//
//go:generate mockery --name=Repository --output=repository --inpackage
type Repository interface { //You, 9 hours ago * register and login _
	SaveUser(ctx context.Context, email, name, password string) error
	GetUserByEmail(ctx context.Context, email string) (*entity.User, error)
}

type repo struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) Repository {
	return &repo{
		db: db,
	}
}
