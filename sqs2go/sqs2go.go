package sqs2go

import (
	"errors"
	"flag"
	"fmt"
	"github.com/chaseisabelle/flagz"
	"github.com/chaseisabelle/sqsc"
	"github.com/chaseisabelle/stop"
	"log"
	"os"
	"sync"
	"time"
)

type SQS2Go struct {
	configuration *Configuration
	handler       func(string) error
	logger        func(error)
}

type Configuration struct {
	Workers int
	Backoff int
	SQSC    *sqsc.Config
}

var configuration *Configuration

func init() {
	flag.Int("workers", 1, "the number of parallel workers to run")
	flag.Int("backoff", 250, "interval (milliseconds) between checking for new message after receiving no message")
	flag.String("id", "", "aws account id (leave blank for no-auth)")
	flag.String("key", "", "aws account key (leave blank for no-auth)")
	flag.String("secret", "", "aws account secret (leave blank for no-auth)")
	flag.String("region", "", "aws region (i.e. us-east-1)")
	flag.String("url", "", "the sqs queue url")
	flag.String("queue", "", "the queue name")
	flag.String("endpoint", "", "the aws endpoint")
	flag.Int("retries", -1, "the workers number of retries")
	flag.Int("timeout", 30, "the message visibility timeout in seconds")
	flag.Int("wait", 0, "wait time in seconds")
}

func New(han func(string) error, lgr func(error)) (*SQS2Go, error) {
	if han == nil {
		return nil, errors.New("handler required")
	}

	if lgr == nil {
		lgr = func(err error) {
			_, err = fmt.Fprintln(os.Stderr, err)

			if err != nil {
				log.Println(err)
			}
		}
	}

	s2g := &SQS2Go{
		handler: han,
		logger:  lgr,
	}

	return s2g, nil
}

func (s *SQS2Go) Configure(cfg *Configuration) error {
	if cfg != nil {
		s.configuration = cfg

		return nil
	}

	if !flag.Parsed() {
		flag.Parse()
	}

	workers, err := flagz.Int("workers")

	if err != nil {
		return err
	}

	backoff, err := flagz.Int("backoff")

	if err != nil {
		return err
	}

	id, err := flagz.String("id")

	if err != nil {
		return err
	}

	key, err := flagz.String("key")

	if err != nil {
		return err
	}

	secret, err := flagz.String("secret")

	if err != nil {
		return err
	}

	region, err := flagz.String("region")

	if err != nil {
		return err
	}

	endpoint, err := flagz.String("endpoint")

	if err != nil {
		return err
	}

	queue, err := flagz.String("queue")

	if err != nil {
		return err
	}

	url, err := flagz.String("url")

	if err != nil {
		return err
	}

	retries, err := flagz.Int("retries")

	if err != nil {
		return err
	}

	timeout, err := flagz.Int("timeout")

	if err != nil {
		return err
	}

	wait, err := flagz.Int("wait")

	if err != nil {
		return err
	}

	configuration = &Configuration{
		Workers: workers,
		Backoff: backoff,
		SQSC: &sqsc.Config{
			ID:       id,
			Key:      key,
			Secret:   secret,
			Region:   region,
			Endpoint: endpoint,
			Queue:    queue,
			URL:      url,
			Retries:  retries,
			Timeout:  timeout,
			Wait:     wait,
		},
	}

	if configuration.Workers <= 0 {
		return errors.New("must have 1 or more workers")
	}

	if configuration.Backoff < 0 {
		return errors.New("backoff must be greater than or equal to 0")
	}

	s.configuration = configuration

	return nil
}

func (s *SQS2Go) Configuration() *Configuration {
	return s.configuration
}

func (s *SQS2Go) Handler() func(string) error {
	return s.handler
}

func (s *SQS2Go) Logger() func(error) {
	return s.logger
}

func (s *SQS2Go) Start() error {
	cfg := s.Configuration()

	if cfg == nil {
		return errors.New("not configured")
	}

	if cfg.Workers < 1 {
		return fmt.Errorf("1 or more workers required. invalid value %d", cfg.Workers)
	}

	if cfg.Backoff < 0 {
		return fmt.Errorf("0 or more millisecond backoff required. invalid value %d", cfg.Backoff)
	}

	han := s.Handler()
	lgr := s.Logger()

	cli, err := sqsc.New(s.configuration.SQSC)

	if err != nil {
		return err
	}

	bo := time.Duration(cfg.Backoff) * time.Millisecond
	wg := sync.WaitGroup{}

	for w := 0; w < cfg.Workers; w++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			for !stop.Stopped() {
				bod, rh, err := cli.Consume()

				if err != nil {
					lgr(err)

					continue
				}

				if bod == "" && rh == "" {
					time.Sleep(bo)

					continue
				}

				err = han(bod)

				if err != nil {
					lgr(err)

					continue
				}

				_, err = cli.Delete(rh)

				if err != nil {
					lgr(err)
				}
			}
		}()
	}

	wg.Wait()

	return nil
}
