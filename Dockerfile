FROM golang:1.21.6-bookworm

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download & go mod verify
ADD fserver ./fserver
ADD handlers ./handlers
ADD logic ./logic
ADD models ./models
ADD repository ./repository
COPY main.go cfg.yml ./

CMD [ "go", "run", "main.go" ]