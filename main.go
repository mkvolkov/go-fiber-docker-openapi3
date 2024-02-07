package main

import (
	"employees/fserver"
	"employees/handlers"
	"employees/logic"
	"employees/oapi"
	"employees/repository"
	"flag"
	"log"

	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
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

	k := koanf.New(".")
	err := k.Load(file.Provider("./cfg.yml"), yaml.Parser())
	if err != nil {
		log.Fatalln("Error loading configuration: ", err)
	}

	srvAddr := k.String("server.host") + ":" + k.String("server.port")

	pgCfg := repository.PgCfg{}
	pgCfg.Host = k.String("postgres.host")
	pgCfg.Port = k.String("postgres.port")
	pgCfg.Username = k.String("postgres.username")
	pgCfg.Password = k.String("postgres.password")
	pgCfg.DBName = k.String("postgres.dbname")

	dbConn, err := repository.ConnectDB(&pgCfg)
	if err != nil {
		log.Fatalln("Error connecting to database: ", err)
	}

	fSrv := fserver.NewFServer()

	pgRepo := repository.NewRepository(dbConn)
	empLogic := logic.NewLogic(pgRepo)
	fHandlers := handlers.NewHandler(empLogic)

	fHandlers.CreateRoutes(fSrv.FiberApp)
	fSrv.Run(srvAddr)
}
