run:
	sudo docker-compose -f zk-kafka.yml up -d
	go run consumer/cmd/main.go

stop: 
	sudo docker-compose -f zk-kafka.yml stop
	sudo docker-compose -f zk-kafka.yml down