package user

import (
	"context"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (h Handler) Lists(c *gin.Context) {
	collection := h.DB.Collection("users")
	findOptions := options.Find()
	findOptions.SetLimit(10)
	var results []*User
	cur, err := collection.Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		log.Fatal(err)
	}
	for cur.Next(context.TODO()) {
		var elem User
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, &elem)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	cur.Close(context.TODO())
	c.JSON(200, results)
}

func (h Handler) CreateNew(c *gin.Context) {
	name := c.PostForm("name")
	age, _ := strconv.Atoi(c.PostForm("age"))
	city := c.PostForm("city")
	collection := h.DB.Collection("users")
	addUserNew := User{
		Name: name,
		Age:  age,
		City: city,
	}
	insertResult, err := collection.InsertOne(context.TODO(), addUserNew)
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(200, insertResult.InsertedID)
}
