package repo

import (
	"context"
	"employees/internal/models"
	"employees/internal/usecase"
	"fmt"
	"sync"

	"github.com/guregu/null"
	"github.com/jackc/pgx/v5"
)

type Repository struct {
	mtx  sync.Mutex
	conn *pgx.Conn

	usecase.Employee
}

func NewRepository(pgConn *pgx.Conn) *Repository {
	return &Repository{conn: pgConn}
}

func (r *Repository) HireEmployee(emp *models.PEmployee) (models.FEmployee, error) {
	sqlQuery := `
		INSERT INTO employees
		(name, phone, gender, age, email, address, vdays)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id, name, phone, gender, age, email, address, vdays`

	r.mtx.Lock()
	hired := r.conn.QueryRow(
		context.Background(),
		sqlQuery,
		emp.Name,
		emp.Phone,
		emp.Gender,
		emp.Age,
		emp.Email,
		emp.Address,
		emp.Vdays,
	)
	r.mtx.Unlock()

	var hiredEmp models.FEmployee
	var vdaysNull = null.NewInt(0, true)
	err := hired.Scan(
		&hiredEmp.ID,
		&hiredEmp.Name,
		&hiredEmp.Phone,
		&hiredEmp.Gender,
		&hiredEmp.Age,
		&hiredEmp.Email,
		&hiredEmp.Address,
		&vdaysNull,
	)
	if err != nil {
		return hiredEmp, err
	}

	if vdaysNull.IsZero() {
		hiredEmp.Vdays = 0
	} else {
		hiredEmp.Vdays = int(vdaysNull.Int64)
	}

	return hiredEmp, nil
}

func (e *Repository) FireEmployee(empID int) (models.FEmployee, error) {
	sqlQuery := `DELETE FROM employees WHERE id = $1
		RETURNING id, name, phone, gender, age, email, address, vdays`

	e.mtx.Lock()
	fired := e.conn.QueryRow(
		context.Background(),
		sqlQuery,
		empID,
	)
	e.mtx.Unlock()

	var firedEmp models.FEmployee
	var vdaysNull = null.NewInt(0, true)
	err := fired.Scan(
		&firedEmp.ID,
		&firedEmp.Name,
		&firedEmp.Phone,
		&firedEmp.Gender,
		&firedEmp.Age,
		&firedEmp.Email,
		&firedEmp.Address,
		&vdaysNull,
	)
	if err != nil {
		return firedEmp, err
	}

	if vdaysNull.IsZero() {
		firedEmp.Vdays = 0
	} else {
		firedEmp.Vdays = int(vdaysNull.Int64)
	}

	return firedEmp, nil
}

func (e *Repository) GetVacationDays(empID int) (models.Vdays, error) {
	sqlQuery := `SELECT vdays FROM employees WHERE id = $1`

	daysRes := e.conn.QueryRow(
		context.Background(),
		sqlQuery,
		empID,
	)

	var days null.Int
	var vdaysRet models.Vdays
	err := daysRes.Scan(&days)
	if err != nil {
		return vdaysRet, err
	}

	if days.IsZero() {
		vdaysRet.Vdays = 0
	} else {
		vdaysRet.Vdays = int(days.Int64)
	}

	return vdaysRet, nil
}

func (e *Repository) GetEmployeeByName(name string) ([]models.FEmployee, error) {
	sqlQuery := fmt.Sprintf("SELECT * FROM employees WHERE name ILIKE '%%%s%%'", name)

	emplsRes, err := e.conn.Query(
		context.Background(),
		sqlQuery,
	)

	if err != nil {
		return nil, err
	}
	defer emplsRes.Close()

	var empls []models.FEmployee
	for emplsRes.Next() {
		var emp models.FEmployee
		var vdaysNull = null.NewInt(0, true)
		err = emplsRes.Scan(
			&emp.ID,
			&emp.Name,
			&emp.Phone,
			&emp.Gender,
			&emp.Age,
			&emp.Email,
			&emp.Address,
			&vdaysNull,
		)
		if err != nil {
			return nil, err
		}

		if vdaysNull.IsZero() {
			emp.Vdays = 0
		} else {
			emp.Vdays = int(vdaysNull.Int64)
		}

		empls = append(empls, emp)
	}

	return empls, nil
}
