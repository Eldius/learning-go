package termui

import (
	"fmt"
	"log"

	"github.com/Eldius/terminal-gui-tests/tools/tweet"
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"github.com/kurrik/twittergo"
)

const (
	mainPanelHeight = 30
	mainPanelWidth  = 130
	displayHeight   = 5
)

func getTopLeft(i int) (int, int) {
	return 1, 1 + i*displayHeight
}

func getBottomRight(i int, maxWidh int) (int, int) {
	return maxWidh, 5 + i*displayHeight
}

/*
Main tests gizak/termui
*/
func Main() {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	//termWidth, termHeight := ui.TerminalDimensions()
	termWidth, termHeight := ui.TerminalDimensions()

	//grid := ui.NewGrid()
	//grid.SetRect(0, 0, termWidth, termHeight)
	p := widgets.NewParagraph()
	p.Title = "Loading..."
	p.Text = "Fetching Tweets..."
	p.SetRect(5, 5, termWidth, termHeight)

	ui.Render(p)
	statuses := tweet.FetchTweets(10)

	//fmt.Println(len(statuses))
	//var statusToShow []ui.GridItem
	ui.Clear()
	for i, s := range statuses {
		displayTweet(i, s, termHeight, termWidth)
	}

	//grid.Set(statusToShow)

	//ui.Render(grid)

	for e := range ui.PollEvents() {
		if e.Type == ui.KeyboardEvent {
			break
		}
	}
}

func displayTweet(i int, s twittergo.Tweet, maxHeight int, maxWidh int) {
	x0, y0 := getTopLeft(i)
	x1, y1 := getBottomRight(i, maxWidh)
	if y1 > maxHeight {
		return
	}
	p := widgets.NewParagraph()
	p.Title = fmt.Sprintf("user: %s", s.User()["screen_name"])
	p.SetRect(x0, y0, x1, y1)
	p.Text = fmt.Sprintf("status: %s", s.Text())

	ui.Render(p)

}

func setupGrid(grid *ui.Grid, rows ...ui.GridItem) {
	grid.Set(rows)
}
