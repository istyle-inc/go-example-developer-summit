.PHONY: build
build:
	dep ensure
	go build -o goproducer
