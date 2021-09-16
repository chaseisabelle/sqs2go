package main

import (
	"flag"
	"fmt"
	"github.com/chaseisabelle/sqs2go/config"
	"github.com/chaseisabelle/sqs2go/sqs2go"
	"github.com/streadway/amqp"
)

var channel *amqp.Channel
var queue amqp.Queue
var exchange *string
var mandatory *bool
var immediate *bool
var contentType *string

func main() {
	hst := flag.String("rabbitmq-host", "localhost:5672", "the rabbitmq host")
	usr := flag.String("rabbitmq-user", "guest", "the rabbitmq username")
	pwd := flag.String("rabbitmq-pass", "guest", "the rabbitmq password")
	que := flag.String("rabbitmq-queue", "", "the rabbitmq queue name")
	dur := flag.Bool("durable", false, "durable? idk")
	ad := flag.Bool("auto-delete", false, "auto delete?")
	exc := flag.Bool("exclusive", false, "exclusive?")
	nw := flag.Bool("no-wait", false, "no wait?")

	exchange = flag.String("exchange", "", "exchange?")
	mandatory = flag.Bool("mandatory", false, "mandatory?")
	immediate = flag.Bool("immediate", false, "immediate?")
	contentType = flag.String("content-type", "text/plain", "the content type")

	sqs, err := sqs2go.New(config.Load(), handler, func(err error) {
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
		*que, // name
		*dur, // durable
		*ad,  // delete when unused
		*exc, // exclusive
		*nw,  // no-wait
		nil,  // arguments
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
		*exchange,  // exchange
		queue.Name, // routing key
		*mandatory, // mandatory
		*immediate, // immediate
		amqp.Publishing{
			ContentType: *contentType,
			Body:        []byte(bod),
		})
}
