package main

import (
	"github.com/MilesChou/gou/view"
	ui "github.com/gizak/termui/v3"
	"log"
)

func main() {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	uiView := view.NewView()
	uiView.SetupGrid()
	uiView.Render()

	uiEvents := ui.PollEvents()

	for {
		e := <-uiEvents
		switch e.ID {
		case "q", "<C-c>":
			return
		case "<Resize>":
			payload := e.Payload.(ui.Resize)
			uiView.Resize(payload.Width, payload.Height)
			ui.Clear()
		}

		uiView.Render()
	}
}
