package forms

import (
	"context"
	"fmt"
	"university-db-admin/internal/domain"
	"university-db-admin/internal/repository"
	"university-db-admin/internal/ui/utils"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func ShowEmployeeForm(content *fyne.Container, action string, r *repository.Repository) {
	content.Objects = nil

	switch action {
	case "Добавить":
		showAddEmployeeForm(content, r)
	case "Удалить":
		showDeleteEmployeeForm(content, r)
	case "Обновить":
		showUpdateEmployeeForm(content, r)
	case "Просмотреть":
		showEmployeeList(content, r)
	}

	content.Refresh()
}

func showAddEmployeeForm(content *fyne.Container, r *repository.Repository) {
	nameEntry := widget.NewEntry()
	nameEntry.SetPlaceHolder("Имя")

	passportEntry := widget.NewEntry()
	passportEntry.SetPlaceHolder("Паспорт")

	positionEntry := widget.NewEntry()
	positionEntry.SetPlaceHolder("ID Должности")

	submitButton := widget.NewButton("Добавить", func() {
		employee := domain.Employee{
			Name:       nameEntry.Text,
			Passport:   passportEntry.Text,
			PositionID: utils.ParseUint64(positionEntry.Text),
		}

		err := r.Employees.Create(context.Background(), employee)
		utils.ShowResult(content, err, "Сотрудник успешно добавлен")
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

func showDeleteEmployeeForm(content *fyne.Container, r *repository.Repository) {
	idEntry := widget.NewEntry()
	idEntry.SetPlaceHolder("ID сотрудника")

	deleteButton := widget.NewButton("Удалить", func() {
		err := r.Employees.Delete(context.Background(), utils.ParseUint64(idEntry.Text))
		utils.ShowResult(content, err, "Сотрудник удалён")
	})

	form := container.NewVBox(
		widget.NewLabel("Удаление сотрудника"),
		idEntry,
		deleteButton,
	)

	content.Add(form)
}

func showUpdateEmployeeForm(content *fyne.Container, r *repository.Repository) {
	idEntry := widget.NewEntry()
	idEntry.SetPlaceHolder("ID сотрудника")

	nameEntry := widget.NewEntry()
	nameEntry.SetPlaceHolder("Новое имя")

	passportEntry := widget.NewEntry()
	passportEntry.SetPlaceHolder("Новый паспорт")

	positionEntry := widget.NewEntry()
	positionEntry.SetPlaceHolder("Новый ID Должности")

	updateButton := widget.NewButton("Обновить", func() {
		employee := domain.Employee{
			ID:         utils.ParseUint64(idEntry.Text),
			Name:       nameEntry.Text,
			Passport:   passportEntry.Text,
			PositionID: utils.ParseUint64(positionEntry.Text),
		}

		err := r.Employees.Update(context.Background(), employee.ID, employee)
		utils.ShowResult(content, err, "Сотрудник обновлён")
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

func showEmployeeList(content *fyne.Container, r *repository.Repository) {
	content.Objects = nil

	headers := []string{"ID", "Имя", "Паспорт", "ID Должности"}
	var data [][]string

	filterEntry := widget.NewEntry()
	filterEntry.SetPlaceHolder("Введите значение")

	filterOptions := map[string]uint8{
		"Все":          0,
		"ID":           1,
		"Имя":          2,
		"Паспорт":      3,
		"ID Должности": 4,
	}

	var selectedField uint8
	filterSelect := widget.NewSelect([]string{"Все", "ID", "Имя", "Паспорт", "ID Должности"}, func(value string) {
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
			employees []domain.Employee
			emp       domain.Employee
			err       error
		)

		switch selectedField {
		case 0:
			employees, err = r.Employees.FindAll(context.Background())
		case 1:
			emp, err = r.Employees.FindOne(context.Background(), utils.ParseUint64(filterEntry.Text))
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
			employees, err = r.Employees.FindByPosition(context.Background(), utils.ParseUint64(filterEntry.Text))
		}

		if err != nil {
			utils.ShowResult(content, err, "Ошибка при поиске")
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

	employees, _ := r.Employees.FindAll(context.Background())
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
	content.Refresh()
}

func updateTable(headers []string, data [][]string) *fyne.Container {
	table := widget.NewTable(
		func() (int, int) { return len(data) + 1, len(headers) },
		func() fyne.CanvasObject {
			return container.NewHScroll(widget.NewLabel("")) // HScroll wrapper for long strings like names to scroll them
		},
		func(cell widget.TableCellID, obj fyne.CanvasObject) {
			scroll := obj.(*container.Scroll)
			label := scroll.Content.(*widget.Label)

			if cell.Row == 0 {
				label.SetText(headers[cell.Col])
				label.TextStyle.Bold = true
			} else {
				label.SetText(data[cell.Row-1][cell.Col])
			}
		},
	)

	table.SetColumnWidth(0, 50)
	table.SetColumnWidth(1, 800)
	table.SetColumnWidth(2, 100)
	table.SetColumnWidth(3, 120)

	scrollContainer := container.NewVScroll(table)
	scrollContainer.SetMinSize(fyne.NewSize(500, 450))

	return container.NewVBox(scrollContainer)
}
