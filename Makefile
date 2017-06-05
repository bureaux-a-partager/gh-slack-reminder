GOFILES = $(shell find . -name '*.go' -not -path './vendor/*')
GOPACKAGES = $(shell go list ./...  | grep -v /vendor/)

build-linux :
	mkdir -p build/linux/
	env GOOS=linux GOARCH=386 go build -o build/linux/gh-slack-reminder

build-macos :
	mkdir -p build/macos/
	env GOOS=darwin GOARCH=386 go build -o build/macos/gh-slack-reminder

build-all : build-linux build-macos

test: test-all

test-all :
	@go test -v $(GOPACKAGES)
