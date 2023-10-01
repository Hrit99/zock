package main

import (
	"log"

	"github.com/Hrit99/zock.git/config"
	database "github.com/Hrit99/zock.git/db"
	"github.com/gin-gonic/gin"
)

func SetUpRouter() *gin.Engine {
	//load env variables
	err := config.Loadenv()
	if err != nil {
		log.Fatalf("Unable to load env variables. Err: %s", err)
	}

	//mongodb connection
	db, err = database.ConnectDb(config.Uri)
	if err != nil {
		log.Fatalf("Unable to connect to database. Err: %s", err)
	}

	router := gin.Default()
	return router
}
