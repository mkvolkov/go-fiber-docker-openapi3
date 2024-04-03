job "project" {
    datacenters = ["dc1"]
    type = "service"

    group "empls" {
        count = 4

        network {
            port "httpemp" {}
        }

        task "empls" {
            driver = "docker"

            config {
                network_mode = "host"
                image = "mkvolkov/employees:3.1.1"
                ports = ["httpemp"]

                args = [
                    "-port", "${NOMAD_PORT_httpemp}",
                ]
            }

            service {
                name = "employees"
                port = "httpemp"
                check {
                    type = "http"
                    path = "/health"
                    interval = "10s"
                    timeout = "2s"
                }
            }
        }
    }

    group "pgsql" {
        count = 1

        network {
            port "pgsql" {}
        }

        task "pgsql" {
            driver = "docker"

            env {
                POSTGRES_USER = "mike"
                POSTGRES_PASSWORD = "postpass"
                POSTGRES_DB = "emp_db"
            }

            config {
                network_mode = "host"
                image = "postgres"
                ports = ["pgsql"]

                volumes = [
                    "/home/mike/tutorials/go-fiber-docker-openapi3/init:/docker-entrypoint-initdb.d"
                ]
            }
        }
    }
}