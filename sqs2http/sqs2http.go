package main

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"github.com/chaseisabelle/flagz"
	"github.com/chaseisabelle/sqs2go/sqs2go"
	"net/http"
	"strings"
)

var client *http.Client
var to *string
var method *string
var headers http.Header
var accept []int

func init() {
	client = &http.Client{}
	headers = http.Header{}
}

func main() {
	to = flag.String("to", "", "the url to forward the messages to")
	method = flag.String("method", "GET", "the request method to send the message with")

	var rfz flagz.Flagz
	var hfz flagz.Flagz

	flag.Var(&rfz, "accept", "acceptable http status code(s) - i.e. it will not requeue when these codes are received from the http endpoint")
	flag.Var(&hfz, "header", "the http headers")

	s2g, err := sqs2go.New(handler, nil)

	if err != nil {
		panic(err)
	}

	err = s2g.Configure(nil)

	if err != nil {
		panic(err)
	}

	accept, err = rfz.Intz()

	if err != nil {
		panic(err)
	}

	for _, hdr := range hfz.Stringz() {
		spl := strings.SplitAfterN(hdr, ":", 2)

		if len(spl) != 2 {
			panic(fmt.Errorf("invalid header: %s", hdr))
		}

		hk := strings.TrimSpace(spl[0])

		if hk == "" {
			panic(fmt.Errorf("invalid header key: %s", hdr))
		}

		headers.Add(hk, strings.TrimSpace(spl[1]))
	}

	err = s2g.Start()

	if err != nil {
		panic(err)
	}
}

func handler(bod string) error {
	req, err := http.NewRequest(*method, *to, bytes.NewBufferString(bod))

	if err != nil {
		return err
	}

	req.Header = headers

	res, err := client.Do(req)

	if res == nil {
		if err == nil {
			err = errors.New("received nil response with no error")
		}

		return err
	}

	rsc := res.StatusCode

	for _, acc := range accept {
		if acc != rsc {
			return fmt.Errorf("unacceptable http status code %d - requeueing", rsc)
		}
	}

	return nil
}
