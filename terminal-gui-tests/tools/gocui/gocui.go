package gocui

import (
	"fmt"
	"log"

	"github.com/Eldius/learning-go/terminal-gui-tests/tools/tweet"
	"github.com/jroimartin/gocui"
)

const (
	displayHeight = 3
)

var client tweet.MyTwitterClient

/*
Main tests jroimartin/gocui
*/
func Main() {
	//statuses := tweet.FetchTweets()
	//for _, s := range statuses {
	//	fmt.Println(s.Text())
	//	fmt.Println("---")
	//	//tools.Debug(s)
	//}

	client = tweet.MyTwitterClient{}

	if err := client.Connect(); err != nil {
		log.Fatalf("failed to connect to Twitter: %v", err.Error())
	}

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

func getTopLeft(i int, maxY int) (int, int) {
	return 1, 1 + i*displayHeight
}

func getBottomRight(i int, maxX int) (int, int) {
	return maxX - 1, 3 + i*displayHeight
}

func layout(g *gocui.Gui) error {
	//maxX, maxY := g.Size()
	maxX, maxY := g.Size()

	//fmt.Println(fmt.Sprintf("{x: %d, y: %d}", maxX, maxY))
	statuses := client.FetchTimeline(15)

	for i, s := range statuses {
		fmt.Println(s.FullText)
		x0, y0 := getTopLeft(i, maxY)
		x1, y1 := getBottomRight(i, maxX)
		if y1 >= maxY {
			continue
		}
		if screen, err := g.SetView(fmt.Sprintf("v%d", i), x0, y0, x1, y1); err != nil {
			if err != gocui.ErrUnknownView {
				return err
			}
			screen.Title = fmt.Sprintf("%s:", s.User.Name)
			//fmt.Fprintf(screen, "{\"x\": %d, \"y\": %d}", maxX, maxY)
			//fmt.Fprintf(screen, "%s", s.Text())
			fmt.Fprintf(screen, "{x0: %d, y0: %d, x1: %d, y1: %d, i: %d, maxX: %d, maxY: %d}", x0, y0, x1, y1, i, maxX, maxY)
			//screen.
		}
	}

	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
