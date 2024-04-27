all: clean ui backend

install-dev:
	$(MAKE) -C ui install-dev

docker-install-dev:
	$(MAKE) -C ui docker-install-dev

docker-ui:
	$(MAKE) -C ui docker-all

docker-build: clean docker-ui backend

backend:
	go build -o bin/reblog -ldflags "-w -s" -gcflags "-N -l"

gen:
	go run cmd/gen.go

apidoc:
	swag fmt
	swag init -g server/server.go -o internal/docs --parseDependency --parseInternal
	redocly build-docs internal/docs/swagger.yaml -o apidoc/index.html

document: apidoc
	$(MAKE) -C document html
	mkdir -p document/_build/html/apidoc
	cp apidoc/index.html document/_build/html/apidoc/index.html

fmt:
	gofmt -w .
	prettier -w ui
	swag fmt

ui:
	$(MAKE) -C ui

dev:
	go build -o bin/reblog-dev
	REBLOG_DEV=true ./bin/reblog-dev

clean:
	rm -rf bin/*

.PHONY: clean apidoc document ui
