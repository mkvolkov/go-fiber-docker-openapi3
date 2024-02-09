package main

import (
	"employees/config"
	"employees/handlers"
	"employees/logic"
	"employees/oapi"
	"employees/pkg/fserver"
	"employees/repository"
	"flag"
	"log"
	"os"
	"os/signal"
)

func main() {
	genSpec := flag.Bool("gen", false, "generate openapi3 specification")
	flag.Parse()

	if *genSpec {
		log.Println("Generating OpenAPI 3 Specification...")

		oapi.CreateOpenAPI()

		log.Println("Done.")
		return
	}

	Cfg, err := config.ReadConfig("./config/cfg.yml")
	if err != nil {
		log.Fatalln("Error loading configuration: ", err)
	}

	srvAddr := Cfg.Srv.Host + ":" + Cfg.Srv.Port

	dbConn, err := repository.ConnectDB(&Cfg.PgCfg)
	if err != nil {
		log.Fatalln("Error connecting to database: ", err)
	}

	fSrv := fserver.NewFServer()

	pgRepo := repository.NewRepository(dbConn)
	empLogic := logic.NewLogic(pgRepo)
	fHandlers := handlers.NewHandler(empLogic)

	fHandlers.CreateRoutes(fSrv.FiberApp)

	doneCh := make(chan os.Signal, 1)
	signal.Notify(doneCh, os.Interrupt)
	go func() {
		<-doneCh
		log.Println("Interrupt signal received, graceful shutdown...")
		fSrv.Shutdown()
	}()

	fSrv.Run(srvAddr)
}
