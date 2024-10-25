package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	weathersubs "github.com/saanvijay/weathernotify/weathersubs"
)

func kafkaProduceForcast() {
	producer, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": ":9092"})
	if err != nil {
		log.Printf("Failed to create kafka producer: %s\n", err)
	}
	defer producer.Close()

	// Event Listener
	// Listen to all the events on the default events channel
	go func() {
		for e := range producer.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				// The message delivery report, indicating success or
				// permanent failure after retries have been exhausted.
				// Application level retries won't help since the client
				// is already configured to do that.
				m := ev
				if m.TopicPartition.Error != nil {
					fmt.Printf("Delivery failed: %v\n", m.TopicPartition.Error)
				} else {
					fmt.Printf("Delivered message to topic %s [%d] at offset %v\n",
						*m.TopicPartition.Topic, m.TopicPartition.Partition, m.TopicPartition.Offset)
				}
			case kafka.Error:
				// Generic client instance-level errors, such as
				// broker connection failures, authentication issues, etc.
				//
				// These errors should generally be considered informational
				// as the underlying client will automatically try to
				// recover from any errors encountered, the application
				// does not need to take action on them.
				fmt.Printf("Error: %v\n", ev)
			default:
				fmt.Printf("Ignored event: %s\n", ev)
			}
		}
	}()

	location, err := weathersubs.GetCurrentLocation()
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	forecast, err := weathersubs.GetForeCast(location)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	jsonData, _ := json.MarshalIndent(forecast.Properties.Periods, "", " ")

	topics := []string{"This Afternoon", "Tonight"}

	for i := 0; i < len(topics); i++ {
		topic := topics[i]
		message := string(jsonData[i])
		err1 := producer.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{
				Topic:     &topic,
				Partition: kafka.PartitionAny},
			Value: []byte(message),
		}, nil)
		if err1 != nil {
			log.Printf("Failed to produce message: %s\n", err1)
		}
	}

	producer.Flush(15 * 1000)

}
