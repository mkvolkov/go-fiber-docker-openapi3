FROM golang:1.21.6-bookworm

WORKDIR /app

COPY go.mod go.sum ./
ADD fserver ./fserver
ADD handlers ./handlers
ADD logic ./logic
ADD models ./models
ADD repository ./repository
COPY main.go cfg.yml ./
RUN go mod tidy

CMD [ "go", "run", "main.go" ]