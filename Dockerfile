FROM golang:1.22.0-bookworm

WORKDIR /app

COPY go.mod go.sum ./
ADD cmd ./cmd
ADD config ./config
ADD internal ./internal
ADD pkg ./pkg
RUN go mod tidy

CMD [ "go", "run", "cmd/app/main.go" ]