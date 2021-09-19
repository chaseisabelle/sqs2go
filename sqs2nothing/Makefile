build:
	docker build -f Dockerfile -t sqs2go/sqs2nothing:latest .

push:
	docker push chaseisabelle/sqs2nothing:latest

release:
	make build && make push

run:
	docker run -it --rm sqs2go/sqs2nothing:latest

go:
	make build && make run

rmi:
	 docker rmi -f sqs2go/sqs2nothing:latest

get:
	go get github.com/chaseisabelle/sqs2go/sqs2go

tidy:
	go mod tidy

sync:
	make get && make tidy