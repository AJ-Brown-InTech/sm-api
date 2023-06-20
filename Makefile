clean:
	find . -name '*.go' -exec gofmt -w {} \;

run:
	go mod download
	go run .

run-docker:
	docker-compose up -d --force-recreate --quiet-pull 