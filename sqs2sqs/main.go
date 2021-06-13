package main

import (
	"flag"
	"github.com/chaseisabelle/sqs2go"
	"github.com/chaseisabelle/sqs2go/config"
	"github.com/chaseisabelle/sqsc"
)

var client *sqsc.SQSC
var delay *int

func main() {
	id2 := flag.String("id2", "", "aws account id (leave blank for no-auth)")
	key2 := flag.String("key2", "", "aws account key (leave blank for no-auth)")
	secret2 := flag.String("secret2", "", "aws account secret (leave blank for no-auth)")
	region2 := flag.String("region2", "", "aws region (i.e. us-east-1)")
	url2 := flag.String("url2", "", "the sqs queue url")
	queue2 := flag.String("queue2", "", "the queue name")
	endpoint2 := flag.String("endpoint2", "", "the aws endpoint")
	delay = flag.Int("delay", 0, "the delay for the produced message")

	sqs, err := sqs2go.New(config.Load(), handler, func(err error) {
		println(err.Error())
	})

	if err != nil {
		panic(err)
	}

	client, err = sqsc.New(&sqsc.Config{
		ID:       *id2,
		Key:      *key2,
		Secret:   *secret2,
		Region:   *region2,
		Endpoint: *endpoint2,
		Queue:    *queue2,
		URL:      *url2,
	})

	if err != nil {
		panic(err)
	}

	err = sqs.Start()

	if err != nil {
		panic(err)
	}
}

func handler(bod string) error {
	_, err := client.Produce(bod, *delay)

	return err
}
