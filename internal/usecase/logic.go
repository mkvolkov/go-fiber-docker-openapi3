package usecase

import (
	"employees/internal/models"
)

type EmpUsecase struct {
	repo Employee
}

func NewUsecase(repoUC Employee) *EmpUsecase {
	return &EmpUsecase{
		repo: NewEmpUsecase(repoUC),
	}
}

func NewEmpUsecase(eRepo Employee) *EmpUsecase {
	return &EmpUsecase{repo: eRepo}
}

func (er *EmpUsecase) HireEmployee(emp *models.PEmployee) (models.FEmployee, error) {
	return er.repo.HireEmployee(emp)
}

func (er *EmpUsecase) FireEmployee(empID int) (models.FEmployee, error) {
	return er.repo.FireEmployee(empID)
}

func (er *EmpUsecase) GetVacationDays(empID int) (models.Vdays, error) {
	return er.repo.GetVacationDays(empID)
}

func (er *EmpUsecase) GetEmployeeByName(name string) ([]models.FEmployee, error) {
	return er.repo.GetEmployeeByName(name)
}
