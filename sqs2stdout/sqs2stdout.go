package main

import (
	"flag"
	"fmt"
	"github.com/chaseisabelle/sqs2go/sqs2go"
)

var delimiter *string

func main() {
	delimiter = flag.String("delimiter", "", "what to append to each write")

	s2g, err := sqs2go.New(handler, nil)

	if err != nil {
		panic(err)
	}

	err = s2g.Configure(nil)

	if err != nil {
		panic(err)
	}

	err = s2g.Start()

	if err != nil {
		panic(err)
	}
}

func handler(bod string) error {
	fmt.Printf("%s%s", bod, *delimiter)

	return nil
}
