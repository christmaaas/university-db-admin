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

func ShowLessonsForm(content *fyne.Container, action int, r *repository.Repository) {
	content.Objects = nil

	switch action {
	case 0:
		showAddLessonsForm(content, r)
	case 1:
		showDeleteLessonsForm(content, r)
	case 2:
		showUpdateLessonsForm(content, r)
	case 3:
		showLessonsList(content, r)
	}

	content.Refresh()
}

func showAddLessonsForm(content *fyne.Container, r *repository.Repository) {
	groupEntry := widget.NewEntry()
	groupEntry.SetPlaceHolder("ID группы")

	subjectEntry := widget.NewEntry()
	subjectEntry.SetPlaceHolder("ID предмета")

	lTypeEntry := widget.NewEntry()
	lTypeEntry.SetPlaceHolder("ID типа занятия")

	weekEntry := widget.NewEntry()
	weekEntry.SetPlaceHolder("Неделя")

	weekdayEntry := widget.NewEntry()
	weekdayEntry.SetPlaceHolder("День недели")

	roomEntry := widget.NewEntry()
	roomEntry.SetPlaceHolder("Аудитория")

	submitButton := widget.NewButton("Добавить", func() {
		err := validation.ValidateEmptyStrings(
			groupEntry.Text,
			subjectEntry.Text,
			lTypeEntry.Text,
			weekEntry.Text,
			weekdayEntry.Text,
			roomEntry.Text,
		)
		if err != nil {
			showResult(content, err, "")
			return
		}

		lesson := domain.Lesson{
			GroupID:      parseUint64(groupEntry.Text),
			SubjectID:    parseUint64(subjectEntry.Text),
			LessonTypeID: parseUint64(lTypeEntry.Text),
			Week:         parseUint16(weekEntry.Text),
			Weekday:      parseUint16(weekdayEntry.Text),
			Room:         parseUint64(roomEntry.Text),
		}

		if err = validation.ValidateStruct(lesson); err != nil {
			showResult(content, err, "")
			return
		}

		err = r.Lessons.Create(context.Background(), lesson)
		showResult(content, err, "Занятие успешно добавлено")
	})

	form := container.NewVBox(
		widget.NewLabel("Добавление занятия"),
		groupEntry,
		subjectEntry,
		lTypeEntry,
		weekEntry,
		weekdayEntry,
		roomEntry,
		submitButton,
	)

	content.Add(form)
}

func showDeleteLessonsForm(content *fyne.Container, r *repository.Repository) {
	idEntry := widget.NewEntry()
	idEntry.SetPlaceHolder("ID занятия")

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

		err = r.Lessons.Delete(context.Background(), id)
		showResult(content, err, "Занятие удалено")
	})

	form := container.NewVBox(
		widget.NewLabel("Удаление занятия"),
		idEntry,
		deleteButton,
	)

	content.Add(form)
}

func showUpdateLessonsForm(content *fyne.Container, r *repository.Repository) {
	idEntry := widget.NewEntry()
	idEntry.SetPlaceHolder("ID занятия")

	groupEntry := widget.NewEntry()
	groupEntry.SetPlaceHolder("Новый ID группы")

	subjectEntry := widget.NewEntry()
	subjectEntry.SetPlaceHolder("Новый ID предмета")

	lTypeEntry := widget.NewEntry()
	lTypeEntry.SetPlaceHolder("Новый ID типа занятия")

	weekEntry := widget.NewEntry()
	weekEntry.SetPlaceHolder("Новая неделя")

	weekdayEntry := widget.NewEntry()
	weekdayEntry.SetPlaceHolder("Новый день недели")

	roomEntry := widget.NewEntry()
	roomEntry.SetPlaceHolder("Новая аудитория")

	updateButton := widget.NewButton("Обновить", func() {
		err := validation.ValidateEmptyStrings(
			idEntry.Text,
			groupEntry.Text,
			subjectEntry.Text,
			lTypeEntry.Text,
			weekEntry.Text,
			weekdayEntry.Text,
			roomEntry.Text,
		)
		if err != nil {
			showResult(content, err, "")
			return
		}

		lesson := domain.Lesson{
			ID:           parseUint64(idEntry.Text),
			GroupID:      parseUint64(groupEntry.Text),
			SubjectID:    parseUint64(subjectEntry.Text),
			LessonTypeID: parseUint64(lTypeEntry.Text),
			Week:         parseUint16(weekEntry.Text),
			Weekday:      parseUint16(weekdayEntry.Text),
			Room:         parseUint64(roomEntry.Text),
		}

		if err = validation.ValidateStruct(lesson); err != nil {
			showResult(content, err, "")
			return
		}

		err = r.Lessons.Update(context.Background(), lesson.ID, lesson)
		showResult(content, err, "Занятие обновлено")
	})

	form := container.NewVBox(
		widget.NewLabel("Обновление занятия"),
		idEntry,
		groupEntry,
		subjectEntry,
		lTypeEntry,
		weekEntry,
		weekdayEntry,
		roomEntry,
		updateButton,
	)

	content.Add(form)
}

