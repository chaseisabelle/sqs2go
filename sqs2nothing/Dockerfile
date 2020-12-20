FROM golang:alpine

RUN apk update && apk add --no-cache git

WORKDIR $GOPATH/src/sqs2nothing/

COPY main.go .

RUN go get -d -v
RUN go build -o /go/bin/sqs2nothing

ENTRYPOINT ["/go/bin/sqs2nothing"]
