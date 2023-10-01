package consumer

import (
	"context"
	"errors"
	"math/rand"
	"os"
	"testing"

	"github.com/Hrit99/zock.git/config"
	database "github.com/Hrit99/zock.git/db"
	"github.com/IBM/sarama"
	"github.com/IBM/sarama/mocks"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
)

func Test_Consume(t *testing.T) {
	consumer := mocks.NewConsumer(t, mocks.NewTestConfig())
	defer func() {
		if err := consumer.Close(); err != nil {
			t.Error(err)
		}
	}()

	consumer.ExpectConsumePartition(config.CONST_TOPIC, 0, sarama.OffsetOldest).YieldMessage(&sarama.ConsumerMessage{Value: []byte("hello world")})
	consumer.ExpectConsumePartition(config.CONST_TOPIC, 0, sarama.OffsetOldest).YieldError(sarama.ErrOutOfBrokers)
	consumer.ExpectConsumePartition(config.CONST_TOPIC, 1, sarama.OffsetOldest).YieldMessage(&sarama.ConsumerMessage{Value: []byte("hello world again")})
	consumer.ExpectConsumePartition("other", 0, mocks.AnyOffset).YieldMessage(&sarama.ConsumerMessage{Value: []byte("hello other")})

	pc_test0, err := consumer.ConsumePartition(config.CONST_TOPIC, 0, sarama.OffsetOldest)
	if err != nil {
		t.Fatal(err)
	}
	test0_msg := <-pc_test0.Messages()
	if test0_msg.Topic != config.CONST_TOPIC || test0_msg.Partition != 0 || string(test0_msg.Value) != "hello world" {
		t.Error("Message was not as expected:", test0_msg)
	}
	test0_err := <-pc_test0.Errors()
	if !errors.Is(test0_err, sarama.ErrOutOfBrokers) {
		t.Error("Expected sarama.ErrOutOfBrokers, found:", test0_err.Err)
	}

	pc_test1, err := consumer.ConsumePartition(config.CONST_TOPIC, 1, sarama.OffsetOldest)
	if err != nil {
		t.Fatal(err)
	}
	test1_msg := <-pc_test1.Messages()
	if test1_msg.Topic != config.CONST_TOPIC || test1_msg.Partition != 1 || string(test1_msg.Value) != "hello world again" {
		t.Error("Message was not as expected:", test1_msg)
	}

	pc_other0, err := consumer.ConsumePartition("other", 0, sarama.OffsetNewest)
	if err != nil {
		t.Fatal(err)
	}
	other0_msg := <-pc_other0.Messages()
	if other0_msg.Topic != "other" || other0_msg.Partition != 0 || string(other0_msg.Value) != "hello other" {
		t.Error("Message was not as expected:", other0_msg)
	}
}

func Test_initializeDbForConsumer(t *testing.T) {
	result := godotenv.Load("../local.env")

	if result != nil {
		t.Errorf("\"Loadenv()\" FAILED, expected -> <nil>, got -> %v", result)
	} else {
		t.Logf("\"Loadenv()\" PASSED, expected -> <nil>, got -> %v", result)
	}

	_, result = database.ConnectDb(os.Getenv("MONGO_URI"))

	if result != nil {
		t.Errorf("\"ConnectDb()\" FAILED, expected -> nil, got -> %v", result)
	} else {
		t.Logf("\"ConnectDb()\" PASSED, expected -> <nil>, got -> %v", result)
	}
}

func Benchmark_initializeDbForConsumer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		godotenv.Load("../local.env")
		database.ConnectDb(os.Getenv("MONGO_URI"))

	}
}

func Test_GetProduct(t *testing.T) {
	godotenv.Load("../local.env")
	db, err := database.ConnectDb(os.Getenv("MONGO_URI"))
	if err != nil {
		t.Errorf("\"ConnectDb()\" FAILED, expected -> nil, got -> %v", err)
	} else {
		t.Logf("\"ConnectDb()\" PASSED, expected -> <nil>, got -> %v", err)
	}
	defer func() {
		db.Client().Disconnect(context.TODO())
	}()
	n, err := db.Collection("products").CountDocuments(context.TODO(), bson.D{})
	if err != nil {
		t.Errorf("\"ConnectDb()\" FAILED, expected -> nil, got -> %v", err)
	} else {
		t.Logf("\"ConnectDb()\" PASSED, expected -> <nil>, got -> %v", err)
	}
	pid := rand.Intn(int(n-1)) + 1
	resp := GetProduct(pid, db)
	if resp.Product_id == pid {
		t.Logf("\"ConnectDb()\" PASSED, expected -> <nil>, got -> %v", nil)
	} else {
		t.Errorf("\"ConnectDb()\" FAILED, expected -> nil, got -> %v", err)
	}

}

func Benchmark_GetProduct(b *testing.B) {
	for i := 0; i < b.N; i++ {
		godotenv.Load("../local.env")
		db, _ := database.ConnectDb(os.Getenv("MONGO_URI"))
		defer func() {
			db.Client().Disconnect(context.TODO())
		}()
		n, _ := db.Collection("products").CountDocuments(context.TODO(), bson.D{})
		pid := rand.Intn(int(n-1)) + 1
		GetProduct(pid, db)

	}
}

func Test_downloadFile(t *testing.T) {
	url := "https://www.gstatic.com/webp/gallery/3.jpg"
	path := "saved_images_test/test.webp"
	err := downloadFile(url, path)
	defer func() {
		os.Remove(path)
		os.Remove("saved_images_test")
	}()
	if err != nil {
		t.Errorf("\"downloadFile()\" FAILED, expected -> nil, got -> %v", err)
	} else {
		t.Logf("\"downloadFile()\" PASSED, expected -> <nil>, got -> %v", err)
	}
}

func Benchmark_downloadFile(b *testing.B) {
	for i := 0; i < b.N; i++ {
		url := "https://www.gstatic.com/webp/gallery/3.jpg"
		path := "saved_images_test/test.webp"
		downloadFile(url, path)
		os.Remove(path)
		os.Remove("saved_images_test")
	}
}

func Test_compressImage(t *testing.T) {
	url := "https://www.gstatic.com/webp/gallery/3.jpg"
	path := "saved_images_test/test.webp"
	err := downloadFile(url, path)
	defer func() {
		os.Remove(path)
		os.Remove("saved_images_test")
	}()
	if err != nil {
		t.Errorf("\"compressImage()\" FAILED, expected -> true, got -> %v", err)
	}
	data, err := os.ReadFile(path)
	if err != nil {
		t.Errorf("\"compressImage()\" FAILED, expected -> true, got -> %v", err)
	}
	cdata, err := compressImage(data, 10)
	if err != nil {
		t.Errorf("\"compressImage()\" FAILED, expected -> true, got -> %v", err)
	}
	if len(data) > len(cdata) {
		t.Logf("\"compressImage()\" FAILED, expected -> true, got -> %v", true)
	} else {
		t.Errorf("\"compressImage()\" FAILED, expected -> true, got -> %v", false)
	}

}

func Benchmark_compressImage(b *testing.B) {
	for i := 0; i < b.N; i++ {
		url := "https://www.gstatic.com/webp/gallery/3.jpg"
		path := "saved_images_test/test.webp"
		downloadFile(url, path)
		data, _ := os.ReadFile(path)
		compressImage(data, 10)
		os.Remove(path)
		os.Remove("saved_images_test")
	}
}
