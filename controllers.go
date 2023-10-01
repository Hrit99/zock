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
	var newreq postreq

	if err := c.BindJSON(&newreq); err != nil {
		return
	}

	//get unique product_id
	product_id, err := db.Collection("products").CountDocuments(context.TODO(), bson.D{})
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	//add new product to db
	_, err = db.Collection("products").InsertOne(context.TODO(), &product{
		Product_id:                int(product_id) + 1,
		Product_name:              newreq.Product_name,
		Product_description:       newreq.Product_description,
		Product_images:            newreq.Product_images,
		Product_price:             newreq.Product_price,
		Compressed_product_images: []string{},
	})
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	//add new user to db
	_, err = db.Collection("users").InsertOne(context.TODO(), &user{
		Id: newreq.Id,
	})
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	//send product_id to kafka queue
	producer.Produce(int(product_id))

	//send added product as response
	c.IndentedJSON(http.StatusCreated, &product{
		Product_id:                int(product_id) + 1,
		Product_name:              newreq.Product_name,
		Product_description:       newreq.Product_description,
		Product_images:            newreq.Product_images,
		Product_price:             newreq.Product_price,
		Compressed_product_images: []string{},
	})
}
