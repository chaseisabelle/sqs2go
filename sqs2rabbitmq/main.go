package main

import (
	"flag"
	"fmt"
	"github.com/chaseisabelle/sqs2_"
	"github.com/chaseisabelle/sqs2_/config"
	"github.com/streadway/amqp"
)

var channel *amqp.Channel
var queue amqp.Queue

func main() {
	hst := flag.String("rabbitmq-host", "localhost:5672", "the rabbitmq host")
	usr := flag.String("rabbitmq-user", "guest", "the rabbitmq username")
	pwd := flag.String("rabbitmq-pass", "guest", "the rabbitmq password")
	que := flag.String("rabbitmq-queue", "", "the rabbitmq queue name")
	dur := flag.Bool("durable")

	sqs, err := sqs2_.New(config.Load(), handler, func(err error) {
		println(err.Error())
	})

	if err != nil {
		panic(err)
	}

	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s/", *usr, *pwd, *hst))

	if err != nil {
		panic(err)
	}

	defer func() {
		err := conn.Close()

		if err != nil {
			panic(err)
		}
	}()

	channel, err = conn.Channel()

	if err != nil {
		panic(err)
	}

	defer func() {
		err := channel.Close()

		if err != nil {
			panic(err)
		}
	}()

	queue, err = channel.QueueDeclare(
		*que,  // name
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)

	if err != nil {
		panic(err)
	}

	err = sqs.Start()

	if err != nil {
		panic(err)
	}
}

func handler(bod string) error {
	return channel.Publish(
		"",     // exchange
		queue.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(bod),
		})
}
