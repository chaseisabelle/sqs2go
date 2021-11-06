package main

import (
	"flag"
	"fmt"
	"github.com/chaseisabelle/sqs2go/sqs2go"
	"log"
)

var delimiter *string

func main() {
	delimiter = flag.String("delimiter", "", "what to append to each write")

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
	fmt.Printf("%s%s", bod, *delimiter)

	return nil
}
