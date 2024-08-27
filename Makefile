VERSION := $(shell git describe --tags --abbrev=0 --match 'v*')
COMMIT := $(shell git rev-parse --short HEAD)
EXTERNAL_VERSION :=

all: clean ui backend

install-dev:
	pnpm install --frozen-lockfile

backend:
	go build -o bin/reblog -ldflags "-w -s -X 'github.com/redish101/reblog/internal/version.Version=$(VERSION)$(EXTERNAL_VERSION)' -X 'github.com/redish101/reblog/internal/version.Commit=$(COMMIT)'" -v

gen:
	go run cmd/gen.go

apidoc:
	swag fmt
	swag init -g server/server.go -o internal/docs --parseDependency --parseInternal
	redocly build-docs internal/docs/swagger.yaml -o apidoc/index.html

fmt:
	go fmt ./...
	prettier -w packages
	swag fmt

ui:
	pnpm build
	rm -rf internal/ui/dist
	cp -r packages/dashboard/dist internal/ui/dist

dev:
	go build -o bin/reblog-dev -v
	./bin/reblog-dev

test:
	TESTPWD=$(shell pwd) go test ./...

clean:
	rm -rf bin/*

.PHONY: clean apidoc document ui
