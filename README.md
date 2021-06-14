# sqs2go

_golang forwarding agents for sqs_

---

### usage

#### get the package

`go get github.com/chaseisabelle/sqs2go`

#### implement the package

```go
package main

import (
	"github.com/chaseisabelle/sqs2go"
	"github.com/chaseisabelle/sqs2go/config"
)

func main() {
	s2g, err := sqs2go.New(config.Load(), handler, nil)

	if err != nil {
		panic(err)
	}

	err = s2g.Start()

	if err != nil {
		panic(err)
	}
}

// this is where you can implement your custom logic
func handler(msg string) error {
	...

	return nil //<< return nil on success, or error to retry
}
```

#### implement a custom error logger

```go
logger := func(err error) {
    fmt.Fprintln(os.Stderr, err.Error())
}

s2g, err := sqs2go.New(config.Load(), handler, logger)
```

or supress error logging

```go
s2g, err := sqs2go.New(config.Load(), handler, func (_ error) {})
```

#### default flags/options

```
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
