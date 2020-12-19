package main

import (
	"github.com/chaseisabelle/sqs2_"
	"github.com/chaseisabelle/sqs2_/config"
)

func main() {
	sqs, err := sqs2_.New(config.Load(), handler, func(err error) {
		println(err.Error())
	})

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
