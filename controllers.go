package main

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func PostProduct(c *gin.Context) {
	var newProduct product

	if err := c.BindJSON(&newProduct); err != nil {
		return
	}

	//add new product to db
	doc, err := db.Collection("products").InsertOne(context.TODO(), newProduct)
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	//update product_id value to _id value
	_, err = db.Collection("products").UpdateByID(context.TODO(), doc.InsertedID, bson.D{{"$set", bson.D{{"product_id", doc.InsertedID}}}})
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	//send added product as response
	c.IndentedJSON(http.StatusCreated, newProduct)
}
