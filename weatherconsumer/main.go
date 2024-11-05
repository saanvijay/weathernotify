package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":        "localhost:9092",
		"broker.address.family":    "v4",
		"session.timeout.ms":       6000,
		"group.id":                 "vijay-v3-group", // Consumer group ID
		"auto.offset.reset":        "earliest",       // Start reading from the earliest message
		"enable.auto.offset.store": false,
	})
	if err != nil {
		log.Printf("Failed to create Consumer: %s\n", err)
		return
	}
	defer consumer.Close()

	topics := []string{"Afternoon", "Tonight"}
	//for i := 0; i < len(topics); i++ {
	err = consumer.SubscribeTopics(topics, nil)
	if err != nil {
		log.Printf("Failed to subscribe the topic: %s\n", err)
		return
	}
	//}

	run := true

	for run {
		select {
		case sig := <-sigchan:
			fmt.Printf("Caught signal %v: terminating\n", sig)
			run = false
		default:
			ev := consumer.Poll(100)
			if ev == nil {
				continue
			}

			switch e := ev.(type) {
			case *kafka.Message:
				// Process the message received.
				fmt.Printf("%% Message on %s:\n%s\n",
					e.TopicPartition, string(e.Value))
				if e.Headers != nil {
					fmt.Printf("%% Headers: %v\n", e.Headers)
				}

				_, err := consumer.StoreMessage(e)
				if err != nil {
					fmt.Fprintf(os.Stderr, "%% Error storing offset after message %s:\n",
						e.TopicPartition)
				}
			case kafka.Error:
				// Errors should generally be considered
				// informational, the client will try to
				// automatically recover.
				// But in this example we choose to terminate
				// the application if all brokers are down.
				fmt.Fprintf(os.Stderr, "%% Error: %v: %v\n", e.Code(), e)
				if e.Code() == kafka.ErrAllBrokersDown {
					run = false
				}
			default:
				fmt.Printf("Ignored %v\n", e)
			}
		}
	}

	fmt.Printf("Closing consumer\n")
	consumer.Close()
}
