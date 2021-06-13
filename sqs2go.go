package sqs2go

import (
	"errors"
	"fmt"
	"github.com/chaseisabelle/sqs2go/config"
	"github.com/chaseisabelle/sqsc"
	"github.com/chaseisabelle/stop"
	"os"
	"sync"
)

type SQS2Go struct {
	config  *config.Config
	client  *sqsc.SQSC
	handler func(string) error
	logger  func(error)
}

func New(cfg *config.Config, han func(string) error, lgr func(error)) (*SQS2Go, error) {
	if han == nil {
		return nil, errors.New("handler required")
	}

	if cfg.Workers < 1 {
		return nil, fmt.Errorf("1 or more workers required. invalid value %d", cfg.Workers)
	}

	con, err := cfg.SQSC()

	if err != nil {
		return nil, err
	}

	cli, err := sqsc.New(con)

	if lgr == nil {
		lgr = func(err error) {
			fmt.Fprintln(os.Stderr, err.Error())
		}
	}

	return &SQS2Go{
		config:  cfg,
		client:  cli,
		handler: han,
		logger:  lgr,
	}, err
}

func (s *SQS2Go) Config() *config.Config {
	return s.config
}

func (s *SQS2Go) Client() *sqsc.SQSC {
	return s.client
}

func (s *SQS2Go) Handler() func(string) error {
	return s.handler
}

func (s *SQS2Go) Logger() func(error) {
	return s.logger
}

func (s *SQS2Go) Start() error {
	cfg := s.Config()
	cli := s.Client()
	han := s.Handler()
	lgr := s.Logger()
	wg := sync.WaitGroup{}

	for w := 0; w < cfg.Workers; w++ {
		wg.Add(1)

		go func(w int) {
			defer wg.Done()

			for !stop.Stopped() {
				bod, rh, err := cli.Consume()

				if err != nil {
					lgr(err)

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
		}(w)
	}

	wg.Wait()

	return nil
}
