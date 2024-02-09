FROM golang:1.22.0-bookworm

WORKDIR /app

COPY go.mod go.sum ./
ADD cmd ./cmd
ADD handlers ./handlers
ADD logic ./logic
ADD models ./models
ADD pkg ./pkg
ADD repository ./repository
COPY cfg.yml ./
RUN go mod tidy

CMD [ "go", "run", "cmd/app/main.go" ]