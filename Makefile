BINARY=minesweeper-api
VERSION=0.1.0

build:
	go build -o ${BINARY} cmd/main.go

run:
	@go run cmd/main.go

test:
	@go test ./...

# docker
up:
	docker-compose up --build
down:
	docker-compose down --remove-orphans
