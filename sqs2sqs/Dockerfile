# builder image
FROM golang:1.17.1-alpine3.14 AS sqs2sqs-builder

RUN mkdir /build

ADD *.go /build/
ADD go.mod /build/
ADD go.sum /build/

WORKDIR /build

RUN go get -d -v
RUN CGO_ENABLED=0 GOOS=linux go build -a -o sqs2sqs .


# final image
FROM alpine:3.14 AS sqs2sqs

COPY --from=sqs2sqs-builder /build/sqs2sqs .

CMD ./sqs2sqs \
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
    --producer-id=$PRODUCER_ID \
    --producer-key=$PRODUCER_KEY \
    --producer-secret=$PRODUCER_SECRET \
    --producer-region=$PRODUCER_REGION \
    --producer-url=$PRODUCER_URL \
    --producer-queue=$PRODUCER_QUEUE \
    --producer-endpoint=$PRODUCER_ENDPOINT \
    --producer-delay=$PRODUCER_DELAY