package consumer

import (
	// Change With Your Package Name

	"context"
	"errors"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	kafkaConfig "github.com/Hrit99/zock.git/config"
	database "github.com/Hrit99/zock.git/db"
	"github.com/IBM/sarama"
	"github.com/h2non/bimg"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"log"
)

type product struct {
	Product_id                int      `json:"product_id"`
	Product_name              string   `json:"product_name"`
	Product_description       string   `json:"product_description"`
	Product_images            []string `json:"product_images"`
	Product_price             float64  `json:"product_price"`
	Compressed_product_images []string `json:"compressed_product_images"`
}

func Consume(topic string) {

	//intialize consumer
	config := sarama.NewConfig()

	consumer, err := sarama.NewConsumer([]string{kafkaConfig.CONST_HOST}, config)
	if err != nil {
		log.Fatal("NewConsumer err: ", err)
	}
	defer consumer.Close()

	partitionConsumer, err := consumer.ConsumePartition(topic, 0, sarama.OffsetNewest)
	if err != nil {
		log.Fatal("ConsumePartition err: ", err)
	}
	defer partitionConsumer.Close()

	//initializing mongo database client
	db, err := initializeDbForConsumer()
	if err != nil {
		log.Fatal("Error occured while connecting to MongoDB: ", err)
	}

	//disconnect mongo-client socket after fetching product.
	defer func() {
		db.Client().Disconnect(context.TODO())
	}()

	for message := range partitionConsumer.Messages() {
		log.Printf("product_id passed value: %s\n", string(message.Value))
		intVar, _ := strconv.Atoi(string(message.Value))
		prod := GetProduct(intVar, db)

		//download every image of the product
		for i, v := range prod.Product_images {
			filePathName := "saved_images/" + prod.Product_name + "_" + strconv.Itoa(prod.Product_id) + "/" + strconv.Itoa(i) + ".webp"
			URL := v
			err := downloadFile(URL, filePathName)
			if err != nil {
				log.Fatal(err)
			}
			prod.Compressed_product_images = append(prod.Compressed_product_images, filePathName)
		}
		//store filePathName to product table in database
		storFileLocalPath(prod.Product_id, prod.Compressed_product_images, db)
	}

}

func initializeDbForConsumer() (*mongo.Database, error) {
	//initializing env variables for consumer
	err := godotenv.Load("local.env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
		return nil, err
	}

	//connect db for consume
	db, err := database.ConnectDb(os.Getenv("MONGO_URI"))
	if err != nil {
		log.Fatal("db connection err: ", err)
		return nil, err
	}
	return db, nil
}

// module to update local paths in database
func storFileLocalPath(product_id int, paths []string, db *mongo.Database) {
	filter := bson.D{primitive.E{Key: "product_id", Value: product_id}}

	update := bson.D{primitive.E{Key: "$set", Value: bson.D{
		primitive.E{Key: "compressed_product_images", Value: paths},
	}}}

	db.Collection("products").FindOneAndUpdate(context.TODO(), filter, update)
}

func GetProduct(product_id int, db *mongo.Database) *product {

	//get docuemnt using product_id
	var responseProduct product
	err := db.Collection("products").FindOne(context.TODO(), bson.D{{"product_id", product_id}}).Decode(&responseProduct)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// This error means your query did not match any documents.
			log.Fatal("No documents matched in the collection ", err)
			return nil
		}
		log.Fatal("Some error occured ", err)
		return nil
	}
	return &responseProduct //return product
}

func downloadFile(URL, filePath string) error {
	//Get the response bytes from the url
	response, err := http.Get(URL)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return errors.New("Received non 200 response code")
	}
	//Create a empty file with directories
	file, err := createFile(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	//compress image
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}
	compressedImg, err := compressImage(body, 40)
	if err != nil {
		return err
	}

	//Write the bytes to the file
	_, err = file.Write(compressedImg)
	if err != nil {
		return err
	}

	return nil
}

func createFile(p string) (*os.File, error) {
	if err := os.MkdirAll(filepath.Dir(p), 0770); err != nil {
		return nil, err
	}
	return os.Create(p)
}

func compressImage(buffer []byte, quality int) ([]byte, error) {

	converted, err := bimg.NewImage(buffer).Convert(bimg.WEBP)
	if err != nil {
		return nil, err
	}

	processed, err := bimg.NewImage(converted).Process(bimg.Options{Quality: quality})
	if err != nil {
		return nil, err
	}
	return processed, nil
}
