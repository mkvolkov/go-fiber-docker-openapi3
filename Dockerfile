FROM golang:1.22.0-bookworm

WORKDIR /app

COPY . .
RUN go build -o /empl cmd/main.go

EXPOSE 8080

ENTRYPOINT ["/empl"]