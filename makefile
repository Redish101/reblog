all: clean backend

backend:
	go build -o bin/reblog -ldflags "-w -s" -gcflags "-N -l"

gen:
	go run cmd/gen.go

apidoc:
	swag fmt
	swag init -g server/server.go --parseDependency --parseInternal 
	redocly build-docs docs/swagger.yaml -o apidoc/index.html

dev:
	go build -o bin/reblog-dev
	./bin/reblog-dev

clean:
	rm -rf bin/*

.PHONY: clean apidoc