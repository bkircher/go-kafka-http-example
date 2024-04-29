default: build

build:
	go build

build-debug:
	go build -gcflags="all=-N -l"

lint: build
	staticcheck ./...

clean:
	go clean

.PHONY: build build-debug lint clean
