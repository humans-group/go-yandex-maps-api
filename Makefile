test:
	go test -race -coverprofile=coverage.txt -covermode=atomic -v ./...
# Make sure that you write in an API key before run the command
example:
	go run ./example/main.go