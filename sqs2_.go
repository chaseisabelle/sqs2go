package sqs2_

import (
	"errors"
	"fmt"
	"github.com/chaseisabelle/sqs2_/config"
	"github.com/chaseisabelle/sqsc"
	"github.com/chaseisabelle/stop"
	"sync"
)

type SQS2_ struct {
	config  *config.Config
	client  *sqsc.SQSC
	handler func(string) error
	logger  func(error)
}

func New(cfg *config.Config, han func(string) error, lgr func(error)) (*SQS2_, error) {
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

	return &SQS2_{
		config:  cfg,
		client:  cli,
		handler: han,
		logger:  lgr,
	}, err
}

func (s *SQS2_) Config() *config.Config {
	return s.config
}

func (s *SQS2_) Client() *sqsc.SQSC {
	return s.client
}

func (s *SQS2_) Handler() func(string) error {
	return s.handler
}

func (s *SQS2_) Start() error {
	cfg := s.Config()
	cli := s.Client()
	wg := sync.WaitGroup{}

	for w := 0; w < cfg.Workers; w++ {
		wg.Add(1)

		go func(w int) {
			defer wg.Done()

			for !stop.Stopped() {
				bod, rh, err := cli.Consume()

				if err != nil {
					s.logger(err)

					continue
				}

				err = s.Handler()(bod)

				if err != nil {
					s.logger(err)

					continue
				}

				_, err = cli.Delete(rh)

				if err != nil {
					s.logger(err)
				}
			}
		}(w)
	}

	wg.Wait()

	return nil
}
