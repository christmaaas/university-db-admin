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

func ShowSubjectsForm(content *fyne.Container, action int, r *repository.Repository) {
	content.Objects = nil

	switch action {
	case 0:
		showAddSubjectsForm(content, r)
	case 1:
		showDeleteSubjectsForm(content, r)
	case 2:
		showUpdateSubjectsForm(content, r)
	case 3:
		showSubjectsList(content, r)
	}

	content.Refresh()
}

func showAddSubjectsForm(content *fyne.Container, r *repository.Repository) {
	nameEntry := widget.NewEntry()
	nameEntry.SetPlaceHolder("Название")

	dscrEntry := widget.NewEntry()
	dscrEntry.SetPlaceHolder("Описание")

	submitButton := widget.NewButton("Добавить", func() {
		err := validation.ValidateEmptyStrings(nameEntry.Text, dscrEntry.Text)
		if err != nil {
			showResult(content, err, "")
			return
		}

		sbj := domain.Subject{
			Name:        nameEntry.Text,
			Description: dscrEntry.Text,
		}

		if err = validation.ValidateStruct(sbj); err != nil {
			showResult(content, err, "")
			return
		}

		err = r.Subjects.Create(context.Background(), sbj)
		showResult(content, err, "Предмет добавлен")
	})

	form := container.NewVBox(
		widget.NewLabel("Добавление предмета"),
		nameEntry,
		dscrEntry,
		submitButton,
	)

	content.Add(form)
}

func showDeleteSubjectsForm(content *fyne.Container, r *repository.Repository) {
	idEntry := widget.NewEntry()
	idEntry.SetPlaceHolder("ID предмета")

	deleteButton := widget.NewButton("Удалить", func() {
		err := validation.ValidateEmptyStrings(idEntry.Text)
		if err != nil {
			showResult(content, err, "")
			return
		}

		id := parseUint64(idEntry.Text)
		err = validation.ValidatePositiveNumbers(id)
		if err != nil {
			showResult(content, err, "")
			return
		}

		err = r.Subjects.Delete(context.Background(), id)
		showResult(content, err, "Предмет удален")
	})

	form := container.NewVBox(
		widget.NewLabel("Удаление предмета"),
		idEntry,
		deleteButton,
	)

	content.Add(form)
}

func showUpdateSubjectsForm(content *fyne.Container, r *repository.Repository) {
	idEntry := widget.NewEntry()
	idEntry.SetPlaceHolder("ID предмета")

	nameEntry := widget.NewEntry()
	nameEntry.SetPlaceHolder("Новое название")

	dscrEntry := widget.NewEntry()
	dscrEntry.SetPlaceHolder("Новое описание")

	updateButton := widget.NewButton("Обновить", func() {
		err := validation.ValidateEmptyStrings(
			idEntry.Text,
			nameEntry.Text,
			dscrEntry.Text,
		)
		if err != nil {
			showResult(content, err, "")
			return
		}

		sbj := domain.Subject{
			ID:          parseUint64(idEntry.Text),
			Name:        nameEntry.Text,
			Description: dscrEntry.Text,
		}

		if err = validation.ValidateStruct(sbj); err != nil {
			showResult(content, err, "")
			return
		}

		err = r.Subjects.Update(context.Background(), sbj.ID, sbj)
		showResult(content, err, "Предмет обновлен")
	})

	form := container.NewVBox(
		widget.NewLabel("Обновление предмета"),
		idEntry,
		nameEntry,
		dscrEntry,
		updateButton,
	)

	content.Add(form)
}

func showSubjectsList(content *fyne.Container, r *repository.Repository) {
	headers := []string{
		"ID предмета",
		"Название",
		"Описание",
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
			subjects []domain.Subject
			subject  domain.Subject
			err      error
		)

		switch selectedField {
		case 0:
			subjects, err = r.Subjects.FindAll(context.Background())
		case 1:
			subject, err = r.Subjects.FindOne(context.Background(), parseUint64(filterEntry.Text))
			if err == nil {
				subjects = append(subjects, subject)
			}
		case 2:
			subject, err = r.Subjects.FindByName(context.Background(), filterEntry.Text)
			if err == nil {
				subjects = append(subjects, subject)
			}
		}

		if err != nil {
			showResult(content, err, "Ошибка при поиске")
			return
		}

		for _, s := range subjects {
			data = append(data, []string{
				fmt.Sprintf("%d", s.ID),
				s.Name,
				s.Description,
			})
		}

		content.Objects = content.Objects[:1] // Only filter widgets remain
		content.Add(updateTable(headers, data))
		content.Refresh()
	})

	subjects, _ := r.Subjects.FindAll(context.Background())
	for _, s := range subjects {
		data = append(data, []string{
			fmt.Sprintf("%d", s.ID),
			s.Name,
			s.Description,
		})
	}

	filterContainer := container.NewVBox(
		widget.NewLabel("Фильтрация предметов"),
		filterSelect,
		filterEntry,
		applyFilterButton,
	)

	content.Add(filterContainer)
	content.Add(updateTable(headers, data))
}
