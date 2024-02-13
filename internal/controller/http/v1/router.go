package httpv1

import (
	"employees/logic"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	actions logic.Employee
}

func NewHandler(emplogic logic.Employee) *Handler {
	return &Handler{actions: emplogic}
}

func (h *Handler) CreateRoutes(app *fiber.App) {
	app.Put("/hire", h.hireEmployee)
	app.Delete("/fire/:id", h.fireEmployee)
	app.Get("/vdays/:id", h.getVacationDays)
	app.Get("/find/:name", h.findEmployeeByName)
}
