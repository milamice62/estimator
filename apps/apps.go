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

	//New App and Icon
	app := app.New()

	//Windows and widgets
	w := app.NewWindow("Demo")

	//Form fields
	UnitPrice := widget.NewEntry()
	Square := widget.NewEntry()
	Item := widget.NewEntry()
	entry := &Entry{UnitPrice, Square, Item}

	grid := fyne.NewContainerWithLayout(layout.NewGridLayout(2))
	group1 := widget.NewGroupWithScroller("ProductsInfo")
	group2 := widget.NewGroupWithScroller("Results")
	grid.AddObject(group1)
	grid.AddObject(group2)

	productsInfo := []fyne.CanvasObject{
		widget.NewForm(
			widget.NewFormItem("Item: ", Item),
			widget.NewFormItem("Price/sq.ft: ", UnitPrice),
			widget.NewFormItem("Square(mm): ", Square),
		),
		widget.NewButton("Calculate", func() {
			calculateAndSave(entry, currentFile, group2)
		}),
		widget.NewButton("CreateNewFile", createNewFile()),
		widget.NewButton("SelectFile", func() {
			selectFile(currentFile)
			// currentFile.Name = "sample.csv"
		}),
		widget.NewButton("Quit", func() {
			app.Quit()
		}),
	}

	result := []fyne.CanvasObject{
		fyne.NewContainerWithLayout(
			layout.NewGridLayout(3),
			widget.NewLabelWithStyle("Item", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
			widget.NewLabelWithStyle("Square(ft)", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
			widget.NewLabelWithStyle("Price", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		),
	}

	for _, pi := range productsInfo {
		group1.Append(pi)
	}

	for _, r := range result {
		group2.Append(r)
	}

	w.SetContent(grid)
	return w
}
