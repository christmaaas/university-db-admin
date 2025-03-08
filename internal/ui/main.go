package ui

import (
	"university-db-admin/internal/repository"
	"university-db-admin/internal/ui/forms"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func Run(r *repository.Repository) {
	a := app.New()
	w := a.NewWindow("База данных \"Университет\"")
	w.Resize(fyne.NewSize(1100, 750))

	titleLabel := widget.NewLabelWithStyle("Выберите действие", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})
	actionSelect := widget.NewSelect([]string{"Добавить", "Удалить", "Обновить", "Просмотреть"}, nil)

	options := []string{
		"Сотрудники",
		"Группы",
		"Типы занятий",
		"Занятия",
		"Оценки",
		"Должности",
		"Студенты",
		"Предметы",
	}
	entitySelect := widget.NewSelect(options, nil)

	contentContainer := container.NewVBox()
	executeButton := widget.NewButton("Применить", func() {
		updateContent(contentContainer, actionSelect.Selected, entitySelect.Selected, r)
	})

	mainContent := container.NewVBox(titleLabel, actionSelect, entitySelect, executeButton, contentContainer)
	w.SetContent(mainContent)
	w.ShowAndRun()
}

func updateContent(content *fyne.Container, action, entity string, r *repository.Repository) {
	content.Objects = nil

	switch entity {
	case "Сотрудники":
		forms.ShowEmployeesForm(content, action, r)
	case "Группы":
		forms.ShowGroupsForm(content, action, r)
	case "Типы занятий":
		forms.ShowLessonTypesForm(content, action, r)
	case "Студенты":
		forms.ShowStudentsForm(content, action, r)
	}

	content.Refresh()
}
