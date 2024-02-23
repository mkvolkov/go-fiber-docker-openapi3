package httpv1

import (
	"employees/internal/usecase"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	actions usecase.Employee
}

func NewHandler(emplogic usecase.Employee) *Handler {
	return &Handler{actions: emplogic}
}

func (h *Handler) CreateRoutes(app *fiber.App) {
	app.Post("/hire", h.hireEmployee)
	app.Delete("/fire/:id", h.fireEmployee)
	app.Get("/vdays/:id", h.getVacationDays)
	app.Get("/find/:name", h.findEmployeeByName)
}
