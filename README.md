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
func handler(_ string) error {
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
