package adapter

import (
	"ApiRestAct1/src/asignatures/application/repositories"
	"ApiRestAct1/src/asignatures/domain/entities"
	"github.com/goccy/go-json"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
)

type RabbitMQAdapter struct {
	conn *amqp.Connection
	ch   *amqp.Channel
}

var _ repositories.IMessageService = (*RabbitMQAdapter)(nil)

func NewRabbitMQAdapter() (*RabbitMQAdapter, error) {
	conn, err := amqp.Dial("amqp://toledo:12345@35.170.134.124:5672/")
	if err != nil {
		log.Println("Error connecting to RabbitMQ:", err)
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Println("Error opening channel:", err)
		return nil, err
	}

	_, err = ch.QueueDeclare(
		"asignatures", // Queue name
		true,          // Durable
		false,         // Auto-delete
		false,         // Exclusive
		false,         // No-wait
		nil,           // Arguments
	)
	if err != nil {
		log.Println("Error declaring queue:", err)
		return nil, err
	}

	err = ch.Confirm(false)
	if err != nil {
		log.Println("Error enabling message confirmations:", err)
		return nil, err
	}

	return &RabbitMQAdapter{conn: conn, ch: ch}, nil
}

func (r *RabbitMQAdapter) PublishEvent(eventType string, asignature entities.Asignature) error {
	body, err := json.Marshal(asignature)
	if err != nil {
		log.Println("Error converting event to JSON:", err)
		return err
	}

	ack, nack := r.ch.NotifyConfirm(make(chan uint64, 1), make(chan uint64, 1))

	err = r.ch.Publish(
		"",            // Exchange
		"asignatures", // Routing key (queue name)
		true,          // Mandatory
		false,         // Immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
	if err != nil {
		log.Println("Error sending message to RabbitMQ:", err)
		return err
	}

	select {
	case <-ack:
		log.Println("Message confirmed by RabbitMQ")
	case <-nack:
		log.Println("Message was not confirmed")
	}

	log.Println("Event published:", eventType)
	return nil
}

func (r *RabbitMQAdapter) Close() {
	if err := r.ch.Close(); err != nil {
		log.Printf("Error closing RabbitMQ channel: %v", err)
	}
	if err := r.conn.Close(); err != nil {
		log.Printf("Error closing RabbitMQ connection: %v", err)
	}
}
