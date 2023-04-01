// Package queue dequeues plot requests.
package queue

import (
	"bytes"
	"encoding/gob"
	"log"

	"github.com/ebeeton/buddhabrot-go/plotter/parameters"
	amqp "github.com/rabbitmq/amqp091-go"
)

// Dequeue dequeues plot requests.
func Dequeue() {
	// TODO:: Make the hostname configurable.
	conn, err := amqp.Dial("amqp://guest:guest@rabbitmq:5672/")
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
			r := bytes.NewReader(m.Body)
			dec := gob.NewDecoder(r)
			var plot parameters.Plot
			if err := dec.Decode(&plot); err != nil {
				log.Fatal(err)
			}

			log.Printf("Received plot request: %v", plot)
		}
	}()

	log.Printf("Waiting for messages.")
	<-forever
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}