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

	//initialize product_id value
	product_id, err := db.Collection("products").CountDocuments(context.TODO(), bson.D{})
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}
	_, err = db.Collection("products").UpdateByID(context.TODO(), doc.InsertedID, bson.D{{"$set", bson.D{{"product_id", product_id}}}})
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	//send added product as response
	c.IndentedJSON(http.StatusCreated, newProduct)
}
