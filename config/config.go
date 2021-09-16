package config

import (
	"errors"
	"flag"
	"github.com/chaseisabelle/sqsc"
)

type Config struct {
	Workers int
	AWS     struct {
		ID       string
		Key      string
		Secret   string
		Region   string
		Endpoint string
	}
	SQS struct {
		Queue   string
		URL     string
		Retries int
		Timeout int
		Wait    int
	}
}

func Load() *Config {
	workers := flag.Int("workers", 1, "the number of parallel workers to run")
	id := flag.String("id", "", "aws account id (leave blank for no-auth)")
	key := flag.String("key", "", "aws account key (leave blank for no-auth)")
	secret := flag.String("secret", "", "aws account secret (leave blank for no-auth)")
	region := flag.String("region", "", "aws region (i.e. us-east-1)")
	url := flag.String("url", "", "the sqs queue url")
	queue := flag.String("queue", "", "the queue name")
	endpoint := flag.String("endpoint", "", "the aws endpoint")
	retries := flag.Int("retries", -1, "the workers number of retries")
	timeout := flag.Int("timeout", 30, "the message visibility timeout in seconds")
	wait := flag.Int("wait", 0, "wait time in seconds")

	flag.Parse()

	return &Config{
		Workers: *workers,
		AWS: struct {
			ID       string
			Key      string
			Secret   string
			Region   string
			Endpoint string
		}{
			ID:       *id,
			Key:      *key,
			Secret:   *secret,
			Region:   *region,
			Endpoint: *endpoint,
		},
		SQS: struct {
			Queue   string
			URL     string
			Retries int
			Timeout int
			Wait    int
		}{
			Queue:   *queue,
			URL:     *url,
			Retries: *retries,
			Timeout: *timeout,
			Wait:    *wait,
		},
	}
}

func (c *Config) SQSC() (*sqsc.Config, error) {
	reg := c.AWS.Region

	if reg == "" {
		return nil, errors.New("aws region required")
	}

	cfg := &sqsc.Config{
		ID:       c.AWS.ID,
		Key:      c.AWS.Key,
		Secret:   c.AWS.Secret,
		Region:   reg,
		Endpoint: c.AWS.Endpoint,
		Queue:    c.SQS.Queue,
		URL:      c.SQS.URL,
		Retries:  c.SQS.Retries,
		Timeout:  c.SQS.Timeout,
		Wait:     c.SQS.Wait,
	}

	return cfg, nil
}
