.PHONY: test format install

EXEC=main

format:
	go fmt .

test:
	go test

install:
	go build -o todizer	cmd/main.go 
	chmod +x todizer
	cp todizer ~/go/bin
