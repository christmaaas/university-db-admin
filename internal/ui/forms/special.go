package forms

import (
	"context"
	"fmt"
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

	data, err := r.Employees.FindAllNamePassport(context.Background())
	if err != nil {
		showResult(content, "Ошибка: "+err.Error())
		return
	}

	rows := make([][]string, len(data))
	for i, dto := range data {
		rows[i] = []string{
			dto.Name,
			dto.Passport,
		}
	}
	content.Add(updateTable(headers, rows))
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
			showResult(content, "Ошибка: "+err.Error())
			return
		}

		id := parseUint64(idEntry.Text)
		err = validation.ValidatePositiveNumbers(id)
		if err != nil {
			showResult(content, "Ошибка: "+err.Error())
			return
		}

		data, err := r.Employees.FindNamePassportByID(context.Background(), id)
		if err != nil {
			showResult(content, "Ошибка: "+err.Error())
			return
		}

		rows := [][]string{{data.Name, data.Passport}}
		content.Objects = content.Objects[:1]
		content.Add(updateTable(headers, rows))
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

	data, err := r.Students.FindAllWithNoCurator(context.Background())
	if err != nil {
		showResult(content, "Ошибка: "+err.Error())
		return
	}

	rows := make([][]string, len(data))
	for i, dto := range data {
		rows[i] = []string{
			dto.Name,
			dto.Passport,
			fmt.Sprintf("%d", dto.GroupID),
		}
	}
	content.Add(updateTable(headers, rows))
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
			showResult(content, "Ошибка: "+err.Error())
			return
		}

		firstID := parseUint64(firstIDEntry.Text)
		secondID := parseUint64(secondIDEntry.Text)
		err = validation.ValidatePositiveNumbers(firstID, secondID)
		if err != nil {
			showResult(content, "Ошибка: "+err.Error())
			return
		}

		data, err := r.Employees.FindAllByPositions(context.Background(), firstID, secondID)
		if err != nil {
			showResult(content, "Ошибка: "+err.Error())
			return
		}

		rows := make([][]string, len(data))
		for i, dto := range data {
			rows[i] = []string{
				dto.Name,
			}
		}
		content.Objects = content.Objects[:1]
		content.Add(updateTable(headers, rows))
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
			showResult(content, "Ошибка: "+err.Error())
			return
		}

		id := parseUint64(subjectIDEntry.Text)
		mark := parseUint16(markEntry.Text)
		err = validation.ValidatePositiveNumbers(id, mark)
		if err != nil {
			showResult(content, "Ошибка: "+err.Error())
			return
		}

		data, err := r.Marks.FindAllBySubject(context.Background(), id, mark)
		if err != nil {
			showResult(content, "Ошибка: "+err.Error())
			return
		}

		rows := make([][]string, len(data))
		for i, dto := range data {
			rows[i] = []string{
				fmt.Sprintf("%d", dto.StudentID),
				fmt.Sprintf("%d", dto.Mark),
				dto.Date.Format("2006-01-02"),
			}
		}
		content.Objects = content.Objects[:1]
		content.Add(updateTable(headers, rows))
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
			showResult(content, "Ошибка: "+err.Error())
			return
		}

		data, err := r.Students.FindAllByMiddlename(context.Background(), seqEntry.Text)
		if err != nil {
			showResult(content, "Ошибка: "+err.Error())
			return
		}

		rows := make([][]string, len(data))
		for i, dto := range data {
			rows[i] = []string{
				dto.Name,
				dto.Passport,
			}
		}
		content.Objects = content.Objects[:1]
		content.Add(updateTable(headers, rows))
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

	data, err := r.Subjects.FindAllSorted(context.Background())
	if err != nil {
		showResult(content, "Ошибка: "+err.Error())
		return
	}

	rows := make([][]string, len(data))
	for i, dto := range data {
		rows[i] = []string{
			dto.Name,
		}
	}
	content.Add(updateTable(headers, rows))
}

