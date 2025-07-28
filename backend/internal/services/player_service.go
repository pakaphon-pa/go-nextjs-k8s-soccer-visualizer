package services

type (
	IPlayerService interface {
		GetPlayerByCriteria()
	}

	PlayerService struct {
		*Service
	}
)

func NewPlayerService(service *Service) *PlayerService {
	return &PlayerService{
		Service: service,
	}
}
