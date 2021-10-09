# builder image
FROM golang:1.17.1-alpine3.14 AS sqs2nsq-builder

RUN mkdir /build

ADD *.go /build/
ADD go.mod /build/
ADD go.sum /build/

WORKDIR /build

RUN go get -d -v
RUN CGO_ENABLED=0 GOOS=linux go build -a -o sqs2nsq .


# final image
FROM alpine:3.14 AS sqs2nsq

COPY --from=sqs2nsq-builder /build/sqs2nsq .

CMD ./sqs2nsq \
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
    --wait=$WAIT_TIME \
    --to=$TO \
    --topic=$TOPIC
