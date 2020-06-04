package apps

import (
	"fmt"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	"github.com/milamice62/estimator/calculator"
)

var (
	unitPrice *widget.Entry
	quantity  *widget.Entry
)

func NewGUI() fyne.Window {
	//New App and Icon
	app := app.New()

	//Windows and widgets
	w := app.NewWindow("Demo")

	//Form fields
	unitPrice = widget.NewEntry()
	quantity = widget.NewEntry()

	grid := fyne.NewContainerWithLayout(layout.NewGridLayout(2))
	group1 := widget.NewGroup("ProductsInfo")
	group2 := widget.NewGroup("Results")
	grid.AddObject(group1)
	grid.AddObject(group2)

	canvasObjets := []fyne.CanvasObject{
		widget.NewForm(
			widget.NewFormItem("Unit Price ", unitPrice),
			widget.NewFormItem("Quantity: ", quantity),
		),
		widget.NewButton("Calculate", func() {
			price := calculator.CalculatePrice(unitPrice.Text, quantity.Text)
			group2.Append(
				widget.NewForm(
					widget.NewFormItem(
						"TotalPrice",
						widget.NewLabel(fmt.Sprintf("%f", price)),
					),
				),
			)
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
