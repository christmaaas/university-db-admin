package forms

import (
	"log"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func parseUint64(value string) uint64 {
	num, err := strconv.ParseUint(value, 10, 64)
	if err != nil {
		log.Fatalln(err) // if application cant convert string to int it should be stopped
	}
	return num
}

func showResult(content *fyne.Container, err error, successMessage string) {
	content.Objects = nil
	var labelText string

	if err != nil {
		labelText = "Ошибка: " + err.Error()
	} else {
		labelText = successMessage
	}

	content.Add(widget.NewLabel(labelText))
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
