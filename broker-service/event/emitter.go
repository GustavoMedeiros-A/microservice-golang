package event

import (
	"context"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Emitter struct {
	connetion *amqp.Connection
}

func (e *Emitter) setUp() error {
	channel, err := e.connetion.Channel()

	if err != nil {
		return err
	}

	defer channel.Close()
	return declareExchange(channel)
}

func (e *Emitter) Push(event string, severity string) error {
	channel, err := e.connetion.Channel()

	if err != nil {
		return err
	}

	defer channel.Close()
	log.Println("Pushing to channel")

	// err = channel.Publish("log_topic", severity, false, false, amqp.Publishing{ContentType: "text/plain", Body: []byte(event)})
	err = channel.PublishWithContext(
		context.Background(),
		"logs_topic",
		severity,
		false,
		false,
		amqp.Publishing{ContentType: "text/plain",
			Body: []byte(event)})

	if err != nil {
		return err
	}

	return nil
}

func NewEventEmitter(conn *amqp.Connection) (Emitter, error) {
	emitter := Emitter{
		connetion: conn,
	}

	err := emitter.setUp()
	if err != nil {
		return Emitter{}, err
	}

	return emitter, nil

}
