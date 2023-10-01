package main

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

func Test_PostProduct(t *testing.T) {
	r := SetUpRouter()
	r.POST("/product", PostProduct)
	newproduct := product{
		Product_id:                0,
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

	//delete created document while testing
	defer func() {
		if w.Code == 201 {
			data, err := io.ReadAll(w.Result().Body)
			if err != nil {
				t.Errorf("Error occured while reading the result %v", err)
			}
			var resProduct product
			err = json.Unmarshal(data, &resProduct)
			if err != nil {
				t.Errorf("Error occured while unmarshaling the result %v", err)
			}
			db.Collection("products").DeleteOne(context.TODO(), bson.D{{"product_id", resProduct.Product_id}})
		}

	}()
}

// func Benchmark_PostProduct(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		r := SetUpRouter()
// 		r.POST("/product", PostProduct)
// 		newproduct := product{
// 			Product_id:                0,
// 			Product_name:              "lakme",
// 			Product_description:       "a description for product",
// 			Product_images:            []string{"link1", "link2"},
// 			Product_price:             10.2,
// 			Compressed_product_images: []string{"link1", "link2"},
// 		}
// 		jsonValue, _ := json.Marshal(newproduct)
// 		req, _ := http.NewRequest("POST", "/product", bytes.NewBuffer(jsonValue))
// 		w := httptest.NewRecorder()
// 		r.ServeHTTP(w, req)
// 		assert.Equal(b, http.StatusCreated, w.Code)

// 		//delete created document while testing
// 		defer func() {
// 			if w.Code == 201 {
// 				data, err := io.ReadAll(w.Result().Body)
// 				if err != nil {
// 					b.Errorf("Error occured while reading the result %v", err)
// 				}
// 				var resProduct product
// 				err = json.Unmarshal(data, &resProduct)
// 				if err != nil {
// 					b.Errorf("Error occured while unmarshaling the result %v", err)
// 				}
// 				db.Collection("products").DeleteOne(context.TODO(), bson.D{{"product_id", resProduct.Product_id}})
// 			}
// 		}()
// 	}
// }
