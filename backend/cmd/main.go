package main

import (
	"go-mongo/internal/handler"
	"go-mongo/internal/infrastructure"
	"go-mongo/internal/repository"
	"go-mongo/internal/services"

	"github.com/labstack/echo/v4"
)


var (
	cfg *infrastructure.Config
	connection *repository.Repository
	service *services.Service
	handlers *handler.Handler
)

func init() {
	cfg = infrastructure.NewConfig()
	connection = repository.NewRepository(cfg)
	service = services.NewService(connection)
	handlers = handler.NewHandler(service)
}

func main() {
	e := echo.New()



	e.Logger.Fatal(e.Start(":" + "8080"))
}

