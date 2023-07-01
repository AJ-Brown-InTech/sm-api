clean:
	sudo find . -name '*.go' -exec gofmt -w {} \;

run:
	sudo go mod download
	sudo go run .

run-d:
	docker compose up -d --force-recreate --quiet-pull 