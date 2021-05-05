package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/apache/pulsar-client-go/pulsar"
)

type Product struct {
	Id int
	Name string
}

func main() {
    client, err := pulsar.NewClient(pulsar.ClientOptions{
        URL:               "pulsar://localhost:6650",
        OperationTimeout:  30 * time.Second,
        ConnectionTimeout: 30 * time.Second,
    })
    if err != nil {
        log.Fatalf("Could not instantiate Pulsar client: %v", err)
    }

    defer client.Close()

		producer, err := client.CreateProducer(pulsar.ProducerOptions{
			Topic: "inventory-service",
		})
	
		data, err := json.Marshal(&Product{
			Id: 1, 
			Name: "Pulsar rocks!",
		})

		if err != nil {
			fmt.Println("Unable to marshal json: ", err)
		}

		_, err = producer.Send(context.Background(), &pulsar.ProducerMessage{
			Payload: data,
		})
		
		if err != nil {
			fmt.Println("Failed to publish message", err)
		}
		
		fmt.Println("Published message")

		defer producer.Close()
	}
