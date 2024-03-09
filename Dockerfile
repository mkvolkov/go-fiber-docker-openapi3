FROM golang:1.22.0-bookworm

WORKDIR /app

COPY go.mod go.sum .env ./
ADD cmd ./cmd
ADD pkg ./pkg
RUN go mod tidy

CMD [ "go", "run", "cmd/main.go" ]