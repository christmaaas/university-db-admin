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

func ShowEmployeesSubjectsForm(content *fyne.Container, action int, r *repository.Repository) {
	content.Objects = nil

	switch action {
	case 0:
		showAddEmployeesSubjectsForm(content, r)
	case 1:
		showDeleteEmployeesSubjectsForm(content, r)
	case 2:
		showUpdateEmployeesSubjectsForm(content, r)
	case 3:
		showEmployeesSubjectsList(content, r)
	}

	content.Refresh()
}

func showAddEmployeesSubjectsForm(content *fyne.Container, r *repository.Repository) {
	employeeEntry := widget.NewEntry()
	employeeEntry.SetPlaceHolder("ID преподавателя")

	subjectEntry := widget.NewEntry()
	subjectEntry.SetPlaceHolder("ID предмета")

	submitButton := widget.NewButton("Добавить", func() {
		err := validation.ValidateEmptyStrings(employeeEntry.Text, subjectEntry.Text)
		if err != nil {
			showResult(content, "Ошибка: "+err.Error())
			return
		}

		es := domain.EmployeeSubject{
			EmployeeID: parseUint64(employeeEntry.Text),
			SubjectID:  parseUint64(subjectEntry.Text),
		}

		if err = validation.ValidateStruct(es); err != nil {
			showResult(content, "Ошибка: "+err.Error())
			return
		}

		res, err := r.Employees.IsTeacher(context.Background(), parseUint64(employeeEntry.Text))
		if !res.IsTeacher {
			if err != nil {
				showResult(content, "Ошибка: "+err.Error())
			} else {
				showResult(content, "Ошибка: Указанный сотрудник не является преподавателем")
			}
			return
		}

		if err = r.EmployeesSubjects.Create(context.Background(), es); err != nil {
			showResult(content, "Ошибка: "+err.Error())
			return
		}
		showResult(content, "Знание предмета успешно добавлено")
	})

	form := container.NewVBox(
		widget.NewLabel("Добавление знания предмета"),
		employeeEntry,
		subjectEntry,
		submitButton,
	)

	content.Add(form)
}

func showDeleteEmployeesSubjectsForm(content *fyne.Container, r *repository.Repository) {
	employeeEntry := widget.NewEntry()
	employeeEntry.SetPlaceHolder("ID преподавателя")

	subjectEntry := widget.NewEntry()
	subjectEntry.SetPlaceHolder("ID предмета")

	deleteButton := widget.NewButton("Удалить", func() {
		err := validation.ValidateEmptyStrings(employeeEntry.Text, subjectEntry.Text)
		if err != nil {
			showResult(content, "Ошибка: "+err.Error())
			return
		}

		employeeID := parseUint64(employeeEntry.Text)
		subjectID := parseUint64(subjectEntry.Text)
		err = validation.ValidatePositiveNumbers(employeeID, subjectID)
		if err != nil {
			showResult(content, "Ошибка: "+err.Error())
			return
		}

		if err = r.EmployeesSubjects.Delete(context.Background(), employeeID, subjectID); err != nil {
			showResult(content, "Ошибка: "+err.Error())
			return
		}
		showResult(content, "Знание предмета удалено")
	})

	form := container.NewVBox(
		widget.NewLabel("Удаление знания предмета"),
		employeeEntry,
		subjectEntry,
		deleteButton,
	)

	content.Add(form)
}

func showUpdateEmployeesSubjectsForm(content *fyne.Container, r *repository.Repository) {
	employeeEntry := widget.NewEntry()
	employeeEntry.SetPlaceHolder("ID преподавателя")

	subjectEntry := widget.NewEntry()
	subjectEntry.SetPlaceHolder("ID предмета")

	newEmployeeEntry := widget.NewEntry()
	newEmployeeEntry.SetPlaceHolder("Новый ID преподавателя")

	newSubjectEntry := widget.NewEntry()
	newSubjectEntry.SetPlaceHolder("Новый ID предмета")

	updateButton := widget.NewButton("Обновить", func() {
		err := validation.ValidateEmptyStrings(
			employeeEntry.Text,
			subjectEntry.Text,
			newEmployeeEntry.Text,
			newSubjectEntry.Text,
		)
		if err != nil {
			showResult(content, "Ошибка: "+err.Error())
			return
		}

		eid := parseUint64(employeeEntry.Text)
		sid := parseUint64(subjectEntry.Text)
		err = validation.ValidatePositiveNumbers(eid, sid)
		if err != nil {
			showResult(content, "Ошибка: "+err.Error())
			return
		}

		es := domain.EmployeeSubject{
			EmployeeID: parseUint64(newEmployeeEntry.Text),
			SubjectID:  parseUint64(newSubjectEntry.Text),
		}

		if err = validation.ValidateStruct(es); err != nil {
			showResult(content, "Ошибка: "+err.Error())
			return
		}

		res, err := r.Employees.IsTeacher(context.Background(), parseUint64(employeeEntry.Text))
		if !res.IsTeacher {
			if err != nil {
				showResult(content, "Ошибка: "+err.Error())
			} else {
				showResult(content, "Ошибка: указанный сотрудник не является преподавателем")
			}
			return
		}

		if err = r.EmployeesSubjects.Update(context.Background(), eid, sid, es); err != nil {
			showResult(content, "Ошибка: "+err.Error())
			return
		}
		showResult(content, "Знание предмета обновлено")
	})

	form := container.NewVBox(
		widget.NewLabel("Обновление знания предмета"),
		employeeEntry,
		subjectEntry,
		newEmployeeEntry,
		newSubjectEntry,
		updateButton,
	)

	content.Add(form)
}

func showEmployeesSubjectsList(content *fyne.Container, r *repository.Repository) {
	headers := []string{
		"ID преподавателя",
		"ID предмета",
	}
	options := []string{
		"Все",
		"ID преподавателя",
		"ID предмета",
	}
	filterOptions := map[string]uint8{
		"Все":              0,
		"ID преподавателя": 1,
		"ID предмета":      2,
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
			empSbjs []domain.EmployeeSubject
			err     error
		)

		switch selectedField {
		case 0:
			empSbjs, err = r.EmployeesSubjects.FindAll(context.Background())
		case 1:
			empSbjs, err = r.EmployeesSubjects.FindByEmployeeID(context.Background(), parseUint64(filterEntry.Text))
		case 2:
			empSbjs, err = r.EmployeesSubjects.FindBySubjectID(context.Background(), parseUint64(filterEntry.Text))
		}

		if err != nil {
			showResult(content, "Ошибка: "+err.Error())
			return
		}

		for _, e := range empSbjs {
			data = append(data, []string{
				fmt.Sprintf("%d", e.EmployeeID),
				fmt.Sprintf("%d", e.SubjectID),
			})
		}

		content.Objects = content.Objects[:1] // Only filter widgets remain
		content.Add(updateTable(headers, data))
		content.Refresh()
	})

	empSbjs, err := r.EmployeesSubjects.FindAll(context.Background())
	if err != nil {
		showResult(content, "Ошибка: "+err.Error())
		return
	}
	for _, e := range empSbjs {
		data = append(data, []string{
			fmt.Sprintf("%d", e.EmployeeID),
			fmt.Sprintf("%d", e.SubjectID),
		})
	}

	filterContainer := container.NewVBox(
		widget.NewLabel("Фильтрация знания предмета"),
		filterSelect,
		filterEntry,
		applyFilterButton,
	)

	content.Add(filterContainer)
	content.Add(updateTable(headers, data))
}
