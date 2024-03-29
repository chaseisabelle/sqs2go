package main

import (
	"flag"
	"fmt"
	"github.com/chaseisabelle/sqs2go/sqs2go"
	"log"
	"os"
)

var file *os.File
var filename *string
var delimiter *string
var permissions *uint64

func main() {
	filename = flag.String("filename", "", "the name of the file to write to")
	delimiter = flag.String("delimiter", "", "what to append to each write")
	permissions = flag.Uint64("permissions", 0644, "file permissions")

	s2g, err := sqs2go.New(handler, nil)

	if err != nil {
		log.Fatalln(err)
	}

	err = s2g.Configure(nil)

	if err != nil {
		log.Fatalln(err)
	}

	file, err = os.OpenFile(*filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.FileMode(*permissions))

	if err != nil {
		log.Fatalln(err)
	}

	err = s2g.Start()

	if err != nil {
		log.Fatalln(err)
	}
}

func handler(bod string) error {
	_, err := file.WriteString(fmt.Sprintf("%s%s", bod, *delimiter))

	return err
}
