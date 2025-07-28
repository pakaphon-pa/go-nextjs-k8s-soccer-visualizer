package repository

import (
	"context"
	"fmt"
	"go-mongo/internal/infrastructure"
	"time"

	"github.com/labstack/gommon/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type (
	Repository struct {
		Client *mongo.Client
		DB     *mongo.Database
	}
)

func NewRepository(config *infrastructure.Config) *Repository {
	fmt.Println("Database connection...")

	ctx := context.Background()
	mongoURI := fmt.Sprintf("mongodb://%s:%s@%s:%s", config.Database.Username, config.Database.Password, config.Database.Host, config.Database.Port)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))

	if err != nil {
		errCount := 0
		for i := 1; i <= 3; i++ {
			log.Info("Try To Connect database... [%d]", i)
			time.Sleep(6 * time.Second)
			client, err = mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
			if err != nil {
				errCount++
				continue
			}
			break
		}
		if errCount == 3 {
			panic(err)
		}
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		panic(err)
	}

	fmt.Println("Database connection Successfully")

	return &Repository{
		Client: client,
		DB:     client.Database("transfermarket"),
	}
}
