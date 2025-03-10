package forms

import (
	"context"
	"fmt"
	"university-db-admin/internal/domain"
	"university-db-admin/internal/repository"
	"university-db-admin/pkg/validation"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func ShowMarksForm(content *fyne.Container, action string, r *repository.Repository) {
	content.Objects = nil

	switch action {
	case "Добавить":
		showAddMarksForm(content, r)
	case "Удалить":
		showDeleteMarksForm(content, r)
	case "Обновить":
		showUpdateMarksForm(content, r)
	case "Просмотреть":
		showMarksList(content, r)
	}

	content.Refresh()
}

func showAddMarksForm(content *fyne.Container, r *repository.Repository) {
	employeeEntry := widget.NewEntry()
	employeeEntry.SetPlaceHolder("ID преподавателя")

	studentEntry := widget.NewEntry()
	studentEntry.SetPlaceHolder("ID студента")

	subjectEntry := widget.NewEntry()
	subjectEntry.SetPlaceHolder("ID предмета")

	markEntry := widget.NewEntry()
	markEntry.SetPlaceHolder("Оценка")

	dateEntry := widget.NewEntry()
	dateEntry.SetPlaceHolder("Дата (YYYY-MM-DD)")

	submitButton := widget.NewButton("Добавить", func() {
		err := validation.ValidateEmptyStrings(
			employeeEntry.Text,
			studentEntry.Text,
			subjectEntry.Text,
			markEntry.Text,
			dateEntry.Text,
		)
		if err != nil {
			showResult(content, err, "")
			return
		}

		mark := domain.Mark{
			EmployeeID: parseUint64(employeeEntry.Text),
			StudentID:  parseUint64(studentEntry.Text),
			SubjectID:  parseUint64(subjectEntry.Text),
			Mark:       parseUint16(markEntry.Text),
			Date:       dateEntry.Text,
		}

		if err = validation.ValidateStruct(mark); err != nil {
			showResult(content, err, "")
			return
		}

		isT, err := r.IsEmployeeTeacher(parseUint64(employeeEntry.Text))
		if !isT {
			if err != nil {
				showResult(content, err, "")
			} else {
				showResult(content, err,
					"Ошибка: Указанный сотрудник не является преподавателем",
				)
			}
			return
		}

		err = r.Marks.Create(context.Background(), mark)
		showResult(content, err, "Оценка успешно добавлена")
	})

	form := container.NewVBox(
		widget.NewLabel("Добавление оценки"),
		employeeEntry,
		studentEntry,
		subjectEntry,
		markEntry,
		dateEntry,
		submitButton,
	)

	content.Add(form)
}

func showDeleteMarksForm(content *fyne.Container, r *repository.Repository) {
	idEntry := widget.NewEntry()
	idEntry.SetPlaceHolder("ID оценки")

	deleteButton := widget.NewButton("Удалить", func() {
		err := validation.ValidateEmptyStrings(idEntry.Text)
		if err != nil {
			showResult(content, err, "")
			return
		}

		id := parseUint64(idEntry.Text)
		err = validation.ValidatePositiveNumber(id)
		if err != nil {
			showResult(content, err, "")
			return
		}

		err = r.Marks.Delete(context.Background(), id)
		showResult(content, err, "Оценка удалена")
	})

	form := container.NewVBox(
		widget.NewLabel("Удаление оценки"),
		idEntry,
		deleteButton,
	)

	content.Add(form)
}

func showUpdateMarksForm(content *fyne.Container, r *repository.Repository) {
	idEntry := widget.NewEntry()
	idEntry.SetPlaceHolder("ID оценки")

	employeeEntry := widget.NewEntry()
	employeeEntry.SetPlaceHolder("Новый ID преподавателя")

	studentEntry := widget.NewEntry()
	studentEntry.SetPlaceHolder("Новый ID студента")

	subjectEntry := widget.NewEntry()
	subjectEntry.SetPlaceHolder("Новый ID предмета")

	markEntry := widget.NewEntry()
	markEntry.SetPlaceHolder("Новая оценка")

	dateEntry := widget.NewEntry()
	dateEntry.SetPlaceHolder("Новая дата (YYYY-MM-DD)")

	updateButton := widget.NewButton("Обновить", func() {
		err := validation.ValidateEmptyStrings(
			idEntry.Text,
			employeeEntry.Text,
			studentEntry.Text,
			subjectEntry.Text,
			markEntry.Text,
			dateEntry.Text,
		)
		if err != nil {
			showResult(content, err, "")
			return
		}

		mark := domain.Mark{
			ID:         parseUint64(idEntry.Text),
			EmployeeID: parseUint64(employeeEntry.Text),
			StudentID:  parseUint64(studentEntry.Text),
			SubjectID:  parseUint64(subjectEntry.Text),
			Mark:       parseUint16(markEntry.Text),
			Date:       dateEntry.Text,
		}

		if err = validation.ValidateStruct(mark); err != nil {
			showResult(content, err, "")
			return
		}

		isT, err := r.IsEmployeeTeacher(parseUint64(employeeEntry.Text))
		if !isT {
			if err != nil {
				showResult(content, err, "")
			} else {
				showResult(content, err,
					"Ошибка: Указанный сотрудник не является преподавателем",
				)
			}
			return
		}

		err = r.Marks.Update(context.Background(), mark.ID, mark)
		showResult(content, err, "Оценка обновлена")
	})

	form := container.NewVBox(
		widget.NewLabel("Обновление оценки"),
		idEntry,
		employeeEntry,
		studentEntry,
		subjectEntry,
		markEntry,
		dateEntry,
		updateButton,
	)

	content.Add(form)
}

func showMarksList(content *fyne.Container, r *repository.Repository) {
	content.Objects = nil

	headers := []string{
		"ID оценки",
		"ID преподавателя",
		"ID студента",
		"ID предмета",
		"Оценка",
		"Дата",
	}
	var data [][]string

	filterEntry := widget.NewEntry()
	filterEntry.SetPlaceHolder("Введите значение")

	filterOptions := map[string]uint8{
		"Все":              0,
		"ID":               1,
		"ID преподавателя": 2,
		"ID студента":      3,
		"ID предмета":      4,
		"Оценка":           5,
		"Дата":             6,
	}

	options := []string{
		"Все",
		"ID",
		"ID преподавателя",
		"ID студента",
		"ID предмета",
		"Оценка",
		"Дата",
	}
	var selectedField uint8
	filterSelect := widget.NewSelect(options, func(value string) {
		selectedField = filterOptions[value]

		if selectedField == 0 {
			filterEntry.SetText("")
			filterEntry.Disable()
		} else {
			filterEntry.Enable()
		}
	})

	applyFilterButton := widget.NewButton("Применить фильтр", func() {
		data = nil

		var (
			marks []domain.Mark
			mark  domain.Mark
			err   error
		)

		switch selectedField {
		case 0:
			marks, err = r.Marks.FindAll(context.Background())
		case 1:
			mark, err = r.Marks.FindOne(context.Background(), parseUint64(filterEntry.Text))
			if err == nil {
				marks = append(marks, mark)
			}
		case 2:
			marks, err = r.Marks.FindByEmployeeID(context.Background(), parseUint64(filterEntry.Text))
		case 3:
			marks, err = r.Marks.FindByStudentID(context.Background(), parseUint64(filterEntry.Text))
		case 4:
			marks, err = r.Marks.FindBySubjectID(context.Background(), parseUint64(filterEntry.Text))
		case 5:
			marks, err = r.Marks.FindByMark(context.Background(), parseUint16(filterEntry.Text))
		case 6:
			marks, err = r.Marks.FindByDate(context.Background(), filterEntry.Text)
		}

		if err != nil {
			showResult(content, err, "Ошибка при поиске")
			return
		}

		for _, m := range marks {
			data = append(data, []string{
				fmt.Sprintf("%d", m.ID),
				fmt.Sprintf("%d", m.EmployeeID),
				fmt.Sprintf("%d", m.StudentID),
				fmt.Sprintf("%d", m.SubjectID),
				fmt.Sprintf("%d", m.Mark),
				m.Date,
			})
		}

		content.Objects = content.Objects[:1] // Only filter widgets remain
		content.Add(updateTable(headers, data))
		content.Refresh()
	})

	marks, _ := r.Marks.FindAll(context.Background())
	for _, m := range marks {
		data = append(data, []string{
			fmt.Sprintf("%d", m.ID),
			fmt.Sprintf("%d", m.EmployeeID),
			fmt.Sprintf("%d", m.StudentID),
			fmt.Sprintf("%d", m.SubjectID),
			fmt.Sprintf("%d", m.Mark),
			m.Date,
		})
	}

	filterContainer := container.NewVBox(
		widget.NewLabel("Фильтрация оценок"),
		filterSelect,
		filterEntry,
		applyFilterButton,
	)

	content.Add(filterContainer)
	content.Add(updateTable(headers, data))
	content.Refresh()
}
