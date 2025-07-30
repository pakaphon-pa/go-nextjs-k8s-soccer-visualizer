package main

import (
	"context"
	"fmt"
	"go-mongo/internal/handler"
	"go-mongo/internal/infrastructure"
	"go-mongo/internal/repository"
	"go-mongo/internal/services"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
)

var (
	cfg        *infrastructure.Config
	connection *repository.Repository
	service    *services.Service
	handlers   *handler.Handler
)

func init() {
	cfg = infrastructure.NewConfig()
	connection = repository.NewRepository(cfg)
	service = services.NewService(connection)
	handlers = handler.NewHandler(service)
}

func main() {
	e := echo.New()

	version := e.Group("/v1")
	version.GET("/healthz", func(c echo.Context) error {
		return c.NoContent(http.StatusOK)
	})

	go func() {
		if err := e.Start(":" + cfg.Server.Port); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("Shutting down the server: ", err)

		}
	}()

	quite := make(chan os.Signal, 1)
	signal.Notify(quite, syscall.SIGINT, syscall.SIGTERM)
	<-quite
	fmt.Println("Shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}

	ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := connection.Client.Disconnect(ctx); err != nil {
		e.Logger.Fatal(err)
	} else {
		fmt.Println("MongoDB Disconnected")
	}

}
