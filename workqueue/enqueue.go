// Package workqueue allows callers to enqueue plot requests.
package workqueue

import (
	"context"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

const (
	timeout = 5 * time.Second
)

// Enqueue enqueues plot requests.
func Enqueue(request string) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ.")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel.")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   //exclusive
		false,   // no-wait
		nil,     //arguments
	)
	failOnError(err, "Failed to declare the queue.")

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	err = ch.PublishWithContext(ctx,
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(request),
		})
	failOnError(err, "Failed to publish the request.")

	log.Printf("Sent %s.", request)
}
