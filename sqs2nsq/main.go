package main

import (
	"flag"
	"github.com/chaseisabelle/sqs2go"
	"github.com/chaseisabelle/sqs2go/config"
	"github.com/nsqio/go-nsq"
)

var producer *nsq.Producer
var topic *string

func main() {
	to := flag.String("to", "127.0.0.1:4150", "the nsq host to forward the messages to")
	topic = flag.String("topic", "", "the nsq topic to publish to")

	sqs, err := sqs2go.New(config.Load(), handler, func(err error) {
		println(err.Error())
	})

	if err != nil {
		panic(err)
	}

	producer, err = nsq.NewProducer(*to, nsq.NewConfig())

	if err != nil {
		panic(err)
	}

	err = sqs.Start()

	if err != nil {
		panic(err)
	}
}

func handler(bod string) error {
	return producer.Publish(*topic, []byte(bod))
}
