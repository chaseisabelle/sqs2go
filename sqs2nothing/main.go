package main

import (
	"github.com/chaseisabelle/sqs2go"
	"github.com/chaseisabelle/sqs2go/config"
)

func main() {
	sqs, err := sqs2go.New(config.Load(), handler, nil)

	if err != nil {
		panic(err)
	}

	err = sqs.Start()

	if err != nil {
		panic(err)
	}
}

func handler(_ string) error {
	return nil
}
