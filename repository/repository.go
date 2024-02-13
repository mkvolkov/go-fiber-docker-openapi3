package repository

import (
	"employees/internal/models"

	"github.com/jackc/pgx/v5"
)

type Employee interface {
	HireEmployee(emp *models.PEmployee) (models.FEmployee, error)
	FireEmployee(empID int) (models.FEmployee, error)
	GetVacationDays(empID int) (models.Vdays, error)
	GetEmployeeByName(name string) ([]models.FEmployee, error)
}

type Repository struct {
	Employee
}

func NewRepository(db *pgx.Conn) *Repository {
	return &Repository{Employee: NewEmpPostgres(db)}
}
