package handler

import "go-mongo/internal/services"



type (
	Handler struct {

	}
)


func NewHandler(svc *services.Service) *Handler {
	return &Handler{}
} 