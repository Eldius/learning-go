package gocui


import (
	"fmt"
	"log"

	"github.com/jroimartin/gocui"
	"github.com/Eldius/terminal-gui-tests/tools/tweet"
)

/*
Main tests jroimartin/gocui
*/
func Main() {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.SetManagerFunc(layout)

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}

func getTopLeft(i int) (int, int){
	return 1, 1 + i * 2
}

func getBottomRight(i int, maxX int) (int, int){
	return maxX, 3 + i * 2
}

func layout(g *gocui.Gui) error {
	//maxX, maxY := g.Size()
	maxX, _ := g.Size()

	statuses := tweet.FetchTweets()

	// Overlap (front)
	for i, s := range statuses {
		x0, y0 := getTopLeft(i)
		x1, y1 := getBottomRight(i, maxX)
		if v, err := g.SetView("v1", x0, y0, x1, y1); err != nil {
			if err != gocui.ErrUnknownView {
				return err
			}
			v.Title = "Regular title"
			//fmt.Fprintf(v, "{\"x\": %d, \"y\": %d}", maxX, maxY)
			fmt.Fprintf(v, "%d: %s\n", i, s.User)
		}

		//fmt.Fprintf(v, "%d: %s\n", i, s.User)
	}
	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
