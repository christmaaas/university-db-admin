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
		showEmployeesForm(content, r)
	case 1:
		showEmployeeForm(content, r)
	case 2:
		showStudentsNoCuratorForm(content, r)
	case 3:
		showEmployeesByPositionsForm(content, r)
	case 4:
		showMarksBySubjectForm(content, r)
	case 5:
		showStudentsByMiddlenameForm(content, r)
	case 6:
		showSortedSubjectsForm(content, r)
	case 7:
		showSortedMarksForm(content, r)
	case 8:
		showStudentGroupCombsForm(content, r)
	case 9:
		showLessonsScheduleForm(content, r)
	case 10:
		showStudentsWithCuratorsForm(content, r)
	case 11:
		showCuratorsWithStudentsForm(content, r)
	case 12:
		showAllStudentCuratorPairsForm(content, r)
	case 13:
		showStudentsUppercaseWithLengthForm(content, r)
	}

	content.Refresh()
}

func showEmployeesForm(content *fyne.Container, r *repository.Repository) {
	headers := []string{
		"ФИО",
		"Номер паспорта",
	}

	data, err := r.Special.GetAllEmployees(context.Background())
	if err != nil {
		showResult(content, err, "Ошибка при поиске")
		return
	}

	content.Add(updateTable(headers, data))
}

func showEmployeeForm(content *fyne.Container, r *repository.Repository) {
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

		data, err := r.Special.GetEmployeeByID(context.Background(), id)
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
}

func showStudentsNoCuratorForm(content *fyne.Container, r *repository.Repository) {
	headers := []string{
		"ФИО",
		"Номер паспорта",
		"ID группы",
	}

	data, err := r.Special.GetStudentsNoCurator(context.Background())
	if err != nil {
		showResult(content, err, "Ошибка при поиске")
		return
	}

	content.Add(updateTable(headers, data))
}

func showEmployeesByPositionsForm(content *fyne.Container, r *repository.Repository) {
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

		data, err := r.Special.GetEmployeesByPositions(context.Background(), firstID, secondID)
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
}

func showMarksBySubjectForm(content *fyne.Container, r *repository.Repository) {
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

		data, err := r.Special.GetMarksBySubject(context.Background(), id, mark)
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
}

func showStudentsByMiddlenameForm(content *fyne.Container, r *repository.Repository) {
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

		data, err := r.Special.GetStudentsByMiddlename(context.Background(), seqEntry.Text)
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
}

func showSortedSubjectsForm(content *fyne.Container, r *repository.Repository) {
	headers := []string{"Название"}

	data, err := r.Special.GetSortedSubjects(context.Background())
	if err != nil {
		showResult(content, err, "Ошибка при поиске")
		return
	}

	content.Add(updateTable(headers, data))
}

func showSortedMarksForm(content *fyne.Container, r *repository.Repository) {
	headers := []string{
		"ID студента",
		"Оценка",
		"Дата",
	}

	data, err := r.Special.GetSortedMarks(context.Background())
	if err != nil {
		showResult(content, err, "Ошибка при поиске")
		return
	}

	content.Add(updateTable(headers, data))
}

func showStudentGroupCombsForm(content *fyne.Container, r *repository.Repository) {
	headers := []string{
		"ФИО студента",
		"Номер группы",
	}

	data, err := r.Special.GetStudentGroupCombs(context.Background())
	if err != nil {
		showResult(content, err, "Ошибка при поиске")
		return
	}

	content.Add(updateTable(headers, data))
}

func showLessonsScheduleForm(content *fyne.Container, r *repository.Repository) {
	headers := []string{
		"Номер группы",
		"Название предмета",
		"Тип занятия",
		"Аудитория",
		"Неделя",
		"День недели",
	}

	data, err := r.Special.GetLessonsSchedule(context.Background())
	if err != nil {
		showResult(content, err, "Ошибка при поиске")
		return
	}

	content.Add(updateTable(headers, data))
}

func showStudentsWithCuratorsForm(content *fyne.Container, r *repository.Repository) {
	headers := []string{
		"ФИО студента",
		"Паспорт студента",
		"ФИО куратора",
		"Паспорт куратора",
	}

	data, err := r.Special.GetStudentsWithCurators(context.Background())
	if err != nil {
		showResult(content, err, "Ошибка при поиске")
		return
	}

	content.Add(updateTable(headers, data))
}

func showCuratorsWithStudentsForm(content *fyne.Container, r *repository.Repository) {
	headers := []string{
		"ФИО студента",
		"Паспорт студента",
		"ФИО куратора",
		"Паспорт куратора",
	}

	data, err := r.Special.GetCuratorsWithStudents(context.Background())
	if err != nil {
		showResult(content, err, "Ошибка при поиске")
		return
	}

	content.Add(updateTable(headers, data))
}

func showAllStudentCuratorPairsForm(content *fyne.Container, r *repository.Repository) {
	headers := []string{
		"ФИО студента",
		"Паспорт студента",
		"ФИО куратора",
		"Паспорт куратора",
	}

	data, err := r.Special.GetAllStudentCuratorPairs(context.Background())
	if err != nil {
		showResult(content, err, "Ошибка при поиске")
		return
	}

	content.Add(updateTable(headers, data))
}

func showStudentsUppercaseWithLengthForm(content *fyne.Container, r *repository.Repository) {
	headers := []string{
		"ID студента",
		"ФИО в верхнем регистре",
		"Длина ФИО",
	}

	data, err := r.Special.GetStudentsUppercaseWithLength(context.Background())
	if err != nil {
		showResult(content, err, "Ошибка при поиске")
		return
	}

	content.Add(updateTable(headers, data))
}