func showLessonsList(content *fyne.Container, r *repository.Repository) {
	headers := []string{
		"ID занятия",
		"ID группы",
		"ID предмета",
		"ID типа занятия",
		"Неделя",
		"День недели",
		"Аудитория",
	}
	options := []string{
		"Все",
		"ID",
		"ID группы",
		"ID предмета",
		"ID типа занятия",
		"Неделя",
		"День недели",
		"Аудитория",
	}
	filterOptions := map[string]uint8{
		"Все":             0,
		"ID":              1,
		"ID группы":       2,
		"ID предмета":     3,
		"ID типа занятия": 4,
		"Неделя":          5,
		"День недели":     6,
		"Аудитория":       7,
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
			lessons []domain.Lesson
			lesson  domain.Lesson
			err     error
		)

		switch selectedField {
		case 0:
			lessons, err = r.Lessons.FindAll(context.Background())
		case 1:
			lesson, err = r.Lessons.FindOne(context.Background(), parseUint64(filterEntry.Text))
			if err == nil {
				lessons = append(lessons, lesson)
			}
		case 2:
			lessons, err = r.Lessons.FindByGroupID(context.Background(), parseUint64(filterEntry.Text))
		case 3:
			lessons, err = r.Lessons.FindBySubjectID(context.Background(), parseUint64(filterEntry.Text))
		case 4:
			lessons, err = r.Lessons.FindByLessonTypeID(context.Background(), parseUint64(filterEntry.Text))
		case 5:
			lessons, err = r.Lessons.FindByWeek(context.Background(), parseUint16(filterEntry.Text))
		case 6:
			lessons, err = r.Lessons.FindByWeekday(context.Background(), parseUint16(filterEntry.Text))
		case 7:
			lessons, err = r.Lessons.FindByRoom(context.Background(), parseUint64(filterEntry.Text))
		}

		if err != nil {
			showResult(content, err, "Ошибка при поиске")
			return
		}

		for _, l := range lessons {
			data = append(data, []string{
				fmt.Sprintf("%d", l.ID),
				fmt.Sprintf("%d", l.GroupID),
				fmt.Sprintf("%d", l.SubjectID),
				fmt.Sprintf("%d", l.LessonTypeID),
				fmt.Sprintf("%d", l.Week),
				fmt.Sprintf("%d", l.Weekday),
				fmt.Sprintf("%d", l.Room),
			})
		}

		content.Objects = content.Objects[:1] // Only filter widgets remain
		content.Add(updateTable(headers, data))
		content.Refresh()
	})

	lessons, err := r.Lessons.FindAll(context.Background())
	if err != nil {
		showResult(content, err, "Ошибка при поиске")
		return
	}
	for _, l := range lessons {
		data = append(data, []string{
			fmt.Sprintf("%d", l.ID),
			fmt.Sprintf("%d", l.GroupID),
			fmt.Sprintf("%d", l.SubjectID),
			fmt.Sprintf("%d", l.LessonTypeID),
			fmt.Sprintf("%d", l.Week),
			fmt.Sprintf("%d", l.Weekday),
			fmt.Sprintf("%d", l.Room),
		})
	}

	filterContainer := container.NewVBox(
		widget.NewLabel("Фильтрация занятий"),
		filterSelect,
		filterEntry,
		applyFilterButton,
	)

	content.Add(filterContainer)
	content.Add(updateTable(headers, data))
}
