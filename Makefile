all: build test

init:
	go get -u github.com/prometheus/promu

build:
	go install -v
	promu build

test:
	go test -v -race