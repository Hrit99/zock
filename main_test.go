package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
	"github.com/stretchr/testify/assert"
)

func SetUpRouter() *gin.Engine {
	//load env variables
	err := Loadenv()
	if err != nil {
		log.Fatalf("Unable to load env variables. Err: %s", err)
	}

	//mongodb connection
	db, err = ConnectDb(uri)
	if err != nil {
		log.Fatalf("Unable to connect to database. Err: %s", err)
	}

	router := gin.Default()
	return router
}

func Test_PostProduct(t *testing.T) {
	r := SetUpRouter()
	r.POST("/product", PostProduct)
	productId := xid.New().Pid()
	newproduct := product{
		Product_id:                int(productId),
		Product_name:              "lakme",
		Product_description:       "a description for product",
		Product_images:            []string{"link1", "link2"},
		Product_price:             10.2,
		Compressed_product_images: []string{"link1", "link2"},
	}
	jsonValue, _ := json.Marshal(newproduct)
	req, _ := http.NewRequest("POST", "/product", bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)
}
