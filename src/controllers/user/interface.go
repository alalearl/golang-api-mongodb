package user

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type Handler struct {
	DB *mongo.Database
}

type User struct {
	Name string `db:"name" json:"name"`
	Age  int    `db:"age" json:"age"`
	City string `db:"city" json:"city"`
}
