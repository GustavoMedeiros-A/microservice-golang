package main

// This package amqp091-go is using because of replace a community developed package
// and have a lot of people that use that

import (
	"fmt"
	"listener/event"
	"log"
	"math"
	"os"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	// Try to connect to rabbitMQ
	rabbitConn, err := connect()

	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	defer rabbitConn.Close()
	fmt.Println("We connect to rabbitMQ my Friend...")

	// Start listening for messages
	log.Println("Listening for and consuming RabbitMQ messages...")

	// Create a consumer
	consumer, err := event.NewConsumer(rabbitConn)
	if err != nil {
		panic(err)
	}

	// Watch the queue and consume events (call listeners)
	err = consumer.Listen([]string{"log.INFO", "log.AUTH", "log.WARNING", "log.ERROR"})
	if err != nil {
		log.Println(err, "is a error")
	}
}

func connect() (*amqp.Connection, error) {

	var counts int64
	var backOff = 1 * time.Second
	var connection *amqp.Connection

	// don't continue until rabbit is ready

	for {
		connect, err := amqp.Dial("amqp://rabbitmq:rabbitmq@rabbitmq")
		if err != nil {
			fmt.Println("RabbitMQ not yet ready...")
			counts++
		} else {
			connection = connect
			break
		}
		if counts > 5 {
			fmt.Println(err)
			return nil, err
		}
		// Increase the delay
		backOff = time.Duration(math.Pow(float64(counts), 2)) * time.Second
		fmt.Println("Backing off...")
		time.Sleep(backOff)
		continue
	}

	return connection, nil

}
