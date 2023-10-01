# zock

A backend designed to listen to at an endpoint. It stores Product and User data on MongoDB backend. The repository has a Consumer application also to listen to the Producer application messages through a message queue implemented using Kafka. The consumer application downloads, compresses and stores the images of a product whose product_id is passed to it. 

## Usage

Copy local.env file to the root directory

### to run consumer application
make run   

### to run producer application and API endpoint
go run .
