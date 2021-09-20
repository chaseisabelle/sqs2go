# builder image
FROM golang:1.17.1-alpine3.14 AS sqs2log-builder

RUN mkdir /build

ADD *.go /build/
ADD go.mod /build/
ADD go.sum /build/

WORKDIR /build

RUN go get -d -v
RUN CGO_ENABLED=0 GOOS=linux go build -a -o sqs2log .


# final image
FROM alpine:3.14 AS sqs2log

COPY --from=sqs2log-builder /build/sqs2log .

CMD ./sqs2log \
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