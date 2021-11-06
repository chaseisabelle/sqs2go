# sqs2nothing

_forward messages from sqs to nsq_

---

### description

use this tool to consume messages from an sqs queue and forward them to an nsq queue

### example

see a fully functioning example [here](https://github.com/chaseisabelle/sqs2go-examples/sqs2nsq)

### usage

* `make`
  ```bash
  make go
  ```
  _see [Makefile](./Makefile) for commands_
* docker:
  ```bash
  docker build -t sqs2nsq .
  docker run -e WORKERS=1 -e ... sqs2nsq
  ```
  _see the [Dockerfile](./Dockerfile) for env vars_
* cli:
    ```bash
    go run main.go --help
    Usage of sqs2nsq:
      -endpoint string
            the aws endpoint
      -id string
            aws account id (leave blank for no-auth)
      -key string
            aws account key (leave blank for no-auth)
      -queue string
            the queue name
      -region string
            aws region (i.e. us-east-1)
      -retries int
            the workers number of retries (default -1)
      -secret string
            aws account secret (leave blank for no-auth)
      -timeout int
            the message visibility timeout in seconds (default 30)
      -url string
            the sqs queue url
      -wait int
            wait time in seconds
      -workers int
            the number of parallel workers to run (default 1)
      -auth-secret string
            the nsq auth secret
      -backoff int
            interval (milliseconds) between checking for new message after receiving no message (default 250)
      -backoff-multiplier int
            the nsq backoff multiplier (seconds) (default 1)
      -client-id string
            the nsq client id
      -deflate
            deflate nsq messages?
      -deflate-level int
           nsq message deflate level (1-9) (default 6)
      -endpoint string
           the aws endpoint
      -id string
           aws account id (leave blank for no-auth)
      -key string
            aws account key (leave blank for no-auth)
      -queue string
            the queue name
      -region string
            aws region (i.e. us-east-1)
      -retries int
            the workers number of retries (default -1)
        -secret string
            aws account secret (leave blank for no-auth)
      -timeout int
            the message visibility timeout in seconds (default 30)
      -to string
            the nsq host to forward the messages to (default "127.0.0.1:4150")
      -topic string
            the nsq topic to publish to
      -url string
            the sqs queue url
      -wait int
            wait time in seconds
      -workers int
           the number of parallel workers to run (default 1)

    ```

