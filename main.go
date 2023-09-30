package main

import (
	"log"

	"github.com/Hrit99/zock.git/config"
	"github.com/gin-gonic/gin"

	"go.mongodb.org/mongo-driver/mongo"
)

type product struct {
	Product_id                int      `json:"product_id"`
	Product_name              string   `json:"product_name"`
	Product_description       string   `json:"product_description"`
	Product_images            []string `json:"product_images"`
	Product_price             float64  `json:"product_price"`
	Compressed_product_images []string `json:"compressed_product_images"`
}

type user struct {
	Id        int    `json:"user_id"`
	Name      string `json:"name"`
	Mobile    string `json:"mobile"`
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
}

var db *mongo.Database

func main() {
	//load env variables
	err := config.Loadenv()
	if err != nil {
		log.Fatalf("Unable to load env variables. Err: %s", err)
	}

	//mongodb connection
	db, err = ConnectDb()
	if err != nil {
		log.Fatalf("Unable to connect to database. Err: %s", err)
	}

	//gin server router and handlers
	router := gin.Default()
	router.POST("/product", PostProduct)
	router.Run("localhost:" + config.Port)
}
