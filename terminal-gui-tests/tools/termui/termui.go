package termui

import (
	"fmt"
	"log"
	"os"

	"github.com/Eldius/terminal-gui-tests/tools/tweet"
	"github.com/dghubble/go-twitter/twitter"
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

const (
	mainPanelHeight = 30
	mainPanelWidth  = 130
	displayHeight   = 5
	maxTweets       = 30
)

func getTopLeft(i int) (int, int) {
	return 1, 1 + i*displayHeight
}

func getBottomRight(i int, maxWidh int) (int, int) {
	return maxWidh, 5 + i*displayHeight
}

var currentPosition int = 0
var statuses []twitter.Tweet

/*
Main tests gizak/termui
*/
func Main() {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	termWidth, termHeight := ui.TerminalDimensions()

	p := widgets.NewParagraph()
	p.Title = "Loading..."
	p.Text = "Fetching Tweets..."
	p.SetRect(5, 5, termWidth, termHeight)

	ui.Render(p)
	refreshTweets()
	showTweets()

	for e := range ui.PollEvents() {
		//fmt.Print(e.ID)
		switch e.ID {
		case "q", "<C-c>":
			ui.Close()
			os.Exit(0)
		case "j", "<Down>":
			scrollDown()
		case "k", "<Up>":
			scrollUp()
		case "r", "R":
			refreshTweets()
		}
	}
}

func refreshTweets() {
	statuses = tweet.FetchTweets(maxTweets)
	showTweets()
}

func scrollDown() {
	if currentPosition < len(statuses) {
		currentPosition = currentPosition + 1
		showTweets()
	}
}

func scrollUp() {
	if currentPosition > 0 {
		currentPosition = currentPosition - 1
		showTweets()
	}
}

func showTweets() {
	termWidth, termHeight := ui.TerminalDimensions()
	ui.Clear()
	for i, s := range statuses[currentPosition:] {
		displayTweet(i, s, termHeight, termWidth)
	}
}

func displayTweet(i int, s twitter.Tweet, maxHeight int, maxWidh int) {
	x0, y0 := getTopLeft(i)
	x1, y1 := getBottomRight(i, maxWidh)
	if y1 > maxHeight {
		return
	}
	p := widgets.NewParagraph()
	p.Title = fmt.Sprintf("%s", s.User.Name)
	p.SetRect(x0, y0, x1, y1)
	p.Text = fmt.Sprintf("%s", s.FullText)

	ui.Render(p)

}