func showSortedMarksForm(content *fyne.Container, r *repository.Repository) {
	headers := []string{
		"ID студента",
		"Оценка",
		"Дата",
	}

	data, err := r.Marks.FindAllSorted(context.Background())
	if err != nil {
		showResult(content, "Ошибка: "+err.Error())
		return
	}

	rows := make([][]string, len(data))
	for i, dto := range data {
		rows[i] = []string{
			fmt.Sprintf("%d", dto.StudentID),
			fmt.Sprintf("%d", dto.Mark),
			dto.Date.Format("2006-01-02"),
		}
	}
	content.Add(updateTable(headers, rows))
}

func showStudentGroupCombsForm(content *fyne.Container, r *repository.Repository) {
	headers := []string{
		"ФИО студента",
		"Номер группы",
	}

	data, err := r.Students.FindAllGroupCombs(context.Background())
	if err != nil {
		showResult(content, "Ошибка: "+err.Error())
		return
	}

	rows := make([][]string, len(data))
	for i, dto := range data {
		rows[i] = []string{
			dto.StudentName,
			fmt.Sprintf("%d", dto.GroupNumber),
		}
	}
	content.Add(updateTable(headers, rows))
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

	data, err := r.Lessons.FindSchedule(context.Background())
	if err != nil {
		showResult(content, "Ошибка: "+err.Error())
		return
	}

	rows := make([][]string, len(data))
	for i, dto := range data {
		rows[i] = []string{
			fmt.Sprintf("%d", dto.GroupNumber),
			dto.Subject,
			dto.LessonType,
			fmt.Sprintf("%d", dto.Room),
			fmt.Sprintf("%d", dto.Week),
			fmt.Sprintf("%d", dto.Weekday),
		}
	}
	content.Add(updateTable(headers, rows))
}

func showStudentsWithCuratorsForm(content *fyne.Container, r *repository.Repository) {
	headers := []string{
		"ФИО студента",
		"Паспорт студента",
		"ФИО куратора",
		"Паспорт куратора",
	}

	data, err := r.Students.FindAllWithCurators(context.Background())
	if err != nil {
		showResult(content, "Ошибка: "+err.Error())
		return
	}

	rows := make([][]string, len(data))
	for i, dto := range data {
		rows[i] = []string{
			dto.StudentName,
			dto.StudentPassport,
			dto.CuratorName,
			dto.CuratorPassport,
		}
	}
	content.Add(updateTable(headers, rows))
}

func showCuratorsWithStudentsForm(content *fyne.Container, r *repository.Repository) {
	headers := []string{
		"ФИО студента",
		"Паспорт студента",
		"ФИО куратора",
		"Паспорт куратора",
	}

	data, err := r.Students.FindWithAllCurators(context.Background())
	if err != nil {
		showResult(content, "Ошибка: "+err.Error())
		return
	}

	rows := make([][]string, len(data))
	for i, dto := range data {
		rows[i] = []string{
			dto.StudentName,
			dto.StudentPassport,
			dto.CuratorName,
			dto.CuratorPassport,
		}
	}
	content.Add(updateTable(headers, rows))
}

func showAllStudentCuratorPairsForm(content *fyne.Container, r *repository.Repository) {
	headers := []string{
		"ФИО студента",
		"Паспорт студента",
		"ФИО куратора",
		"Паспорт куратора",
	}

	data, err := r.Students.FindAllPairsWithCurator(context.Background())
	if err != nil {
		showResult(content, "Ошибка: "+err.Error())
		return
	}

	rows := make([][]string, len(data))
	for i, dto := range data {
		rows[i] = []string{
			dto.StudentName,
			dto.StudentPassport,
			dto.CuratorName,
			dto.CuratorPassport,
		}
	}
	content.Add(updateTable(headers, rows))
}

func showStudentsUppercaseWithLengthForm(content *fyne.Container, r *repository.Repository) {
	headers := []string{
		"ID студента",
		"ФИО в верхнем регистре",
		"Длина ФИО",
	}

	data, err := r.Students.FindAllUppercaseWithLength(context.Background())
	if err != nil {
		showResult(content, "Ошибка: "+err.Error())
		return
	}

	rows := make([][]string, len(data))
	for i, dto := range data {
		rows[i] = []string{
			fmt.Sprintf("%d", dto.ID),
			dto.UppercaseName,
			fmt.Sprintf("%d", dto.NameLength),
		}
	}
	content.Add(updateTable(headers, rows))
}
