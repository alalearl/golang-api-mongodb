package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jay-supakorn/api-service/functions/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ConfigDB struct {
	User     string
	Password string
	Host     string
	Port     string
}

func ConnectDB() *mongo.Database {
	cfg := ConfigDB{
		User:     os.Getenv("mongodb_user"),
		Password: os.Getenv("mongodb_paswd"),
		Host:     os.Getenv("mongodb_host"),
		Port:     os.Getenv("mongodb_port"),
	}
	clientOptions := options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s@%s:%s", cfg.User, cfg.Password, cfg.Host, cfg.Port))

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	logger.Info.Println("Connected to MongoDB!")

	return client.Database("golang_api")
}
