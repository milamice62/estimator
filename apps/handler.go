package apps

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	"github.com/milamice62/estimator/calculator"
	"github.com/sqweek/dialog"
)

func calculateAndSave(entry *Entry, file *File, group *widget.Group) func() {
	return func() {
		if file.Name == "" {
			dialog.Message("Please select file!").Title("No File Selected").Info()
			return
		}

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
}

func createNewFile(file *File) func() {
	return func() {
		path, err := dialog.File().Save()
		if err != nil {
			log.Printf("Invalid file path: %v", err)
			return
		}

		_, err = os.Create(path)
		if err != nil {
			log.Printf("Could not create file: %v", err)
			return
		}

		file.Name = path
	}
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

func selectFile(currentFile *File) func() {
	return func() {
		path, err := dialog.File().Load()
		if err != nil {
			dialog.Message("Could not select the file!").Title("File not found").Error()
			return
		}
		currentFile.Name = path
	}
}
