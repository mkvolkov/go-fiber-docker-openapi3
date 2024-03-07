package main

import (
	"employees/config"
	httpv1 "employees/internal/controller/http/v1"
	"employees/internal/repo"
	"employees/internal/usecase"
	"employees/pkg/fserver"
	"employees/pkg/oapi"
	"employees/pkg/postgres"
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

	dbConn, err := postgres.ConnectDB(&Cfg.PgCfg)
	if err != nil {
		log.Fatalln("Error connecting to database: ", err)
	}

	fSrv := fserver.NewFServer()

	pgRepo := repo.NewRepository(dbConn)
	empUsecase := usecase.NewEmpUsecase(pgRepo)
	fHandlers := httpv1.NewHandler(empUsecase)

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
