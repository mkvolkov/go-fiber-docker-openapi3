job "project" {
    datacenters = ["dc1"]
    type = "service"

    group "empls" {
        count = 1
        network {
            mode = "host"

            port "empls" {
                to = 8080
                static = 8080
            }
            
            port "postgr" {
                to = 5432
                static = 5432
            }
        }

        task "postgr" {
            driver = "docker"

            env {
                POSTGRES_USER = "postgres"
                POSTGRES_PASSWORD = "postpass"
                POSTGRES_DB = "emp_db"
            }

            config {
                network_mode = "host"
                image = "postgres"
                ports = ["postgr"]

                volumes = [
                    "/home/mike/tutorials/go-fiber-docker-openapi3/init:/docker-entrypoint-initdb.d"
                ]
            }
        }

        task "empls" {
            driver = "docker"

            config {
                network_mode = "host"
                image = "mkvolkov/employees:latest"
                ports = ["empls"]
            }

            resources {
                cores = 2
                memory = 2048
            }
        }
    }
}