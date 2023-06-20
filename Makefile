clean:
	find . -name '*.go' -exec gofmt -w {} \;

run:
	go mod download
	go run .

db:
	docker build -t my-postgres .
	docker run -p 5432:5432 --name my-postgres-container -d my-postgres
