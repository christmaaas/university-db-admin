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

func ShowLessonTypesForm(content *fyne.Container, action int, r *repository.Repository) {
	content.Objects = nil

	switch action {
	case 0:
		showAddLessonTypesForm(content, r)
	case 1:
		showDeleteLessonTypesForm(content, r)
	case 2:
		showUpdateLessonTypesForm(content, r)
	case 3:
		showLessonTypesList(content, r)
	}

	content.Refresh()
}

func showAddLessonTypesForm(content *fyne.Container, r *repository.Repository) {
	nameEntry := widget.NewEntry()
	nameEntry.SetPlaceHolder("Название")

	submitButton := widget.NewButton("Добавить", func() {
		err := validation.ValidateEmptyStrings(nameEntry.Text)
		if err != nil {
			showResult(content, "Ошибка: "+err.Error())
			return
		}

		lessonType := domain.LessonType{
			Name: nameEntry.Text,
		}

		if err = validation.ValidateStruct(lessonType); err != nil {
			showResult(content, "Ошибка: "+err.Error())
			return
		}

		if err = r.LessonTypes.Create(context.Background(), lessonType); err != nil {
			showResult(content, "Ошибка: "+err.Error())
			return
		}
		showResult(content, "Тип занятия добавлен")
	})

	form := container.NewVBox(
		widget.NewLabel("Добавление типа занятия"),
		nameEntry,
		submitButton,
	)

	content.Add(form)
}

func showDeleteLessonTypesForm(content *fyne.Container, r *repository.Repository) {
	idEntry := widget.NewEntry()
	idEntry.SetPlaceHolder("ID типа занятия")

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

		if err = r.LessonTypes.Delete(context.Background(), id); err != nil {
			showResult(content, "Ошибка: "+err.Error())
			return
		}
		showResult(content, "Тип занятия удален")
	})

	form := container.NewVBox(
		widget.NewLabel("Удаление типа занятия"),
		idEntry,
		deleteButton,
	)

	content.Add(form)
}

func showUpdateLessonTypesForm(content *fyne.Container, r *repository.Repository) {
	idEntry := widget.NewEntry()
	idEntry.SetPlaceHolder("ID типа занятия")

	nameEntry := widget.NewEntry()
	nameEntry.SetPlaceHolder("Новое название")

	updateButton := widget.NewButton("Обновить", func() {
		err := validation.ValidateEmptyStrings(idEntry.Text, nameEntry.Text)
		if err != nil {
			showResult(content, "Ошибка: "+err.Error())
			return
		}

		lType := domain.LessonType{
			ID:   parseUint64(idEntry.Text),
			Name: nameEntry.Text,
		}

		if err = validation.ValidateStruct(lType); err != nil {
			showResult(content, "Ошибка: "+err.Error())
			return
		}

		if err = r.LessonTypes.Update(context.Background(), lType.ID, lType); err != nil {
			showResult(content, "Ошибка: "+err.Error())
			return
		}
		showResult(content, "Тип занятия обновлен")
	})

	form := container.NewVBox(
		widget.NewLabel("Обновление типа занятия"),
		idEntry,
		nameEntry,
		updateButton,
	)

	content.Add(form)
}

func showLessonTypesList(content *fyne.Container, r *repository.Repository) {
	headers := []string{
		"ID типа занятия",
		"Название",
	}
	options := []string{
		"Все",
		"ID",
		"Название",
	}
	filterOptions := map[string]uint8{
		"Все":      0,
		"ID":       1,
		"Название": 2,
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
			lTypes []domain.LessonType
			lType  domain.LessonType
			err    error
		)

		switch selectedField {
		case 0:
			lTypes, err = r.LessonTypes.FindAll(context.Background())
		case 1:
			lType, err = r.LessonTypes.FindOne(context.Background(), parseUint64(filterEntry.Text))
			if err == nil {
				lTypes = append(lTypes, lType)
			}
		case 2:
			lType, err = r.LessonTypes.FindByName(context.Background(), filterEntry.Text)
			if err == nil {
				lTypes = append(lTypes, lType)
			}
		}

		if err != nil {
			showResult(content, "Ошибка: "+err.Error())
			return
		}

		for _, l := range lTypes {
			data = append(data, []string{
				fmt.Sprintf("%d", l.ID),
				l.Name,
			})
		}

		content.Objects = content.Objects[:1] // Only filter widgets remain
		content.Add(updateTable(headers, data))
		content.Refresh()
	})

	lTypes, err := r.LessonTypes.FindAll(context.Background())
	if err != nil {
		showResult(content, "Ошибка: "+err.Error())
		return
	}
	for _, l := range lTypes {
		data = append(data, []string{
			fmt.Sprintf("%d", l.ID),
			l.Name,
		})
	}

	filterContainer := container.NewVBox(
		widget.NewLabel("Фильтрация типов занятий"),
		filterSelect,
		filterEntry,
		applyFilterButton,
	)

	content.Add(filterContainer)
	content.Add(updateTable(headers, data))
}
