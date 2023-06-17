clean:
	find . -name '*.go' -exec gofmt -w {} \;

run:
	go mod download
	go run .