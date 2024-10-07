##
test:
	go test -race -v ./...
# Make sure that you write in an API key before run the command
example:
	go run ./example/main.go

.PHONY: test example