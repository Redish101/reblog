all: clean backend

backend:
	go build -o bin/reblog -ldflags "-w -s" -gcflags "-N -l"

gen:
	go run cmd/gen.go

apidoc:
	swag fmt
	swag init -g server/server.go -o internal/docs --parseDependency --parseInternal
	redocly build-docs internal/docs/swagger.yaml -o apidoc/index.html

document: apidoc
	make -C document html
	mkdir -p document/_build/html/apidoc
	cp apidoc/index.html document/_build/html/apidoc/index.html

fmt:
	gofmt -w .
	prettier -w ui
	swag fmt

dev:
	go build -o bin/reblog-dev
	./bin/reblog-dev

clean:
	rm -rf bin/*

.PHONY: clean apidoc document
