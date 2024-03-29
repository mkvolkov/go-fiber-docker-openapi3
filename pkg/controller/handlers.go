package controller

import (
	"employees/pkg/models"
	"strconv"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
)

func (h *Controller) hireEmployee(c *fiber.Ctx) error {
	var inputEmp models.PEmployee

	err := json.Unmarshal(c.Body(), &inputEmp)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("JSON unmarshal error")
	}

	hiredID, err := h.actions.HireEmployee(&inputEmp)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(hiredID)
}

func (h *Controller) fireEmployee(c *fiber.Ctx) error {
	ID := c.Params("id")

	fireID, err := strconv.Atoi(ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	firedID, err := h.actions.FireEmployee(fireID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(firedID)
}

func (h *Controller) getVacationDays(c *fiber.Ctx) error {
	ID := c.Params("id")

	getvID, err := strconv.Atoi(ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	vDays, err := h.actions.GetVacationDays(getvID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(vDays)
}

func (h *Controller) findEmployeeByName(c *fiber.Ctx) error {
	name := c.Params("name")

	empls, err := h.actions.GetEmployeeByName(name)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(empls)
}

func (h *Controller) healthCheck(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).SendString("OK")
}
