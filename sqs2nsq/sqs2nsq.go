package main

import (
	"flag"
	"github.com/chaseisabelle/sqs2go/sqs2go"
	"github.com/nsqio/go-nsq"
	"log"
	"time"
)

var producer *nsq.Producer
var topic *string

func main() {
	to := flag.String("to", "127.0.0.1:4150", "the nsq host to forward the messages to")
	topic = flag.String("topic", "", "the nsq topic to publish to")

	cid := flag.String("client-id", "", "the nsq client id")
	aut := flag.String("auth-secret", "", "the nsq auth secret")
	bom := flag.Int("backoff-multiplier", 1, "the nsq backoff multiplier (seconds)")
	def := flag.Bool("deflate", false, "deflate nsq messages?")
	dl := flag.Int("deflate-level", 6, "nsq message deflate level (1-9)")

	s2g, err := sqs2go.New(handler, nil)

	if err != nil {
		log.Fatalln(err)
	}

	err = s2g.Configure(nil)

	if err != nil {
		log.Fatalln(err)
	}

	cfg := nsq.NewConfig()

	cfg.ClientID = *cid
	cfg.AuthSecret = *aut
	cfg.Deflate = *def
	cfg.DeflateLevel = *dl
	cfg.BackoffMultiplier = time.Duration(*bom) * time.Second

	producer, err = nsq.NewProducer(*to, cfg)

	if err != nil {
		log.Fatalln(err)
	}

	err = s2g.Start()

	producer.Stop()

	if err != nil {
		log.Fatalln(err)
	}
}

func handler(bod string) error {
	return producer.Publish(*topic, []byte(bod))
}
