.PHONY: build  package start-docker
DIST_PATH = dist
PROJECT_NAME = smart-edu-server
RELEASE_PATH = $(DIST_PATH)/release
BUILD_DATE = $(shell date -u)
BUILD_HASH = $(shell git rev-parse HEAD)
GO = go
GINKGO = ginkgo
GO_LINKER_FLAGS ?= -ldflags \
	   "-X '$(PROJECT_NAME)/model.BuildDate=$(BUILD_DATE)' \
	   -X $(PROJECT_NAME)/model.BuildHash=$(BUILD_HASH)"
BUILDER_GOOS_GOARCH=$(shell $(GO) env GOOS)_$(shell $(GO) env GOARCH)
PACKAGESLISTS=$(shell $(GO) list ./...)
TESTFLAGS ?= -short
PACKAGESLISTS_COMMA=$(shell echo $(PACKAGESLISTS) | tr ' ' ',')
ROOT := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))

build:
	@echo Build Linux amd64
	env GOOS=linux GOARCH=amd64 $(GO) install -i $(GOFLAGS) $(GO_LINKER_FLAGS) ./...

package:
	@echo Packaging $(PROJECT_NAME)
	rm -rf $(DIST_PATH)
	mkdir -p $(RELEASE_PATH)/bin
	mkdir -p $(RELEASE_PATH)/config
	mkdir -p $(RELEASE_PATH)/logs
	cp $(GOBIN)/$(PROJECT_NAME) $(RELEASE_PATH)/bin
	cp config/*.json $(RELEASE_PATH)/config
	tar -C $(DIST_PATH) -czf $(RELEASE_PATH)-$(BUILDER_GOOS_GOARCH).tar.gz release

govet: ## Runs govet against all packages.
	@echo Running GOVET
	$(GO) vet $(GOFLAGS) $(PACKAGESLISTS) || exit 1

gofmt: ## Runs gofmt against all packages.
	@echo Running GOFMT
	@echo $(PACKAGESLISTS)
	@for package in $(PACKAGESLISTS);do \
		echo "Checking "$$package; \
		files=$$(go list -f '{{range .GoFiles}}{{$$.Dir}}/{{.}} {{end}}' $$package); \
		if [ "$$files" ]; then \
			gofmt_output=$$(gofmt -d -s $$files 2>&1); \
			if [ "$$gofmt_output" ]; then \
				echo "$$gofmt_output"; \
				echo "gofmt failure"; \
				exit 1; \
			fi; \
		fi; \
	done
	@echo "gofmt success"; \

test: clean-docker start-docker test-server clean-docker

test-server:
	@echo test-logs
	docker run -d  --name ginkgo-test -p 8089:8089 --network my-net -v $(GOPATH)/src/$(PROJECT_NAME):/go/src/$(PROJECT_NAME)  golang:latest /bin/sh -c  "/go/src/$(PROJECT_NAME)/ginkgo -r -trace -cover  -coverprofile=coverprofile.txt -outputdir=/go/src/$(PROJECT_NAME)   /go/src/$(PROJECT_NAME)"
	docker logs -f ginkgo-test


start-docker: ## Starts the docker containers for local development.
	@echo Starting docker containers

	@if [ $(shell docker ps -a | grep -ci ginkgo-mysql) -eq 0 ]; then \
		echo starting ginkgo-mysql; \
		docker run --network my-net --name ginkgo-mysql  -e MYSQL_ROOT_PASSWORD=123456 \
		-e MYSQL_USER=ginkgo -e MYSQL_PASSWORD=123456 -e MYSQL_DATABASE=ginkgo-test -d mysql:5.7 > /dev/null; \
	elif [ $(shell docker ps | grep -ci ginkgo-mysql) -eq 0 ]; then \
		echo restarting ginkgo-mysql; \
		docker start ginkgo-mysql > /dev/null; \
	fi

	@if [ $(shell docker ps -a | grep -ci ginkgo-redis) -eq 0 ]; then \
		echo starting ginkgo-redis; \
		docker run --network my-net --name ginkgo-redis -d redis:3.2 redis-server --appendonly yes; \
	elif [ $(shell docker ps | grep -ci ginkgo-redis) -eq 0 ]; then \
		echo restarting ginkgo-redis; \
		docker start ginkgo-redis > /dev/null; \
	fi

stop-docker: ## Stops the docker containers for local development.
	@echo Stopping docker containers

	@if [ $(shell docker ps -a | grep -ci ginkgo-mysql) -eq 1 ]; then \
		echo stopping ginkgo-mysql; \
		docker stop ginkgo-mysql > /dev/null; \
	fi

	@if [ $(shell docker ps -a | grep -ci ginkgo-redis) -eq 1 ]; then \
		echo stopping  ginkgo-redis; \
		docker stop  ginkgo-redis > /dev/null; \
	fi

clean-docker: ## Deletes the docker containers for local development.
	@echo Removing docker containers

	@if [ $(shell docker ps -a | grep -ci ginkgo-mysql) -eq 1 ]; then \
		echo removing ginkgo-mysql; \
		docker stop ginkgo-mysql > /dev/null; \
		docker rm -v ginkgo-mysql > /dev/null; \
	fi

	@if [ $(shell docker ps -a | grep -ci ginkgo-redis) -eq 1 ]; then \
		echo removing  ginkgo-redis; \
		docker stop  ginkgo-redis > /dev/null; \
		docker rm -v  ginkgo-redis > /dev/null; \
	fi

	@if [ $(shell docker ps -a | grep -ci ginkgo-test) -eq 1 ]; then \
		echo removing  ginkgo-test; \
		docker stop  ginkgo-test > /dev/null; \
		docker rm -v  ginkgo-test > /dev/null; \
	fi

