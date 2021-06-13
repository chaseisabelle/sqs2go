FROM golang:alpine

RUN apk update && apk add --no-cache git

WORKDIR $GOPATH/src/sqs2nothing/

COPY main.go .
COPY go.* ./

RUN go get -d -v
RUN go build -o /go/bin/sqs2nothing

CMD /go/bin/sqs2nothing \
    --workers=$WORKERS \
    --id=$AWS_ID \
    --key=$AWS_KEY \
    --secret=$AWS_SECRET \
    --region=$AWS_REGION \
    --url=$QUEUE_URL \
    --queue=$QUEUE_NAME \
    --endpoint=$QUEUE_ENDPOINT \
    --retries=$MAX_RETRIES \
    --timeout=$VISIBILITY_TIMEOUT \
    --wait=$WAIT_TIME
