// Package queue dequeues plot requests.
package queue

import (
	"bytes"
	"encoding/gob"
	"log"

	"github.com/ebeeton/buddhabrot-go/plotter/parameters"
	amqp "github.com/rabbitmq/amqp091-go"
)

type process func(parameters.Plot)

// Dequeue dequeues plot requests.
func Dequeue(p process) {
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
		false,  // auto-ack
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

			p(plot)

			// Setting auto-ack to false requires the consumer to acknowledge
			// that the message has been processed and can be deleted. There is
			// a 30 minute default timeout for this. See:
			// https://rabbitmq.com/consumers.html#acknowledgement-timeout
			m.Ack(false)
			log.Println("Plot complete.")
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
