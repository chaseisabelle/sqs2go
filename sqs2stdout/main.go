package main

import (
	"flag"
	"fmt"
	"github.com/chaseisabelle/sqs2go/config"
	"github.com/chaseisabelle/sqs2go/sqs2go"
)

var delimiter *string

func main() {
	delimiter = flag.String("delimiter", "", "what to append to each write")

	sqs, err := sqs2go.New(config.Load(), handler, func(err error) {
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

func handler(bod string) error {
	print(fmt.Sprintf("%s%s", bod, *delimiter))

	return nil
}
