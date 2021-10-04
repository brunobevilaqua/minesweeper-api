build:
	go build -o bin/cmd/minesweeper-api app/cmd/minesweeper-api/main.go
run:
	go run cmd/minesweeper-api/main.go

test:
	go test ./...

# docker
up:
	docker-compose up --build
down:
	docker-compose down --remove-orphans
