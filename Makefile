VERSION := $(shell git describe --tags --abbrev=0 --match 'v*')
COMMIT := $(shell git rev-parse --short HEAD)
EXTERNAL_VERSION :=

all: clean backend

backend:
	go build -o bin/acmeidc -ldflags "-w -s -X 'github.com/ChuqiCloud/acmeidc/internal/version.Version=$(VERSION)$(EXTERNAL_VERSION)' -X 'github.com/ChuqiCloud/acmeidc/internal/version.Commit=$(COMMIT)'" -gcflags "-N -l" -v

gen:
	go run cmd/gen.go

fmt:
	go fmt ./...

dev:
	go build -o bin/acmeidc-dev -v
	./bin/acmeidc-dev

test:
	TESTPWD=$(shell pwd) go test ./...

clean:
	rm -rf bin/*

.PHONY: clean apidoc document ui
