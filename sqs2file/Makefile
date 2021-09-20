build:
	docker build --no-cache -f Dockerfile -t chaseisabelle/sqs2file:latest .

push:
	docker push chaseisabelle/sqs2file:latest

release:
	make build && make push

run:
	docker run -it --rm chaseisabelle/sqs2file:latest

go:
	make build && make run

rmi:
	 docker rmi -f chaseisabelle/sqs2file:latest

get:
	go get github.com/chaseisabelle/sqs2go/sqs2go

tidy:
	go mod tidy

sync:
	go clean --modcache && make get && make tidy