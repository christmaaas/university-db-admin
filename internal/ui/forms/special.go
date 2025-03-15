package forms

import (
	"context"
	"university-db-admin/internal/repository"
	"university-db-admin/pkg/validation"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func ShowSpecialQueryForm(content *fyne.Container, action int, r *repository.Repository) {
	content.Objects = nil

	switch action {
	case 0:
		showEmployeesInfoForm(content, r)
	case 1:
		showEmployeeInfoForm(content, r)
	case 2:
		showStudentsNoCuratorInfoForm(content, r)
	case 3:
		showEmployeesInfoByPositionsForm(content, r)
	case 4:
		showMarksInfoBySubjectForm(content, r)
	case 5:
		showStudnetsInfoByMiddlenameForm(content, r)
	case 6:
		showSortedSubjectsInfoForm(content, r)
	case 7:
		showSortedMarksInfoForm(content, r)
	case 8:
		showAllStudentsGroupsCombsForm(content, r)
	case 9:
		showLessonsScheduleForm(content, r)
	case 10:
		showAllStudentsWithCuratorsForm(content, r)
	case 11:
		showAllCuratorsWithStudentsForm(content, r)
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

func showEmployeeInfoForm(content *fyne.Container, r *repository.Repository) {
	content.Objects = nil

	headers := []string{
		"ФИО",
		"Номер паспорта",
	}

	idEntry := widget.NewEntry()
	idEntry.SetPlaceHolder("ID Сотрудника")

	submitButton := widget.NewButton("Применить", func() {
		err := validation.ValidateEmptyStrings(idEntry.Text)
		if err != nil {
			showResult(content, err, "")
			return
		}

		id := parseUint64(idEntry.Text)
		err = validation.ValidatePositiveNumbers(id)
		if err != nil {
			showResult(content, err, "")
			return
		}

		data, err := r.Special.GetAllEmployeesInfoByID(context.Background(), id)
		if err != nil {
			showResult(content, err, "Ошибка при поиске")
			return
		}

		content.Objects = content.Objects[:1]
		content.Add(updateTable(headers, data))
		content.Refresh()
	})

	form := container.NewVBox(
		widget.NewLabel("Выбор сотрудника"),
		idEntry,
		submitButton,
	)

	content.Add(form)
	content.Refresh()
}

func showStudentsNoCuratorInfoForm(content *fyne.Container, r *repository.Repository) {
	content.Objects = nil

	headers := []string{
		"ФИО",
		"Номер паспорта",
		"ID группы",
	}

	data, err := r.Special.GetStudentsNoCuratorInfo(context.Background())
	if err != nil {
		showResult(content, err, "Ошибка при поиске")
		return
	}

	content.Add(updateTable(headers, data))
	content.Refresh()
}

func showEmployeesInfoByPositionsForm(content *fyne.Container, r *repository.Repository) {
	content.Objects = nil

	headers := []string{"ФИО"}

	firstIDEntry := widget.NewEntry()
	firstIDEntry.SetPlaceHolder("ID должности")

	secondIDEntry := widget.NewEntry()
	secondIDEntry.SetPlaceHolder("ID должности")

	submitButton := widget.NewButton("Применить", func() {
		err := validation.ValidateEmptyStrings(firstIDEntry.Text, secondIDEntry.Text)
		if err != nil {
			showResult(content, err, "")
			return
		}

		firstID := parseUint64(firstIDEntry.Text)
		secondID := parseUint64(secondIDEntry.Text)
		err = validation.ValidatePositiveNumbers(firstID, secondID)
		if err != nil {
			showResult(content, err, "")
			return
		}

		data, err := r.Special.GetEmployeesInfoByPositionsID(context.Background(), firstID, secondID)
		if err != nil {
			showResult(content, err, "Ошибка при поиске")
			return
		}

		content.Objects = content.Objects[:1]
		content.Add(updateTable(headers, data))
		content.Refresh()
	})

	form := container.NewVBox(
		widget.NewLabel("Выбор должностей"),
		firstIDEntry,
		secondIDEntry,
		submitButton,
	)

	content.Add(form)
	content.Refresh()
}

func showMarksInfoBySubjectForm(content *fyne.Container, r *repository.Repository) {
	content.Objects = nil

	headers := []string{
		"ID студента",
		"Оценка",
		"Дата",
	}

	subjectIDEntry := widget.NewEntry()
	subjectIDEntry.SetPlaceHolder("ID предмета")

	markEntry := widget.NewEntry()
	markEntry.SetPlaceHolder("Оценка")

	submitButton := widget.NewButton("Применить", func() {
		err := validation.ValidateEmptyStrings(subjectIDEntry.Text, markEntry.Text)
		if err != nil {
			showResult(content, err, "")
			return
		}

		id := parseUint64(subjectIDEntry.Text)
		mark := parseUint16(markEntry.Text)
		err = validation.ValidatePositiveNumbers(id, mark)
		if err != nil {
			showResult(content, err, "")
			return
		}

		data, err := r.Special.GetMarksInfoBySubjectID(context.Background(), id, mark)
		if err != nil {
			showResult(content, err, "Ошибка при поиске")
			return
		}

		content.Objects = content.Objects[:1]
		content.Add(updateTable(headers, data))
		content.Refresh()
	})

	form := container.NewVBox(
		widget.NewLabel("Выбор предмета и оценки"),
		subjectIDEntry,
		markEntry,
		submitButton,
	)

	content.Add(form)
	content.Refresh()
}

func showStudnetsInfoByMiddlenameForm(content *fyne.Container, r *repository.Repository) {
	content.Objects = nil

	headers := []string{
		"ФИО",
		"Номер паспорта",
	}

	seqEntry := widget.NewEntry()
	seqEntry.SetPlaceHolder("Последовательность")

	submitButton := widget.NewButton("Применить", func() {
		err := validation.ValidateEmptyStrings(seqEntry.Text)
		if err != nil {
			showResult(content, err, "")
			return
		}

		data, err := r.Special.GetStudentsInfoByMiddlename(context.Background(), seqEntry.Text)
		if err != nil {
			showResult(content, err, "Ошибка при поиске")
			return
		}

		content.Objects = content.Objects[:1]
		content.Add(updateTable(headers, data))
		content.Refresh()
	})

	form := container.NewVBox(
		widget.NewLabel("Выбор последовательности"),
		seqEntry,
		submitButton,
	)

	content.Add(form)
	content.Refresh()
}

func showSortedSubjectsInfoForm(content *fyne.Container, r *repository.Repository) {
	content.Objects = nil

	headers := []string{"Название"}

	data, err := r.Special.GetSortedSubjectsInfo(context.Background())
	if err != nil {
		showResult(content, err, "Ошибка при поиске")
		return
	}

	content.Add(updateTable(headers, data))
	content.Refresh()
}

func showSortedMarksInfoForm(content *fyne.Container, r *repository.Repository) {
	content.Objects = nil

	headers := []string{
		"ID студента",
		"Оценка",
		"Дата",
	}

	data, err := r.Special.GetSortedMarksInfo(context.Background())
	if err != nil {
		showResult(content, err, "Ошибка при поиске")
		return
	}

	content.Add(updateTable(headers, data))
	content.Refresh()
}

func showAllStudentsGroupsCombsForm(content *fyne.Container, r *repository.Repository) {
	content.Objects = nil

	headers := []string{
		"ФИО студента",
		"Номер группы",
	}

	data, err := r.Special.GetAllStudentsGroupsCombs(context.Background())
	if err != nil {
		showResult(content, err, "Ошибка при поиске")
		return
	}

	content.Add(updateTable(headers, data))
	content.Refresh()
}

func showAllStudentsWithCuratorsForm(content *fyne.Container, r *repository.Repository) {
	content.Objects = nil

	headers := []string{
		"ФИО студента",
		"Паспорт студента",
		"ФИО куратора",
		"Паспорт куратора",
	}

	data, err := r.Special.GetAllStudentsWithCurators(context.Background())
	if err != nil {
		showResult(content, err, "Ошибка при поиске")
		return
	}

	content.Add(updateTable(headers, data))
	content.Refresh()
}

func showAllCuratorsWithStudentsForm(content *fyne.Container, r *repository.Repository) {
	content.Objects = nil

	headers := []string{
		"ФИО студента",
		"Паспорт студента",
		"ФИО куратора",
		"Паспорт куратора",
	}

	data, err := r.Special.GetAllCuratorsWithStudents(context.Background())
	if err != nil {
		showResult(content, err, "Ошибка при поиске")
		return
	}

	content.Add(updateTable(headers, data))
	content.Refresh()
}
