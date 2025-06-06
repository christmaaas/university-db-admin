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

func ShowEmployeesForm(content *fyne.Container, action int, r *repository.Repository) {
	content.Objects = nil

	switch action {
	case 0:
		showAddEmployeesForm(content, r)
	case 1:
		showDeleteEmployeesForm(content, r)
	case 2:
		showUpdateEmployeesForm(content, r)
	case 3:
		showEmployeesList(content, r)
	}

	content.Refresh()
}

func showAddEmployeesForm(content *fyne.Container, r *repository.Repository) {
	nameEntry := widget.NewEntry()
	nameEntry.SetPlaceHolder("Имя")

	passportEntry := widget.NewEntry()
	passportEntry.SetPlaceHolder("Паспорт")

	positionEntry := widget.NewEntry()
	positionEntry.SetPlaceHolder("ID Должности")

	submitButton := widget.NewButton("Добавить", func() {
		err := validation.ValidateEmptyStrings(
			nameEntry.Text,
			passportEntry.Text,
			positionEntry.Text,
		)
		if err != nil {
			showResult(content, "Ошибка: "+err.Error())
			return
		}

		employee := domain.Employee{
			Name:       nameEntry.Text,
			Passport:   passportEntry.Text,
			PositionID: parseUint64(positionEntry.Text),
		}

		if err = validation.ValidateStruct(employee); err != nil {
			showResult(content, "Ошибка: "+err.Error())
			return
		}

		if err = r.Employees.Create(context.Background(), employee); err != nil {
			showResult(content, "Ошибка: "+err.Error())
			return
		}
		showResult(content, "Сотрудник успешно добавлен")
	})

	form := container.NewVBox(
		widget.NewLabel("Добавление сотрудника"),
		nameEntry,
		passportEntry,
		positionEntry,
		submitButton,
	)

	content.Add(form)
}

func showDeleteEmployeesForm(content *fyne.Container, r *repository.Repository) {
	idEntry := widget.NewEntry()
	idEntry.SetPlaceHolder("ID сотрудника")

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

		if err = r.Employees.Delete(context.Background(), id); err != nil {
			showResult(content, "Ошибка: "+err.Error())
			return
		}
		showResult(content, "Сотрудник удалён")
	})

	form := container.NewVBox(
		widget.NewLabel("Удаление сотрудника"),
		idEntry,
		deleteButton,
	)

	content.Add(form)
}

func showUpdateEmployeesForm(content *fyne.Container, r *repository.Repository) {
	idEntry := widget.NewEntry()
	idEntry.SetPlaceHolder("ID сотрудника")

	nameEntry := widget.NewEntry()
	nameEntry.SetPlaceHolder("Новое имя")

	passportEntry := widget.NewEntry()
	passportEntry.SetPlaceHolder("Новый паспорт")

	positionEntry := widget.NewEntry()
	positionEntry.SetPlaceHolder("Новый ID Должности")

	updateButton := widget.NewButton("Обновить", func() {
		err := validation.ValidateEmptyStrings(
			idEntry.Text,
			nameEntry.Text,
			passportEntry.Text,
			positionEntry.Text,
		)
		if err != nil {
			showResult(content, "Ошибка: "+err.Error())
			return
		}

		employee := domain.Employee{
			ID:         parseUint64(idEntry.Text),
			Name:       nameEntry.Text,
			Passport:   passportEntry.Text,
			PositionID: parseUint64(positionEntry.Text),
		}

		if err = validation.ValidateStruct(employee); err != nil {
			showResult(content, "Ошибка: "+err.Error())
			return
		}

		if err = r.Employees.Update(context.Background(), employee.ID, employee); err != nil {
			showResult(content, "Ошибка: "+err.Error())
			return
		}
		showResult(content, "Сотрудник обновлён")
	})

	form := container.NewVBox(
		widget.NewLabel("Обновление сотрудника"),
		idEntry,
		nameEntry,
		passportEntry,
		positionEntry,
		updateButton,
	)

	content.Add(form)
}

func showEmployeesList(content *fyne.Container, r *repository.Repository) {
	headers := []string{
		"ID сотрудника",
		"Имя",
		"Паспорт",
		"ID Должности",
	}
	options := []string{
		"Все",
		"ID",
		"Имя",
		"Паспорт",
		"ID Должности",
	}
	filterOptions := map[string]uint8{
		"Все":          0,
		"ID":           1,
		"Имя":          2,
		"Паспорт":      3,
		"ID Должности": 4,
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
			employees []domain.Employee
			emp       domain.Employee
			err       error
		)

		switch selectedField {
		case 0:
			employees, err = r.Employees.FindAll(context.Background())
		case 1:
			emp, err = r.Employees.FindOne(context.Background(), parseUint64(filterEntry.Text))
			if err == nil {
				employees = append(employees, emp)
			}
		case 2:
			employees, err = r.Employees.FindByName(context.Background(), filterEntry.Text)
		case 3:
			emp, err = r.Employees.FindByPassport(context.Background(), filterEntry.Text)
			if err == nil {
				employees = append(employees, emp)
			}
		case 4:
			employees, err = r.Employees.FindByPosition(context.Background(), parseUint64(filterEntry.Text))
		}

		if err != nil {
			showResult(content, "Ошибка: "+err.Error())
			return
		}

		for _, e := range employees {
			data = append(data, []string{
				fmt.Sprintf("%d", e.ID),
				e.Name,
				e.Passport,
				fmt.Sprintf("%d", e.PositionID),
			})
		}

		content.Objects = content.Objects[:1] // Only filter widgets remain
		content.Add(updateTable(headers, data))
		content.Refresh()
	})

	employees, err := r.Employees.FindAll(context.Background())
	if err != nil {
		showResult(content, "Ошибка: "+err.Error())
		return
	}
	for _, e := range employees {
		data = append(data, []string{
			fmt.Sprintf("%d", e.ID),
			e.Name,
			e.Passport,
			fmt.Sprintf("%d", e.PositionID),
		})
	}

	filterContainer := container.NewVBox(
		widget.NewLabel("Фильтрация сотрудников"),
		filterSelect,
		filterEntry,
		applyFilterButton,
	)

	content.Add(filterContainer)
	content.Add(updateTable(headers, data))
}
