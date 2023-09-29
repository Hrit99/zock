package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"go.mongodb.org/mongo-driver/mongo"
)

type product struct {
	Product_id                int      `json:"_id"`
	Product_name              string   `json:"product_name"`
	Product_description       string   `json:"product_description"`
	Product_images            []string `json:"product_images"`
	Product_price             float64  `json:"product_price"`
	Compressed_product_images []string `json:"compressed_product_images"`
}

type user struct {
	Id        int    `json:"_id"`
	Name      string `json:"name"`
	Mobile    string `json:"mobile"`
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
}

var db *mongo.Database

func main() {
	//env initialization
	err := godotenv.Load("local.env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	port := os.Getenv("PORT")
	uri := os.Getenv("MONGO_URI")

	//mongodb connection
	db, err = ConnectDb(uri)
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	//gin server router and handlers
	router := gin.Default()
	router.POST("/product", PostProduct)
	router.Run("localhost:" + port)
}
