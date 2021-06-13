package main

import (
	"github.com/chaseisabelle/sqs2go"
	"github.com/chaseisabelle/sqs2go/config"
)

func main() {
	s2g, err := sqs2go.New(config.Load(), handler, nil)

	if err != nil {
		panic(err)
	}

	err = s2g.Start()

	if err != nil {
		panic(err)
	}
}

func handler(_ string) error {
	return nil
}
