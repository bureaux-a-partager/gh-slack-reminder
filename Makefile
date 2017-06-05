build-linux :
	mkdir -p build/linux/
	env GOOS=linux GOARCH=386 go build -o build/linux/gh-slack-reminder

build-macos :
	mkdir -p build/macos/
	env GOOS=darwin GOARCH=386 go build -o build/macos/gh-slack-reminder

build : build-linux build-macos

test :
	go test
