.PHONY: run-background

build:
	go build -o bin/main main.go

run:
	go run main.go

run-build:
	./bin/main

run-background:
	@go run main.go & PID=$$!; \
	echo "Process ID is $$PID"; \
	ps axjf; \
	wait $$PID; \
	echo "Process $$PID has finished."