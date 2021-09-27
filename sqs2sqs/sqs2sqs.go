package main

import (
	"flag"
	"github.com/chaseisabelle/sqs2go/sqs2go"
	"github.com/chaseisabelle/sqsc"
)

var client *sqsc.SQSC
var delay *int

func main() {
	id := flag.String("producer-id", "", "producer aws account id (leave blank for no-auth)")
	key := flag.String("producer-key", "", "producer aws account key (leave blank for no-auth)")
	secret := flag.String("producer-secret", "", "producer aws account secret (leave blank for no-auth)")
	region := flag.String("producer-region", "", "producer aws region (i.e. us-east-1)")
	url := flag.String("producer-url", "", "producer sqs queue url")
	queue := flag.String("producer-queue", "", "producer queue name")
	endpoint := flag.String("producer-endpoint", "", "producer aws endpoint")
	delay = flag.Int("producer-delay", 0, "the delay for the produced message")

	s2g, err := sqs2go.New(handler, nil)

	if err != nil {
		panic(err)
	}

	err = s2g.Configure(nil)

	if err != nil {
		panic(err)
	}

	client, err = sqsc.New(&sqsc.Config{
		ID:       *id,
		Key:      *key,
		Secret:   *secret,
		Region:   *region,
		Endpoint: *endpoint,
		Queue:    *queue,
		URL:      *url,
	})

	if err != nil {
		panic(err)
	}

	err = s2g.Start()

	if err != nil {
		panic(err)
	}
}

func handler(bod string) error {
	_, err := client.Produce(bod, *delay)

	return err
}
