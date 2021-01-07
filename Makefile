GO    := GO15VENDOREXPERIMENT=1 go
PROMU := $(GOPATH)/bin/promu
pkgs   = $(shell $(GO) list ./... | grep -v /vendor/)

PREFIX                  ?= $(shell pwd)
BIN_DIR                 ?= $(shell pwd)
DOCKER_IMAGE_NAME       ?= volume_exporter
DOCKER_IMAGE_TAG        ?= $(subst /,-,$(shell git rev-parse --abbrev-ref HEAD))
TAG 					:= $(shell echo `if [ "$(TRAVIS_BRANCH)" = "master" ] || [ "$(TRAVIS_BRANCH)" = "" ] ; then echo "latest"; else echo $(TRAVIS_BRANCH) ; fi`)

all: build test

init: 
	@$(GO) get -u github.com/prometheus/promu
	@$(GO) get -u github.com/prometheus/client_golang/prometheus
	@$(GO) get -u github.com/prometheus/common/version
	@$(GO) get -u github.com/prometheus/common/log
	@$(GO) get -u github.com/prometheus/client_golang/prometheus/promhttp

build: 
	@$(GO) go install -v
	@$(PROMU) build --prefix $(PREFIX)

docker: 
	@docker build -t "$(DOCKER_IMAGE_NAME):$(DOCKER_IMAGE_TAG)" .

push: 
	@docker login -u $(DOCKER_USERNAME) -p $(DOCKER_PASSWORD)
	@docker tag "$(DOCKER_IMAGE_NAME):$(DOCKER_IMAGE_TAG)" "$(DOCKER_USERNAME)/$(DOCKER_IMAGE_NAME):$(TAG)"
	@docker push "$(DOCKER_USERNAME)/$(DOCKER_IMAGE_NAME):$(TAG)"

test: 
	@$(GO) test -short $(pkgs)