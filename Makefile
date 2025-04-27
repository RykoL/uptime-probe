CONTAINER_NAME=timescaledb
IMAGE=timescale/timescaledb-ha:pg17
PG_PASSWORD=password
DATABASE_URL="postgresql://postgres:$(PG_PASSWORD)@localhost:5432/postgres"

SOURCES=$(shell find . -name "*.go")
MOD_FILES=go.mod go.sum

TEMPL_SOURCES=$(shell find . -name "*.templ")
TEMPL_TARGETS=$(TEMPL_SOURCES:.templ=_templ.go)

BIN=./output/cmd/main

SHELL := /bin/bash

build: generate $(BIN)

generate: $(TEMPL_TARGETS)

test: generate
	go test ./... -v

clean: $(BIN)
	@rm $(BIN)

pre-commit: clean build test


$(BIN): $(SOURCES) $(MOD_FILES) $(TEMPL_TARGETS)
	@echo "Building main binary"
	@go build -o $(BIN) ./cmd/main.go

$(TEMPL_TARGETS): $(TEMPL_SOURCES)
	@echo "Generating templ files"
	@templ generate

run-db:
	@echo "Checking if container '$(CONTAINER_NAME)' is running..."
	@if docker ps --format '{{.Names}}' | grep -q "^$(CONTAINER_NAME)$$"; then\
		echo "Container '$(CONTAINER_NAME)' is already running."; \
	else \
		echo "Container '$(CONTAINER_NAME)' is not running. Starting it..."; \
		docker run --rm -d --name $(CONTAINER_NAME) -p 5432:5432 -e POSTGRES_PASSWORD=$(PG_PASSWORD) $(IMAGE); \
    fi

run-dev: run-db
	@DATABASE_URL=$(DATABASE_URL) templ generate --watch --proxy="http://localhost:8080" --cmd="go run cmd/main.go config/testdata/example.yaml" --open-browser=false


.PHONY: build generate run-db run-dev test clean