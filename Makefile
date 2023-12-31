build:
	@go build -o gobank cmd/gobank/main.go

run: build docker-up
	@./bin/gobank

test:
	@go test -v ./...

docker-up:
	@docker-compose up -d

docker-down:
	@docker-compose down