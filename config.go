package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var port string
var uri string

func Loadenv() error {
	//env initialization
	err := godotenv.Load("local.env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
		return err
	}

	port = os.Getenv("PORT")
	uri = os.Getenv("MONGO_URI")
	return nil
}
