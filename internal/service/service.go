package service

import (
	"context"

	"github.com/JoelBautista1093/ANproyecto/internal/models"
	"github.com/JoelBautista1093/ANproyecto/internal/repository"
)

// Services is the bussines logic of the aplications
//
//go:generate mockery --name=Service --output=service --inpackage
type Service interface {
	RegisterUser(ctx context.Context, email, name, password string) error
	LoginUser(ctx context.Context, email, password string) (*models.User, error)
}

type serv struct {
	repo repository.Repository
}

func New(repo repository.Repository) Service {
	return &serv{
		repo: repo,
	}
}
