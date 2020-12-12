package main

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"github.com/chaseisabelle/flagz"
	"github.com/chaseisabelle/sqs2_"
	"github.com/chaseisabelle/sqs2_/config"
	"net/http"
)

var client *http.Client
var to *string
var method *string
var requeue []int

func main() {
	to = flag.String("to", "", "the url to forward the messages to")
	method = flag.String("method", "GET", "the request method to send the message with")

	var flags flagz.Flagz

	flag.Var(&flags, "requeue", "the http status code to requeue a message for")

	sqs, err := sqs2_.New(config.Load(), handler, func(err error) {
		println(err.Error())
	})

	if err != nil {
		panic(err)
	}

	client = &http.Client{}

	err = sqs.Start()

	if err != nil {
		panic(err)
	}
}

func handler(bod string) error {
	req, err := http.NewRequest(*method, *to, bytes.NewBufferString(bod))

	if err != nil {
		return err
	}

	res, err := client.Do(req)

	if res == nil {
		if err == nil {
			err = errors.New("received nil response with no error")
		}

		return err
	}

	for _, sc := range requeue {
		if sc == res.StatusCode {
			if err == nil {
				err = fmt.Errorf("received %d response with no error", sc)
			}

			return err
		}
	}

	return nil
}
