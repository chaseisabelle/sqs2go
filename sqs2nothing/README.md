# sqs2nothing

_forward messages from sqs to...nothing_

---

### description

use this tool to consume messages from an sqs queue and just drop them. essentially, this will just drain a queue.

### usage

* docker:
  ```bash
  docker build -t sqs2nothing .
  docker run -e WORKERS=1 -e ... sqs2nothing
  ```
  _see the [Dockerfile](./Dockerfile) for env vars_
* cli:
    ```bash
    go run main.go --help
    Usage of sqs2nothing:
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
    ```

