package main

import (
	"fmt"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"group.id":          "group-id",
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		log.Printf("Failed to create Consumer: %s\n", err)
		return
	}
	defer consumer.Close()

	topics := []string{"This Afternoon", "Tonight"}
	for i := 0; i < len(topics); i++ {
		err = consumer.Subscribe(topics[i], nil)
		if err != nil {
			log.Printf("Failed to subscribe the topic: %s\n", err)
			return
		}
	}

	for {
		msg, err := consumer.ReadMessage(-1) // time duration
		if err == nil {
			fmt.Printf("Received message: %s\n", string(msg.Value))
		} else {
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}
}
