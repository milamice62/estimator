package apps

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

type File struct {
	Name string
}

type Entry struct {
	UnitPrice *widget.Entry
	Square    *widget.Entry
	Item      *widget.Entry
}

func NewGUI() fyne.Window {
	currentFile := &File{}
	entry := &Entry{widget.NewEntry(), widget.NewEntry(), widget.NewEntry()}

	//New App and Icon
	app := app.New()
	w := app.NewWindow("Demo")

	//Draw layout
	grid := fyne.NewContainerWithLayout(layout.NewGridLayout(2))
	productGroup := widget.NewGroupWithScroller("ProductsInfo")
	resultGroup := widget.NewGroupWithScroller("Results")
	grid.AddObject(productGroup)
	grid.AddObject(resultGroup)

	productsPanel := []fyne.CanvasObject{
		widget.NewForm(
			widget.NewFormItem("Item: ", entry.Item),
			widget.NewFormItem("Price/sq.ft: ", entry.UnitPrice),
			widget.NewFormItem("Square(mm): ", entry.Square),
		),
		widget.NewButton("Calculate", calculateAndSave(entry, currentFile, resultGroup)),
		widget.NewButton("CreateNewFile", createNewFile(currentFile)),
		widget.NewButton("SelectFile", selectFile(currentFile)),
		widget.NewButton("Quit", func() { app.Quit() }),
	}

	resultPanel := []fyne.CanvasObject{
		fyne.NewContainerWithLayout(
			layout.NewGridLayout(3),
			widget.NewLabelWithStyle("Item", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
			widget.NewLabelWithStyle("Square(ft)", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
			widget.NewLabelWithStyle("Price", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		),
	}

	for _, po := range productsPanel {
		productGroup.Append(po)
	}

	for _, ro := range resultPanel {
		resultGroup.Append(ro)
	}

	w.SetContent(grid)
	return w
}
