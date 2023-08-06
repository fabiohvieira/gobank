build:
	@go build -o bin/gobank

run: build docker-up
	@./bin/gobank

test:
	@go test -v ./...

docker-up:
	@docker-compose up -d

docker-down:
	@docker-compose down