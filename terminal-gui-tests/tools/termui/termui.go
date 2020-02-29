package termui

import (
	"log"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

/*
Main tests gizak/termui
*/
func Main() {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	p1 := widgets.NewParagraph()
	p1.Text = "P1"
	p1.SetRect(10, 0, 60, 5)

	ui.Render(p1)

	p2 := widgets.NewParagraph()
	p2.Text = "P2"
	p2.SetRect(5, 15, 25, 5)

	ui.Render(p2)

	p3 := widgets.NewParagraph()
	p3.Text = "P3"
	p3.SetRect(62, 0, 100, 5)

	ui.Render(p3)

	p4 := widgets.NewParagraph()
	p4.Text = "P4"
	p4.SetRect(0, 15, 40, 35)

	ui.Render(p4)

	for e := range ui.PollEvents() {
		if e.Type == ui.KeyboardEvent {
			break
		}
	}
}

func createTable() {

}
