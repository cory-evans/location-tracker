.PHONY : all
.DEFAULT_GOAL := all


BINARY_NAME=gps-tracker

arm64:
	GOOS=linux GOARCH=arm64 go build -o bin/${BINARY_NAME}_arm64 main.go

amd64:
	GOOS=linux GOARCH=amd64 go build -o bin/${BINARY_NAME}_amd64 main.go

clean:
	go clean
	rm bin/*

all: arm64 amd64