package main

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"github.com/chaseisabelle/flagz"
	"github.com/chaseisabelle/sqs2go/sqs2go"
	"log"
	"net/http"
	"strings"
)

var client *http.Client
var to *string
var method *string
var headers http.Header
var accept map[int]interface{}

func init() {
	client = &http.Client{}
	headers = http.Header{}
	accept = make(map[int]interface{})
}

func main() {
	to = flag.String("to", "", "the url to forward the messages to")
	method = flag.String("method", "GET", "the request method to send the message with")

	var afz flagz.Flagz
	var hfz flagz.Flagz

	flag.Var(&afz, "accept", "acceptable http status code(s) - i.e. it will not requeue when these codes are received from the http endpoint")
	flag.Var(&hfz, "header", "the http headers")

	flag.Parse()

	s2g, err := sqs2go.New(handler, nil)

	if err != nil {
		log.Fatalln(err)
	}

	err = s2g.Configure(nil)

	if err != nil {
		log.Fatalln(err)
	}

	acc, err := afz.Intz()

	if err != nil {
		log.Fatalln(err)
	}

	for _, arc := range acc {
		accept[arc] = nil
	}

	for _, hdr := range hfz.Stringz() {
		spl := strings.SplitAfterN(hdr, ":", 2)

		if len(spl) != 2 {
			log.Fatalln(fmt.Errorf("invalid header: %s", hdr))
		}

		hk := strings.TrimSpace(spl[0])
		hk = hk[:len(hk)-1]

		if hk == "" {
			log.Fatalln(fmt.Errorf("invalid header key: %s", hdr))
		}

		headers.Add(hk, strings.TrimSpace(spl[1]))
	}

	err = s2g.Start()

	if err != nil {
		log.Fatalln(err)
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
	_, ok := accept[rsc]

	if !ok {
		err = fmt.Errorf("unacceptable http status code %d - requeueing", rsc)
	}

	return err
}
