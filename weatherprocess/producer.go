package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	weathersubs "github.com/saanvijay/weathernotify/weathersubs"
)

type WeatherForecast struct {
	Name          string `json:"name"`
	Temperature   int    `json:"temperature"`
	WindSpeed     string `json:"windSpeed"`
	ShortForecast string `json:"shortForecast"`
}

func kafkaProduceForcast() {
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		// User-specific properties that you must set
		"bootstrap.servers": "kafka1:19092",

		// Fixed properties
		"acks": "all"})

	if err != nil {
		fmt.Printf("Failed to create producer: %s", err)
		os.Exit(1)
	}

	// Go-routine to handle message delivery reports and
	// possibly other event types (errors, stats, etc)
	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Failed to deliver message: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("Produced event to topic %s: key = %-10s value = %s\n",
						*ev.TopicPartition.Topic, string(ev.Key), string(ev.Value))
				}
			}
		}
	}()

	// Set up a ticker to produce messages every EMAIL_NOTIFICATION_IN_MINUTES minutes
	minutes := os.Getenv("EMAIL_NOTIFICATION_IN_MINUTES")
	minutesInt, err := strconv.Atoi(minutes)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	timeInterval := time.Duration(minutesInt) * time.Minute
	ticker := time.NewTicker(timeInterval)

	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
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
			weatherData := string(jsonData)
			topic := "weather"

			err = p.Produce(&kafka.Message{
				TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
				Key:            []byte("forcast"),
				Value:          []byte(weatherData),
			}, nil)
			if err != nil {
				fmt.Printf("Failed to produce weather message: %s\n", err)
			} else {
				fmt.Println("Weather message queued for delivery")
			}
		}

		// Wait for all messages to be delivered
		p.Flush(15 * 1000)

	}

	//p.Close()
}
