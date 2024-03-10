package controller

import (
	"employees/pkg/interfaces"
	"employees/pkg/repo"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5"
)

type Controller struct {
	actions interfaces.Employee
}

func NewController(db *pgx.Conn) *Controller {
	return &Controller{
		actions: repo.NewRepository(db),
	}
}

func CreateRoutes(app *fiber.App, db *pgx.Conn) {
	c := NewController(db)

	app.Post("/hire", c.hireEmployee)
	app.Delete("/fire/:id", c.fireEmployee)
	app.Get("/vdays/:id", c.getVacationDays)
	app.Get("/find/:name", c.findEmployeeByName)
	app.Get("/health", c.healthCheck)
}
