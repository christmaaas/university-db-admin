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

func ShowStudentsForm(content *fyne.Container, action int, r *repository.Repository) {
	content.Objects = nil

	switch action {
	case 0:
		showAddStudentsForm(content, r)
	case 1:
		showDeleteStudentsForm(content, r)
	case 2:
		showUpdateStudentsForm(content, r)
	case 3:
		showStudentsList(content, r)
	}

	content.Refresh()
}

func showAddStudentsForm(content *fyne.Container, r *repository.Repository) {
	nameEntry := widget.NewEntry()
	nameEntry.SetPlaceHolder("Имя")

	passportEntry := widget.NewEntry()
	passportEntry.SetPlaceHolder("Паспорт")

	employeeEntry := widget.NewEntry()
	employeeEntry.SetPlaceHolder("ID Куратора")

	groupEntry := widget.NewEntry()
	groupEntry.SetPlaceHolder("ID Группы")

	submitButton := widget.NewButton("Добавить", func() {
		err := validation.ValidateEmptyStrings(
			nameEntry.Text,
			passportEntry.Text,
			employeeEntry.Text,
			groupEntry.Text,
		)
		if err != nil {
			showResult(content, "Ошибка: "+err.Error())
			return
		}

		student := domain.Student{
			Name:       nameEntry.Text,
			Passport:   passportEntry.Text,
			EmployeeID: parseUint64(employeeEntry.Text),
			GroupID:    parseUint64(groupEntry.Text),
		}

		if err = validation.ValidateStruct(student); err != nil {
			showResult(content, "Ошибка: "+err.Error())
			return
		}

		if err = r.Students.Create(context.Background(), student); err != nil {
			showResult(content, "Ошибка: "+err.Error())
			return
		}
		showResult(content, "Студент успешно добавлен")
	})

	form := container.NewVBox(
		widget.NewLabel("Добавление студента"),
		nameEntry,
		passportEntry,
		employeeEntry,
		groupEntry,
		submitButton,
	)

	content.Add(form)
}

func showDeleteStudentsForm(content *fyne.Container, r *repository.Repository) {
	idEntry := widget.NewEntry()
	idEntry.SetPlaceHolder("ID студента")

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

		if err = r.Students.Delete(context.Background(), id); err != nil {
			showResult(content, "Ошибка: "+err.Error())
			return
		}
		showResult(content, "Студент удалён")
	})

	form := container.NewVBox(
		widget.NewLabel("Удаление студента"),
		idEntry,
		deleteButton,
	)

	content.Add(form)
}

func showUpdateStudentsForm(content *fyne.Container, r *repository.Repository) {
	idEntry := widget.NewEntry()
	idEntry.SetPlaceHolder("ID студента")

	nameEntry := widget.NewEntry()
	nameEntry.SetPlaceHolder("Новое имя")

	passportEntry := widget.NewEntry()
	passportEntry.SetPlaceHolder("Новый паспорт")

	employeeEntry := widget.NewEntry()
	employeeEntry.SetPlaceHolder("Новый ID Куратора")

	groupEntry := widget.NewEntry()
	groupEntry.SetPlaceHolder("Новый ID Группы")

	updateButton := widget.NewButton("Обновить", func() {
		err := validation.ValidateEmptyStrings(
			idEntry.Text,
			nameEntry.Text,
			passportEntry.Text,
			employeeEntry.Text,
			groupEntry.Text,
		)
		if err != nil {
			showResult(content, "Ошибка: "+err.Error())
			return
		}

		student := domain.Student{
			ID:         parseUint64(idEntry.Text),
			Name:       nameEntry.Text,
			Passport:   passportEntry.Text,
			EmployeeID: parseUint64(employeeEntry.Text),
			GroupID:    parseUint64(groupEntry.Text),
		}

		if err = validation.ValidateStruct(student); err != nil {
			showResult(content, "Ошибка: "+err.Error())
			return
		}

		if err = r.Students.Update(context.Background(), student.ID, student); err != nil {
			showResult(content, "Ошибка: "+err.Error())
			return
		}
		showResult(content, "Студент обновлён")
	})

	form := container.NewVBox(
		widget.NewLabel("Обновление студента"),
		idEntry,
		nameEntry,
		passportEntry,
		employeeEntry,
		groupEntry,
		updateButton,
	)

	content.Add(form)
}

func showStudentsList(content *fyne.Container, r *repository.Repository) {
	headers := []string{
		"ID студента",
		"Имя",
		"Паспорт",
		"ID Куратора",
		"ID Группы",
	}
	options := []string{
		"Все",
		"ID",
		"Имя",
		"Паспорт",
		"ID Куратора",
		"ID Группы",
	}
	filterOptions := map[string]uint8{
		"Все":         0,
		"ID":          1,
		"Имя":         2,
		"Паспорт":     3,
		"ID Куратора": 4,
		"ID Группы":   5,
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
			students []domain.Student
			stud     domain.Student
			err      error
		)

		switch selectedField {
		case 0:
			students, err = r.Students.FindAll(context.Background())
		case 1:
			stud, err = r.Students.FindOne(context.Background(), parseUint64(filterEntry.Text))
			if err == nil {
				students = append(students, stud)
			}
		case 2:
			students, err = r.Students.FindByName(context.Background(), filterEntry.Text)
		case 3:
			stud, err = r.Students.FindByPassport(context.Background(), filterEntry.Text)
			if err == nil {
				students = append(students, stud)
			}
		case 4:
			students, err = r.Students.FindByEmployeeID(context.Background(), parseUint64(filterEntry.Text))
		case 5:
			students, err = r.Students.FindByGroupID(context.Background(), parseUint64(filterEntry.Text))
		}

		if err != nil {
			showResult(content, "Ошибка: "+err.Error())
			return
		}

		for _, s := range students {
			data = append(data, []string{
				fmt.Sprintf("%d", s.ID),
				s.Name,
				s.Passport,
				fmt.Sprintf("%d", s.EmployeeID),
				fmt.Sprintf("%d", s.GroupID),
			})
		}

		content.Objects = content.Objects[:1] // Only filter widgets remain
		content.Add(updateTable(headers, data))
		content.Refresh()
	})

	students, err := r.Students.FindAll(context.Background())
	if err != nil {
		showResult(content, "Ошибка: "+err.Error())
		return
	}
	for _, s := range students {
		data = append(data, []string{
			fmt.Sprintf("%d", s.ID),
			s.Name,
			s.Passport,
			fmt.Sprintf("%d", s.EmployeeID),
			fmt.Sprintf("%d", s.GroupID),
		})
	}

	filterContainer := container.NewVBox(
		widget.NewLabel("Фильтрация студентов"),
		filterSelect,
		filterEntry,
		applyFilterButton,
	)

	content.Add(filterContainer)
	content.Add(updateTable(headers, data))
}
