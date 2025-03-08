package forms

import (
	"context"
	"fmt"
	"university-db-admin/internal/domain"
	"university-db-admin/internal/repository"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func ShowStudentsForm(content *fyne.Container, action string, r *repository.Repository) {
	content.Objects = nil

	switch action {
	case "Добавить":
		showAddStudentsForm(content, r)
	case "Удалить":
		showDeleteStudentsForm(content, r)
	case "Обновить":
		showUpdateStudentsForm(content, r)
	case "Просмотреть":
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
		student := domain.Student{
			Name:       nameEntry.Text,
			Passport:   passportEntry.Text,
			EmployeeID: parseUint64(employeeEntry.Text),
			GroupID:    parseUint64(groupEntry.Text),
		}

		err := r.Students.Create(context.Background(), student)
		showResult(content, err, "Студент успешно добавлен")
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
		err := r.Students.Delete(context.Background(), parseUint64(idEntry.Text))
		showResult(content, err, "Студент удалён")
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
		student := domain.Student{
			ID:         parseUint64(idEntry.Text),
			Name:       nameEntry.Text,
			Passport:   passportEntry.Text,
			EmployeeID: parseUint64(employeeEntry.Text),
			GroupID:    parseUint64(groupEntry.Text),
		}

		err := r.Students.Update(context.Background(), student.ID, student)
		showResult(content, err, "Студент обновлён")
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
	content.Objects = nil

	headers := []string{
		"ID студента",
		"Имя",
		"Паспорт",
		"ID Куратора",
		"ID Группы",
	}
	var data [][]string

	filterEntry := widget.NewEntry()
	filterEntry.SetPlaceHolder("Введите значение")

	filterOptions := map[string]uint8{
		"Все":         0,
		"ID":          1,
		"Имя":         2,
		"Паспорт":     3,
		"ID Куратора": 4,
		"ID Группы":   5,
	}

	options := []string{
		"Все",
		"ID",
		"Имя",
		"Паспорт",
		"ID Куратора",
		"ID Группы",
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
			showResult(content, err, "Ошибка при поиске")
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

	students, _ := r.Students.FindAll(context.Background())
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
	content.Refresh()
}
