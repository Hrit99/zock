package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
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

func main() {
	//env initialization
	err := godotenv.Load("local.env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	port := os.Getenv("PORT")

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
	c.IndentedJSON(http.StatusCreated, newProduct)
}
