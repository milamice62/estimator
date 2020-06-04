package apps

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"time"

	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	"github.com/milamice62/estimator/calculator"
	"github.com/sqweek/dialog"
)

func createNewFile() func() {
	return func() {
		_, err := os.Create("File-" + time.Now().String() + ".csv")
		if err != nil {
			log.Printf("Could not create file: %v", err)
		}
	}
}

func calculateAndSave(entry *Entry, file *File, group *widget.Group) {
	sq, pr := calculator.CalculatePrice(entry.UnitPrice.Text, entry.Square.Text)
	saveToFile(file.Name, entry.Item.Text, sq, pr)
	group.Append(
		fyne.NewContainerWithLayout(
			layout.NewGridLayout(3),
			widget.NewLabelWithStyle(fmt.Sprintf("%v", entry.Item.Text), fyne.TextAlignCenter, fyne.TextStyle{Italic: true}),
			widget.NewLabelWithStyle(fmt.Sprintf("%.2f", sq), fyne.TextAlignCenter, fyne.TextStyle{Italic: true}),
			widget.NewLabelWithStyle(fmt.Sprintf("%.2f", pr), fyne.TextAlignCenter, fyne.TextStyle{Italic: true}),
		),
	)
}

func saveToFile(file string, name string, square float64, price float64) {
	f, err := os.OpenFile(file, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Printf("Could not open file: %v", err)
	}
	defer f.Close()

	record := []string{name, fmt.Sprintf("%.2f", square), fmt.Sprintf("%.2f", price)}
	cw := csv.NewWriter(f)
	cw.Write(record)
	cw.Flush()
}

func selectFile(currentFile *File) {
	fb := dialog.File()
	path, err := fb.Load()
	if err != nil {
		mb := dialog.Message("file")
		mb.Title("Could not open file").Error()
		return
	}
	currentFile.Name = path
}
