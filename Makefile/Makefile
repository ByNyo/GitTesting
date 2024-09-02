hello:
	echo "hello"

build:
	go build -o bin/main main.go

run:
	go run main.go

run-build:
	./bin/main

	.PHONY: run-background

run-background:
	@go run main.go & PID=$$!; \
	echo "Process ID is $$PID"; \
	ps aux | grep $$PID; \
	wait $$PID; \
	echo "Process $$PID has finished."