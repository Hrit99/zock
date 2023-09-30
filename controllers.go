package main

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/Hrit99/zock.git/producer"
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
	newProduct.Product_id = int(product_id)
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}
	_, err = db.Collection("products").UpdateByID(context.TODO(), doc.InsertedID, bson.D{{"$set", bson.D{{"product_id", product_id}}}})
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	//send product_id to kafka queue
	producer.Produce(int(product_id))

	//send added product as response
	c.IndentedJSON(http.StatusCreated, newProduct)
}
