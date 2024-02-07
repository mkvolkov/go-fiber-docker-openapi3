package logic

import (
	"employees/models"
	"employees/repository"
)

type Employee interface {
	HireEmployee(emp *models.PEmployee) (models.FEmployee, error)
	FireEmployee(empID int) (models.FEmployee, error)
	GetVacationDays(empID int) (models.Vdays, error)
	GetEmployeeByName(name string) ([]models.FEmployee, error)
}

type EmpLogic struct {
	Employee
}

func NewLogic(repo *repository.Repository) *EmpLogic {
	return &EmpLogic{
		Employee: NewEmpRepo(repo.Employee),
	}
}
