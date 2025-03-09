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

func ShowGroupsForm(content *fyne.Container, action string, r *repository.Repository) {
	content.Objects = nil

	switch action {
	case "Добавить":
		showAddGroupsForm(content, r)
	case "Удалить":
		showDeleteGroupsForm(content, r)
	case "Обновить":
		showUpdateGroupsForm(content, r)
	case "Просмотреть":
		showGroupsList(content, r)
	}

	content.Refresh()
}

func showAddGroupsForm(content *fyne.Container, r *repository.Repository) {
	numberEntry := widget.NewEntry()
	numberEntry.SetPlaceHolder("Номер")

	submitButton := widget.NewButton("Добавить", func() {
		group := domain.Group{
			Number: parseUint64(numberEntry.Text),
		}

		err := r.Groups.Create(context.Background(), group)
		showResult(content, err, "Группа успешно добавлена")
	})

	form := container.NewVBox(
		widget.NewLabel("Добавление группы"),
		numberEntry,
		submitButton,
	)

	content.Add(form)
}

func showDeleteGroupsForm(content *fyne.Container, r *repository.Repository) {
	idEntry := widget.NewEntry()
	idEntry.SetPlaceHolder("ID группы")

	deleteButton := widget.NewButton("Удалить", func() {
		err := r.Groups.Delete(context.Background(), parseUint64(idEntry.Text))
		showResult(content, err, "Группа удалена")
	})

	form := container.NewVBox(
		widget.NewLabel("Удаление группы"),
		idEntry,
		deleteButton,
	)

	content.Add(form)
}

func showUpdateGroupsForm(content *fyne.Container, r *repository.Repository) {
	idEntry := widget.NewEntry()
	idEntry.SetPlaceHolder("ID группы")

	numberEntry := widget.NewEntry()
	numberEntry.SetPlaceHolder("Новый номер")

	updateButton := widget.NewButton("Обновить", func() {
		group := domain.Group{
			ID:     parseUint64(idEntry.Text),
			Number: parseUint64(numberEntry.Text),
		}

		err := r.Groups.Update(context.Background(), group.ID, group)
		showResult(content, err, "Группа обновлена")
	})

	form := container.NewVBox(
		widget.NewLabel("Обновление группы"),
		idEntry,
		numberEntry,
		updateButton,
	)

	content.Add(form)
}

func showGroupsList(content *fyne.Container, r *repository.Repository) {
	content.Objects = nil

	headers := []string{
		"ID группы",
		"Номер",
	}
	var data [][]string

	filterEntry := widget.NewEntry()
	filterEntry.SetPlaceHolder("Введите значение")

	filterOptions := map[string]uint8{
		"Все":   0,
		"ID":    1,
		"Номер": 2,
	}

	options := []string{
		"Все",
		"ID",
		"Номер",
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
			groups []domain.Group
			grp    domain.Group
			err    error
		)

		switch selectedField {
		case 0:
			groups, err = r.Groups.FindAll(context.Background())
		case 1:
			grp, err = r.Groups.FindOne(context.Background(), parseUint64(filterEntry.Text))
			if err == nil {
				groups = append(groups, grp)
			}
		case 2:
			grp, err = r.Groups.FindByNumber(context.Background(), parseUint64(filterEntry.Text))
			if err == nil {
				groups = append(groups, grp)
			}
		}

		if err != nil {
			showResult(content, err, "Ошибка при поиске")
			return
		}

		for _, g := range groups {
			data = append(data, []string{
				fmt.Sprintf("%d", g.ID),
				fmt.Sprintf("%d", g.Number),
			})
		}

		content.Objects = content.Objects[:1] // Only filter widgets remain
		content.Add(updateTable(headers, data))
		content.Refresh()
	})

	groups, _ := r.Groups.FindAll(context.Background())
	for _, g := range groups {
		data = append(data, []string{
			fmt.Sprintf("%d", g.ID),
			fmt.Sprintf("%d", g.Number),
		})
	}

	filterContainer := container.NewVBox(
		widget.NewLabel("Фильтрация групп"),
		filterSelect,
		filterEntry,
		applyFilterButton,
	)

	content.Add(filterContainer)
	content.Add(updateTable(headers, data))
	content.Refresh()
}
