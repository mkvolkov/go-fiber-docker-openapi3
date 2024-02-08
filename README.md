# Demo CRUD project in Golang

## Techniques covered:

- Clean Architecture in Golang
- koanf (reading config files)
- Fiber framework
- Docker
- docker compose
- PostgreSQL (dockerized; with test data)
- pgx
- OpenAPI 3.0 (generating the specification)

### Run the project:

Basic usage:

```
docker compose up
```

Generate the OpenAPI 3.0 specification:

```
go run main.go -gen
```