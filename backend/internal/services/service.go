package services

import "go-mongo/internal/repository"



type (
	Service struct {

	}

)

func NewService(repository *repository.Repository) *Service {
	return &Service{}
}