package event

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Consumer struct {
	conn      *amqp.Connection
	queueName string
}

func NewConsumer(conn *amqp.Connection) (Consumer, error) {
	consumer := Consumer{
		conn: conn,
	}

	err := consumer.setUp()

	if err != nil {
		return Consumer{}, err
	}

	return consumer, nil
}

func (consumer *Consumer) setUp() error {
	channel, err := consumer.conn.Channel()
	if err != nil {
		return err
	}

	return declareExchange(channel)
}

type Payload struct {
	Name string   `json:"name"`
	Data []string `json:"data"`
}

func (consumer *Consumer) Listen(topics []string) error {
	ch, err := consumer.conn.Channel()
	if err != nil {
		return err
	}

	defer ch.Close()

	queue, err := declareRandomQueue(ch)

	if err != nil {
		return err
	}

	for _, str := range topics {
		ch.QueueBind(
			queue.Name,
			str,
			"logs_topic",
			false, // Not wait
			nil,
		)

		if err != nil {
			return err
		}
	}

	messages, err := ch.Consume(queue.Name, "", true, false, false, false, nil)
	if err != nil {
		return err
	}

	forever := make(chan bool)
	// run in background
	go func() {
		for d := range messages {
			var payload Payload
			_ = json.Unmarshal(d.Body, &payload) // JSON-encoded data and stores the result in the value pointed to by v

			go handlePayload(payload)
		}
	}()

	fmt.Printf("Waiting for message [Exchange, Queue] [logs_topic, %s]\n", queue.Name)
	<-forever
	return nil
}

func handlePayload(payload Payload) {

	switch payload.Name {
	case "log":
		// Log whatever we get
		err := logEvent(payload)
		if err != nil {
			log.Println(err)
		}

	case "auth":
		// to auth
		err := authEvent(payload)
		if err != nil {
			log.Println(err)
		}

	// you can have as many cases as you want, as long as you write the logic

	default:
		err := logEvent(payload)
		if err != nil {
			log.Println(err)
		}

	}
}

func authEvent(entry Payload) error {
	log.Println("entry auth", entry)

	jsonData, _ := json.MarshalIndent(entry, "", "\t")
	authServiceURL := "http://authentication-service/authenticate"
	log.Println("pass to the authServiceURL")

	request, err := http.NewRequest("POST", authServiceURL, bytes.NewBuffer(jsonData))
	log.Println("pass to the post request")
	if err != nil {
		return err
	}

	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	response, err := client.Do(request)
	log.Println("responde auth", response)
	log.Println("responde auth status code", response.StatusCode)
	if err != nil {
		return err
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusAccepted {
		return err
	}

	return nil

}

func logEvent(entry Payload) error {
	log.Println("entry log", entry)
	jsonData, _ := json.MarshalIndent(entry, "", "\t") // not use that on production
	logServiceURL := "http://logger-service/log"
	log.Println("pass to the logServiceURL")

	request, err := http.NewRequest("POST", logServiceURL, bytes.NewBuffer(jsonData))
	log.Println("Pass to the post request")
	if err != nil {
		return err
	}

	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	response, err := client.Do(request)

	if err != nil {
		return err
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusAccepted {
		return err
	}

	return nil
}
