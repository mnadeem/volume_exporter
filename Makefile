all: build test

init:
	go get -u github.com/prometheus/promu
	go get -u github.com/prometheus/client_golang/prometheus
	go get -u github.com/prometheus/common/version
	go get -u github.com/prometheus/common/log
	go get -u github.com/prometheus/client_golang/prometheus/promhttp

build:
	go install -v
	promu build

test:
	go test -v -race