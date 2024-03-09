package oapi

import (
	"log"
	"net/http"
	"os"

	"github.com/swaggest/openapi-go"
	"github.com/swaggest/openapi-go/openapi3"
)

type ErrResponse string

type Employee struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Gender  string `json:"gender"`
	Age     int    `json:"age"`
	Email   string `json:"email"`
	Address string `json:"address"`
	Vdays   int    `json:"vdays"`
}

type EmpID struct {
	ID int `path:"id"`
}

type Vdays struct {
	Vdays int `json:"vdays"`
}

type EmpName struct {
	Name string `path:"name"`
}

func CreateOpenAPI() {
	reflector := openapi3.NewReflector()
	reflector.Spec = &openapi3.Spec{Openapi: "3.0.3"}
	reflector.Spec.Info.
		WithTitle("Employees API").
		WithVersion("2.0").
		WithDescription("Employees API, Demo Golang CRUD & OpenAPI 3.0")

	hireEmp, _ := reflector.NewOperationContext(http.MethodPost, "/hire")
	hireEmp.AddReqStructure(new(Employee))
	hireEmp.AddRespStructure(new(Employee), openapi.WithHTTPStatus(http.StatusOK))
	hireEmp.AddRespStructure(new(ErrResponse), openapi.WithHTTPStatus(http.StatusInternalServerError))
	reflector.AddOperation(hireEmp)

	fireEmp, _ := reflector.NewOperationContext(http.MethodDelete, "/fire/{id}")
	fireEmp.AddReqStructure(new(EmpID))
	fireEmp.AddRespStructure(new(Employee), openapi.WithHTTPStatus(http.StatusOK))
	fireEmp.AddRespStructure(new(ErrResponse), openapi.WithHTTPStatus(http.StatusInternalServerError))
	reflector.AddOperation(fireEmp)

	getVdays, _ := reflector.NewOperationContext(http.MethodGet, "/vdays/{id}")
	getVdays.AddReqStructure(new(EmpID))
	getVdays.AddRespStructure(new(Vdays), openapi.WithHTTPStatus(http.StatusOK))
	getVdays.AddRespStructure(new(ErrResponse), openapi.WithHTTPStatus(http.StatusInternalServerError))
	reflector.AddOperation(getVdays)

	findEmps, _ := reflector.NewOperationContext(http.MethodGet, "/find/{name}")
	findEmps.AddReqStructure(new(EmpName))
	findEmps.AddRespStructure(new([]Employee), openapi.WithHTTPStatus(http.StatusOK))
	findEmps.AddRespStructure(new(ErrResponse), openapi.WithHTTPStatus(http.StatusInternalServerError))
	reflector.AddOperation(findEmps)

	schema, err := reflector.Spec.MarshalYAML()
	if err != nil {
		log.Fatalln("Error marshalling OpenAPI specification:", err)
	}

	if _, err := os.Stat("./docs"); os.IsNotExist(err) {
		err := os.Mkdir("./docs", 0o777)
		if err != nil {
			log.Fatalln("Error creating directory:", err)
		}
	}
	os.WriteFile("./docs/openapi.yml", schema, 0o666)
}
