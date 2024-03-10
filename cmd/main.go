package main

import (
	"employees/pkg/config"
	"employees/pkg/controller"
	"employees/pkg/oapi"
	"employees/pkg/postgres"
	"flag"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {
	genSpec := flag.Bool("gen", false, "Generate OpenAPI spec")
	argPort := flag.String("port", "8080", "Port to listen on")
	flag.Parse()

	if *genSpec {
		log.Println("Generating OpenAPI spec...")

		oapi.CreateOpenAPI()

		log.Println("Done.")
		return
	}

	Cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalln("Error loading configuration: ", err)
	}

	if *argPort != "" {
		Cfg.Port = *argPort
	}

	nomHost := os.Getenv("NOMAD_IP_httpemp")
	if nomHost != "" {
		Cfg.DBHost = nomHost
	}

	dbConn, err := postgres.ConnectDB(&Cfg)
	if err != nil {
		log.Fatalln("Error connecting to database: ", err)
	}

	srvAddr := ":" + Cfg.Port

	app := fiber.New()

	controller.CreateRoutes(app, dbConn)
	app.Listen(srvAddr)
}
