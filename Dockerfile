FROM golang

WORKDIR /app
COPY . .
RUN go build -o bin cmd/minesweeper-api/main.go
EXPOSE 40000
CMD ["./bin"]