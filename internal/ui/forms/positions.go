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

func ShowPositionsForm(content *fyne.Container, action string, r *repository.Repository) {
	content.Objects = nil

	switch action {
	case "Добавить":
		showAddPositionsForm(content, r)
	case "Удалить":
		showDeletePositionsForm(content, r)
	case "Обновить":
		showUpdatePositionsForm(content, r)
	case "Просмотреть":
		showPositionsList(content, r)
	}

	content.Refresh()
}

func showAddPositionsForm(content *fyne.Container, r *repository.Repository) {
	nameEntry := widget.NewEntry()
	nameEntry.SetPlaceHolder("Название")

	submitButton := widget.NewButton("Добавить", func() {
		pos := domain.Position{
			Name: nameEntry.Text,
		}

		err := r.Positions.Create(context.Background(), pos)
		showResult(content, err, "Должность добавлена")
	})

	form := container.NewVBox(
		widget.NewLabel("Добавление должности"),
		nameEntry,
		submitButton,
	)

	content.Add(form)
}

func showDeletePositionsForm(content *fyne.Container, r *repository.Repository) {
	idEntry := widget.NewEntry()
	idEntry.SetPlaceHolder("ID должности")

	deleteButton := widget.NewButton("Удалить", func() {
		err := r.Positions.Delete(context.Background(), parseUint64(idEntry.Text))
		showResult(content, err, "Должность удалена")
	})

	form := container.NewVBox(
		widget.NewLabel("Удаление должности"),
		idEntry,
		deleteButton,
	)

	content.Add(form)
}

func showUpdatePositionsForm(content *fyne.Container, r *repository.Repository) {
	idEntry := widget.NewEntry()
	idEntry.SetPlaceHolder("ID должности")

	nameEntry := widget.NewEntry()
	nameEntry.SetPlaceHolder("Новое название")

	updateButton := widget.NewButton("Обновить", func() {
		pos := domain.Position{
			ID:   parseUint64(idEntry.Text),
			Name: nameEntry.Text,
		}

		err := r.Positions.Update(context.Background(), pos.ID, pos)
		showResult(content, err, "Должность обновлена")
	})

	form := container.NewVBox(
		widget.NewLabel("Обновление должности"),
		idEntry,
		nameEntry,
		updateButton,
	)

	content.Add(form)
}

func showPositionsList(content *fyne.Container, r *repository.Repository) {
	content.Objects = nil

	headers := []string{
		"ID должности",
		"Название",
	}
	var data [][]string

	filterEntry := widget.NewEntry()
	filterEntry.SetPlaceHolder("Введите значение")

	filterOptions := map[string]uint8{
		"Все":      0,
		"ID":       1,
		"Название": 2,
	}

	options := []string{
		"Все",
		"ID",
		"Название",
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
			positions []domain.Position
			position  domain.Position
			err       error
		)

		switch selectedField {
		case 0:
			positions, err = r.Positions.FindAll(context.Background())
		case 1:
			position, err = r.Positions.FindOne(context.Background(), parseUint64(filterEntry.Text))
			if err == nil {
				positions = append(positions, position)
			}
		case 2:
			position, err = r.Positions.FindByName(context.Background(), filterEntry.Text)
			if err == nil {
				positions = append(positions, position)
			}
		}

		if err != nil {
			showResult(content, err, "Ошибка при поиске")
			return
		}

		for _, p := range positions {
			data = append(data, []string{
				fmt.Sprintf("%d", p.ID),
				p.Name,
			})
		}

		content.Objects = content.Objects[:1] // Only filter widgets remain
		content.Add(updateTable(headers, data))
		content.Refresh()
	})

	positions, _ := r.Positions.FindAll(context.Background())
	for _, p := range positions {
		data = append(data, []string{
			fmt.Sprintf("%d", p.ID),
			p.Name,
		})
	}

	filterContainer := container.NewVBox(
		widget.NewLabel("Фильтрация должностей"),
		filterSelect,
		filterEntry,
		applyFilterButton,
	)

	content.Add(filterContainer)
	content.Add(updateTable(headers, data))
	content.Refresh()
}
