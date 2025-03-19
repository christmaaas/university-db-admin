package forms

import (
	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// ui constants for view
const (
	dateLayout = "2006-01-02"
)

// parses uint64 with error handling
func parseUint64(value string) uint64 {
	num, err := strconv.ParseUint(value, 10, 64)
	if err != nil {
		return 0 // if 0 is returned it will be further validated
	}
	return num
}

// parses uint16 with error handling
func parseUint16(value string) uint16 {
	num, err := strconv.ParseUint(value, 10, 16)
	if err != nil {
		return 0 // if 0 is returned it will be further validated
	}
	return uint16(num)
}

// parses time.Time with error handling
func parseDate(dateString string) time.Time {
	parsedTime, err := time.Parse(dateLayout, dateString)
	if err != nil {
		return time.Time{} // if time.Time{} is returned it will be further validated
	}
	return parsedTime
}

// displays a message depending on whether an error occurs
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

// recreates the table with new content
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

	for i := range headers {
		table.SetColumnWidth(i, 300)
	}

	scrollContainer := container.NewVScroll(table)
	scrollContainer.SetMinSize(fyne.NewSize(500, 450))

	return container.NewVBox(scrollContainer)
}
