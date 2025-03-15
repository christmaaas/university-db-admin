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

	contentContainer := container.NewVBox()
	showMainMenu(contentContainer, r)

	w.SetContent(contentContainer)
	w.ShowAndRun()
}

func showMainMenu(content *fyne.Container, r *repository.Repository) {
	content.Objects = nil

	titleLabel := widget.NewLabelWithStyle("Выберите режим работы", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})

	crudButton := widget.NewButton("Операции с сущностями", func() {
		showEntitySelection(content, r)
	})

	queriesButton := widget.NewButton("Специальные SQL-запросы", func() {
		showSpecialQuerySelection(content, r)
	})

	menu := container.NewVBox(
		titleLabel,
		crudButton,
		queriesButton,
	)

	content.Add(menu)
	content.Refresh()
}

func showEntitySelection(content *fyne.Container, r *repository.Repository) {
	content.Objects = nil

	titleLabel := widget.NewLabelWithStyle("Выберите действие", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})
	actionSelect := widget.NewSelect([]string{"Добавить", "Удалить", "Обновить", "Просмотреть"}, nil)

	entityOptions := []string{
		"Сотрудники",
		"Группы",
		"Типы занятий",
		"Занятия",
		"Оценки",
		"Должности",
		"Студенты",
		"Предметы",
		"Знание предметов",
	}
	entitySelect := widget.NewSelect(entityOptions, nil)

	contentContainer := container.NewVBox()
	executeButton := widget.NewButton("Применить", func() {
		updateEntityContent(contentContainer, actionSelect.SelectedIndex(), entitySelect.SelectedIndex(), r)
	})

	backButton := widget.NewButton("Меню", func() {
		showMainMenu(content, r)
	})

	mainContent := container.NewVBox(titleLabel, actionSelect, entitySelect, executeButton, backButton, contentContainer)
	content.Add(mainContent)
	content.Refresh()
}

func showSpecialQuerySelection(content *fyne.Container, r *repository.Repository) {
	content.Objects = nil

	titleLabel := widget.NewLabelWithStyle("Выберите действие", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})

	options := []string{
		"Получить ФИО и номер паспорта всех сотрудников",
		"Получить ФИО и номер паспорта сотрудника по id",
		"Получить информацию о студентах у которых нет куратора",
		"Получить ФИО сотрудников либо с одной должностью, либо другой по id",
		"Получить информацию об оценках выше заданной и выставленных по предмету с заданным id",
		"Получить ФИО и номер паспорта студентов c отчествами, заканчивающимися на <<вна>>",
		"Получить список предметов отсортированных в алфавитном порядке",
		"Получить информацию об оценках отсортированных по дате в порядке возрастания и по значению в порядке убывания",
		"Получить все возможные сочетания студентов и групп",
		"Получить расписание занятий по группам",
		"Получить всех студентов и их кураторов, включая студентов без куратора",
		"Получить всех кураторов и их студентов, включая кураторов без студентов",
		"Получить всех студентов и их кураторов, включая студентов без куратора и кураторов без студентов",
		"Получить ФИО студентов в верхнем регистре и посчитать в них количество символов",
	}
	actionSelect := widget.NewSelect(options, nil)

	contentContainer := container.NewVBox()

	executeButton := widget.NewButton("Выполнить", func() {
		forms.ShowSpecialQueryForm(contentContainer, actionSelect.SelectedIndex(), r)
	})

	backButton := widget.NewButton("Меню", func() {
		showMainMenu(content, r)
	})

	mainContent := container.NewVBox(titleLabel, actionSelect, executeButton, backButton, contentContainer)
	content.Add(mainContent)
	content.Refresh()
}

func updateEntityContent(content *fyne.Container, action, entity int, r *repository.Repository) {
	content.Objects = nil

	switch entity {
	case 0:
		forms.ShowEmployeesForm(content, action, r)
	case 1:
		forms.ShowGroupsForm(content, action, r)
	case 2:
		forms.ShowLessonTypesForm(content, action, r)
	case 3:
		forms.ShowLessonsForm(content, action, r)
	case 4:
		forms.ShowMarksForm(content, action, r)
	case 5:
		forms.ShowPositionsForm(content, action, r)
	case 6:
		forms.ShowStudentsForm(content, action, r)
	case 7:
		forms.ShowSubjectsForm(content, action, r)
	case 8:
		forms.ShowEmployeesSubjectsForm(content, action, r)
	}

	content.Refresh()
}
