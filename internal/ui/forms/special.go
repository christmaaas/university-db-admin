package forms

import (
	"context"
	"university-db-admin/internal/repository"

	"fyne.io/fyne/v2"
)

func ShowSpecialQueryForm(content *fyne.Container, action int, r *repository.Repository) {
	content.Objects = nil

	switch action {
	case 0:
		showEmployeesInfoForm(content, r)
	case 1:
		// TODO
	case 2:
		// TODO
	case 3:
		// TODO
	case 4:
		// TODO
	case 5:
		// TODO
	case 6:
		// TODO
	case 7:
		// TODO
	case 8:
		// TODO
	case 9:
		showLessonsScheduleForm(content, r)
	case 10:
		// TODO
	case 11:
		// TODO
	case 12:
		// TODO
	case 13:
		// TODO
	}

	content.Refresh()
}

func showLessonsScheduleForm(content *fyne.Container, r *repository.Repository) {
	content.Objects = nil

	headers := []string{
		"Номер группы",
		"Название предмета",
		"Тип занятия",
		"Аудитория",
		"Неделя",
		"День недели",
	}

	data, err := r.Special.GetScheduleByGroups(context.Background())
	if err != nil {
		showResult(content, err, "Ошибка при поиске")
		return
	}

	content.Add(updateTable(headers, data))
	content.Refresh()
}

func showEmployeesInfoForm(content *fyne.Container, r *repository.Repository) {
	content.Objects = nil

	headers := []string{
		"ФИО",
		"Номер паспорта",
	}

	data, err := r.Special.GetAllEmployeesInfo(context.Background())
	if err != nil {
		showResult(content, err, "Ошибка при поиске")
		return
	}

	content.Add(updateTable(headers, data))
	content.Refresh()
}
