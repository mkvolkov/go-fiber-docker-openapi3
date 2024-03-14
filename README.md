# Demo CRUD project in Golang with Nomad

## Techniques covered:

- Clean Architecture in Golang
- koanf (reading config files)
- Fiber framework
- Docker
- docker compose
- PostgreSQL (dockerized; with test data)
- pgx
- OpenAPI 3.0 (generating the specification)
- Nomad
- Consul
- HAProxy

### Run the project with Docker Compose:

```
docker compose up
```

Generate the OpenAPI 3.0 specification:

```
go run main.go -gen
```

### Run the project with Nomad, HAProxy and Consul

Start Nomad, using the provided configuration:

```
sudo nomad agent -config=nomad_cfg.hcl -dev
```

View Nomad in browser: http://localhost:4646

Start Consul:

```
consul agent -dev
```

View Consul in browser: http://localhost:8500

Start the project:

```
nomad run start_project.hcl
```

URL to check the project with curl, Insomnia, Postman, Advanced REST Client, etc.:

http://localhost:<port>/find/mikhail

(simple GET request without the body)

The port can be found in the Nomad UI (http://localhost:4646/ui/jobs)

Start HAProxy: use the provided configuration haproxy/haproxy.cfg

Put the file haproxy.cfg into the /etc/haproxy directory and then run:

```
systemctl restart haproxy
```

or

```
systemctl start haproxy
```

if HAProxy isn't still running.

Then, you can access the project with the following URL:

http://localhost/find/mikhail