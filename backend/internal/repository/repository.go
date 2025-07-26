package repository

import (
	"go-mongo/internal/infrastructure"

	"go.mongodb.org/mongo-driver/mongo"
)

type (
	Repository struct {
		db *mongo.Client
	}
)

func NewRepository(config *infrastructure.Config) *Repository {


	return &Repository{
		db: &mongo.Client{},
	}
}