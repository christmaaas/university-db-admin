package repository

/*
	here will be special requests to database
	which the developer defined
*/

import (
	"context"
	"university-db-admin/internal/domain"
)

func (r *Repository) IsEmployeeTeacher(id uint64) (bool, error) {
	var (
		emp domain.Employee
		pos domain.Position
		err error
	)
	const teacherName = "Преподаватель"

	emp, err = r.Employees.FindOne(context.Background(), id)
	if err == nil {
		pos, err = r.Positions.FindOne(context.Background(), emp.PositionID)
		if err == nil {
			if pos.Name == teacherName {
				return true, nil
			}
		}
	}
	return false, err
}
