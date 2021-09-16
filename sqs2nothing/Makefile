build:
	docker build -f Dockerfile -t sqs2go/sqs2nothing:latest .

run:
	docker run -it --rm sqs2go/sqs2nothing:latest

go:
	make build && make run

rmi:
	 docker rmi -f sqs2go/sqs2nothing:latest