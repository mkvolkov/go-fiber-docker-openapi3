job "project" {
    datacenters = ["dc1"]
    type = "service"

    group "empls" {
        count = 1

        network {
            mode = "host"

            port "empls" {
                to = 8080
            }
        }

        task "empls" {
            driver = "docker"

            config {
                network_mode = "host"
                image = "mkvolkov/employees:3.0.0"
                ports = ["empls"]
            }

            resources {
                cores = 1
                memory = 1024
            }
        }
    }

    group "pgsql" {
        count = 1

        network {
            mode = "host"

            port "pgsql" {
                to = 5432
            }
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

            resources {
                cores = 1
                memory = 1024
            }
        }
    }
}