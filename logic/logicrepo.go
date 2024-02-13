package logic

import (
	"employees/internal/models"
	"employees/repository"
)

type EmpRepo struct {
	repo repository.Employee
}

func NewEmpRepo(eRepo repository.Employee) *EmpRepo {
	return &EmpRepo{repo: eRepo}
}

func (er *EmpRepo) HireEmployee(emp *models.PEmployee) (models.FEmployee, error) {
	return er.repo.HireEmployee(emp)
}

func (er *EmpRepo) FireEmployee(empID int) (models.FEmployee, error) {
	return er.repo.FireEmployee(empID)
}

func (er *EmpRepo) GetVacationDays(empID int) (models.Vdays, error) {
	return er.repo.GetVacationDays(empID)
}

func (er *EmpRepo) GetEmployeeByName(name string) ([]models.FEmployee, error) {
	return er.repo.GetEmployeeByName(name)
}
