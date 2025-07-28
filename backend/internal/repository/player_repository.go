package repository

import "go.mongodb.org/mongo-driver/mongo"

type (
	IPlayerRepository interface {
	}

	PlayerRepository struct {
		db         *mongo.Database
		collection *mongo.Collection
	}
)

func NewPlayerRepository(repo *Repository) *PlayerRepository {
	return &PlayerRepository{
		db:         repo.DB,
		collection: repo.DB.Collection("players"),
	}
}
