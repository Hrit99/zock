package main

import (
	"log"

	"github.com/gin-gonic/gin"
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
