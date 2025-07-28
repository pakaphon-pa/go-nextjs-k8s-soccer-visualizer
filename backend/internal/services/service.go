package services

import "go-mongo/internal/repository"

type (
	Service struct {
		playerRepository repository.IPlayerRepository
	}
)

func NewService(repo *repository.Repository) *Service {
	return &Service{
		playerRepository: repository.NewPlayerRepository(repo),
	}
}
