package utils

import (
	"log"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

func ParseUint64(value string) uint64 {
	num, err := strconv.ParseUint(value, 10, 64)
	if err != nil {
		log.Fatalln(err) // if application cant convert string to int it should be stopped
	}
	return num
}

func ShowResult(content *fyne.Container, err error, successMessage string) {
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
