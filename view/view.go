package view

import (
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

type View struct {
	Header     *widgets.Paragraph
	InfoBar    *widgets.Paragraph
	LeftChart  *widgets.BarChart
	RightChart *widgets.BarChart
	NameList   *widgets.List
	InfoList   *widgets.List
}

var (
	grid *ui.Grid
)

func NewView() *View {

	var view = View{}

	view.Header = widgets.NewParagraph()
	view.Header.Border = true
	view.Header.Min.Y = 1
	view.Header.Text = " Gou Test"

	view.InfoBar = widgets.NewParagraph()
	view.InfoBar.Border = true
	view.InfoBar.Text = " Info TEST"

	view.LeftChart = widgets.NewBarChart()
	view.LeftChart.Title = "Chart 1"
	view.LeftChart.Border = true

	view.RightChart = widgets.NewBarChart()
	view.RightChart.Title = "Chart 2"
	view.RightChart.Border = true

	view.NameList = widgets.NewList()
	view.NameList.Title = "List 1"

	view.InfoList = widgets.NewList()
	view.InfoList.Title = "List 2"

	return &view
}

func (self View) Render() {
	ui.Render(grid)
}

func (self View) SetupGrid() {
	grid = ui.NewGrid()

	grid.Set(
		ui.NewRow(1.0/6, self.Header),
		ui.NewRow(1.0/6, self.InfoBar),
		ui.NewRow(1.0/3,
			ui.NewCol(1.0/2, self.LeftChart),
			ui.NewCol(1.0/2, self.RightChart),
		),
		ui.NewRow(1.0/3,
			ui.NewCol(1.0/2, self.NameList),
			ui.NewCol(1.0/2, self.InfoList),
		),
	)

	termWidth, termHeight := ui.TerminalDimensions()

	self.Resize(termWidth, termHeight)
}

func (self View) Resize(termWidth int, termHeight int) {
	grid.SetRect(0, 0, termWidth, termHeight)
}
