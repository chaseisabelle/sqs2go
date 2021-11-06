package main

import (
	"github.com/chaseisabelle/sqs2go/sqs2go"
	"log"
)

func main() {
	s2g, err := sqs2go.New(handler, nil)

	if err != nil {
		log.Fatalln(err)
	}

	err = s2g.Configure(nil)

	if err != nil {
		log.Fatalln(err)
	}

	err = s2g.Start()

	if err != nil {
		log.Fatalln(err)
	}
}

func handler(bod string) error {
	log.Println(bod)

	return nil
}
