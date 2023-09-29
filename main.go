package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type product struct {
	Product_id                int      `json:"product_id"`
	Product_name              string   `json:"product_name"`
	Product_description       string   `json:"product_description"`
	Product_images            []string `json:"product_images"`
	Product_price             float64  `json:"product_price"`
	Compressed_product_images []string `json:"compressed_product_images"`
	Created_at                string   `json:"created_at"`
	Updated_at                string   `json:"updated_at"`
}

type user struct {
	Id         int    `json:"_id"`
	Name       string `json:"name"`
	Mobile     string `json:"mobile"`
	Latitude   string `json:"latitude"`
	Longitude  string `json:"longitude"`
	Created_at string `json:"created_at"`
	Updated_at string `json:"updated_at"`
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
	db, err = connectDb(uri)
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	//gin server router and handlers
	router := gin.Default()
	router.POST("/product", postProduct)
	router.Run("localhost:" + port)
}

func postProduct(c *gin.Context) {
	var newProduct product

	if err := c.BindJSON(&newProduct); err != nil {
		return
	}

	//add new product to db
	db.Collection("products").InsertOne(context.TODO(), newProduct)
	c.IndentedJSON(http.StatusCreated, newProduct)
}

func connectDb(uri string) (*mongo.Database, error) {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)
	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		return nil, err
	}

	// defer func() {
	// 	fmt.Println("disconnecting")
	// 	if err = client.Disconnect(context.TODO()); err != nil {
	// 		return
	// 	}
	// }()

	// Send a ping to confirm a successful connection
	if err := client.Database("zock").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err(); err != nil {
		return nil, err
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")
	db := client.Database("zock")
	return db, nil
}
