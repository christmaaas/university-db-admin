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

func ShowMarksForm(content *fyne.Container, action int, r *repository.Repository) {
	content.Objects = nil

	switch action {
	case 0:
		showAddMarksForm(content, r)
	case 1:
		showDeleteMarksForm(content, r)
	case 2:
		showUpdateMarksForm(content, r)
	case 3:
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
			showResult(content, "Ошибка: "+err.Error())
			return
		}

		mark := domain.Mark{
			EmployeeID: parseUint64(employeeEntry.Text),
			StudentID:  parseUint64(studentEntry.Text),
			SubjectID:  parseUint64(subjectEntry.Text),
			Mark:       parseUint16(markEntry.Text),
			Date:       parseDate(dateEntry.Text),
		}

		if err = validation.ValidateStruct(mark); err != nil {
			showResult(content, "Ошибка: "+err.Error())
			return
		}

		res, err := r.Special.IsTeacher(context.Background(), parseUint64(employeeEntry.Text))
		if !res.IsTeacher {
			if err != nil {
				showResult(content, "Ошибка: "+err.Error())
			} else {
				showResult(content, "Ошибка: указанный сотрудник не является преподавателем")
			}
			return
		}

		if err = r.Marks.Create(context.Background(), mark); err != nil {
			showResult(content, "Ошибка: "+err.Error())
			return
		}
		showResult(content, "Оценка успешно добавлена")
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
			showResult(content, "Ошибка: "+err.Error())
			return
		}

		id := parseUint64(idEntry.Text)
		err = validation.ValidatePositiveNumbers(id)
		if err != nil {
			showResult(content, "Ошибка: "+err.Error())
			return
		}

		if err = r.Marks.Delete(context.Background(), id); err != nil {
			showResult(content, "Ошибка: "+err.Error())
			return
		}
		showResult(content, "Оценка удалена")
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
			showResult(content, "Ошибка: "+err.Error())
			return
		}

		mark := domain.Mark{
			ID:         parseUint64(idEntry.Text),
			EmployeeID: parseUint64(employeeEntry.Text),
			StudentID:  parseUint64(studentEntry.Text),
			SubjectID:  parseUint64(subjectEntry.Text),
			Mark:       parseUint16(markEntry.Text),
			Date:       parseDate(dateEntry.Text),
		}

		if err = validation.ValidateStruct(mark); err != nil {
			showResult(content, "Ошибка: "+err.Error())
			return
		}

		res, err := r.Special.IsTeacher(context.Background(), parseUint64(employeeEntry.Text))
		if !res.IsTeacher {
			if err != nil {
				showResult(content, "Ошибка: "+err.Error())
			} else {
				showResult(content, "Ошибка: указанный сотрудник не является преподавателем")
			}
			return
		}

		if err = r.Marks.Update(context.Background(), mark.ID, mark); err != nil {
			showResult(content, "Ошибка: "+err.Error())
			return
		}
		showResult(content, "Оценка обновлена")
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
	headers := []string{
		"ID оценки",
		"ID преподавателя",
		"ID студента",
		"ID предмета",
		"Оценка",
		"Дата",
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
	filterOptions := map[string]uint8{
		"Все":              0,
		"ID":               1,
		"ID преподавателя": 2,
		"ID студента":      3,
		"ID предмета":      4,
		"Оценка":           5,
		"Дата":             6,
	}

	filterEntry := widget.NewEntry()
	filterEntry.SetPlaceHolder("Введите значение")

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

	var data [][]string
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
			showResult(content, "Ошибка: "+err.Error())
			return
		}

		for _, m := range marks {
			data = append(data, []string{
				fmt.Sprintf("%d", m.ID),
				fmt.Sprintf("%d", m.EmployeeID),
				fmt.Sprintf("%d", m.StudentID),
				fmt.Sprintf("%d", m.SubjectID),
				fmt.Sprintf("%d", m.Mark),
				m.Date.Format(dateLayout),
			})
		}

		content.Objects = content.Objects[:1] // Only filter widgets remain
		content.Add(updateTable(headers, data))
		content.Refresh()
	})

	marks, err := r.Marks.FindAll(context.Background())
	if err != nil {
		showResult(content, "Ошибка: "+err.Error())
		return
	}
	for _, m := range marks {
		data = append(data, []string{
			fmt.Sprintf("%d", m.ID),
			fmt.Sprintf("%d", m.EmployeeID),
			fmt.Sprintf("%d", m.StudentID),
			fmt.Sprintf("%d", m.SubjectID),
			fmt.Sprintf("%d", m.Mark),
			m.Date.Format(dateLayout),
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
}
