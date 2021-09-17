package main

import (
	"github.com/chaseisabelle/sqs2go/sqs2go"
)

func main() {
	s2g, err := sqs2go.New(handler, nil, nil)

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
