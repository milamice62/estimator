package apps

import (
	"fmt"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	"github.com/milamice62/estimator/calculator"
	"github.com/milamice62/estimator/handler"
)

type File struct {
	Name string
}

func NewGUI() fyne.Window {
	currentFile := &File{}

	//New App and Icon
	app := app.New()

	//Windows and widgets
	w := app.NewWindow("Demo")

	//Form fields
	unitPrice := widget.NewEntry()
	square := widget.NewEntry()
	name := widget.NewEntry()

	grid := fyne.NewContainerWithLayout(layout.NewGridLayout(2))
	group1 := widget.NewGroup("ProductsInfo")
	group2 := widget.NewGroup("Results")
	grid.AddObject(group1)
	grid.AddObject(group2)

	canvasObjets := []fyne.CanvasObject{
		widget.NewForm(
			widget.NewFormItem("Name: ", name),
			widget.NewFormItem("Price/sq.ft: ", unitPrice),
			widget.NewFormItem("Square(mm): ", square),
		),
		widget.NewButton("Calculate", func() {
			square, price := calculator.CalculatePrice(unitPrice.Text, square.Text)
			handler.SaveToFile(currentFile.Name, name.Text, square, price)
			group2.Append(
				widget.NewForm(
					widget.NewFormItem(
						"Name",
						widget.NewLabel(fmt.Sprintf("%v", name.Text)),
					),
					widget.NewFormItem(
						"SqureFeet",
						widget.NewLabel(fmt.Sprintf("%.2f", square)),
					),
					widget.NewFormItem(
						"TotalPrice",
						widget.NewLabel(fmt.Sprintf("%.2f", price)),
					),
				),
			)
		}),
		widget.NewButton("CreateNewFile", func() {
			handler.CreateNewFile()
		}),
		widget.NewButton("SelectFile", func() {
			currentFile.Name = "sample.csv"
		}),
		widget.NewButton("Quit", func() {
			app.Quit()
		}),
	}

	for _, co := range canvasObjets {
		group1.Append(co)
	}

	w.SetContent(grid)

	return w
}
