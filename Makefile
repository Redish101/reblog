VERSION := 0.1.0-alpha
COMMIT := $(shell git rev-parse --short HEAD)

all: clean ui backend

install-dev:
	$(MAKE) -C ui install-dev

backend:
	go build -o bin/reblog -ldflags "-w -s -X 'reblog/internal/version.Version=$(VERSION)' -X 'reblog/internal/version.Commit=$(COMMIT)'" -gcflags "-N -l"

gen:
	go run cmd/gen.go

apidoc:
	swag fmt
	swag init -g server/server.go -o internal/docs --parseDependency --parseInternal
	redocly build-docs internal/docs/swagger.yaml -o apidoc/index.html

fmt:
	gofmt -w .
	prettier -w ui
	swag fmt

ui:
	$(MAKE) -C ui

dev:
	go build -o bin/reblog-dev
	./bin/reblog-dev

clean:
	rm -rf bin/*

.PHONY: clean apidoc document ui
