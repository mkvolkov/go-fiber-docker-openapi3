package interfaces

import "employees/pkg/models"

type Employee interface {
	HireEmployee(emp *models.PEmployee) (models.FEmployee, error)
	FireEmployee(empID int) (models.FEmployee, error)
	GetVacationDays(empID int) (models.Vdays, error)
	GetEmployeeByName(name string) ([]models.FEmployee, error)
}
