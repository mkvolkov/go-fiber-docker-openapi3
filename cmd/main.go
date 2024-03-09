package main

import (
	"employees/pkg/config"
	"employees/pkg/controller"
	"employees/pkg/postgres"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {
	Cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalln("Error loading configuration: ", err)
	}

	nomHost := os.Getenv("NOMAD_IP_empls")
	if nomHost != "" {
		Cfg.DBHost = nomHost
	}

	fmt.Println("NOMAD_IP_empls: ", nomHost)

	dbConn, err := postgres.ConnectDB(&Cfg)
	if err != nil {
		log.Fatalln("Error connecting to database: ", err)
	}

	srvAddr := ":" + Cfg.Port

	app := fiber.New()

	controller.CreateRoutes(app, dbConn)
	app.Listen(srvAddr)
}
