FROM golang:1.21-alpine

WORKDIR /app

RUN go install github.com/cosmtrek/air@latest

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o ./tmp/main ./cmd/main.go

EXPOSE 8080

CMD ["air"]