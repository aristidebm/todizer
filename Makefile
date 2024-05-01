.PHONY: test format

EXEC=main

format:
	go fmt .

test:
	go test
