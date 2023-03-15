// Package workqueue allows callers to enqueue plot requests.
package workqueue

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

// Dequeue dequeues plot requests.
func Dequeue() {
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

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer.")

	var forever chan struct{}

	go func() {
		for m := range msgs {
			log.Printf("Received a message: %s", m.Body)
		}
	}()

	log.Printf("Waiting for messages. Press CTRL+C to exit.")
	<-forever
}
