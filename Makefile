clean:
	find . -name '*.go' -exec gofmt -w {} \;
run:
	go run .