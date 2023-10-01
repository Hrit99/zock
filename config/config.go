package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

const (
	CONST_HOST  = "localhost:9092"
	CONST_TOPIC = "zock-topic"
)

var Port string
var Uri string

func Loadenv() error {
	//env initialization
	err := godotenv.Load("local.env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
		return err
	}

	Port = os.Getenv("PORT")
	Uri = os.Getenv("MONGO_URI")
	return nil
}
