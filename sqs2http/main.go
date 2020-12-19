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

var client *http.Client //<< http client
var to *string //<< http endpoint
var method *string //<< http request method
var requeue []int //<< only requeue if http response code meets one of these
var onFail bool //<< if no requeue params given, default to requeue on any !2xx status code

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

	requeue, err = flags.Intz()

	if err != nil {
		panic(err)
	}

	onFail = len(requeue) == 0
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

	rsc := res.StatusCode

	if onFail && (rsc < 200 || rsc > 299) {
		return statusCodeError(rsc, err)
	}

	for _, sc := range requeue {
		if sc == rsc {
			return statusCodeError(rsc, err)
		}
	}

	return nil
}

func statusCodeError(sc int, err error) error {
	if err == nil {
		err = fmt.Errorf("received %d response with no error", sc)
	}

	return err
}
