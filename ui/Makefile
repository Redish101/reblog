all: clean build

clean:
	rm -rf dist

build:
	pnpm build
	rm -rf ../internal/ui/dist
	cp -r dist ../internal/ui/dist

install-dev:
	pnpm install