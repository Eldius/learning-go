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
	maxTweets       = 100
)

var currentPosition int = 0
var statuses []twitter.Tweet
var termWidth, termHeight int
/*
Main tests gizak/termui
*/
func Main() {

	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	termWidth, termHeight = ui.TerminalDimensions()

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
	if currentPosition < len(statuses) - 1 {
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
	ui.Clear()
	termWidth, termHeight = ui.TerminalDimensions()
	for i, s := range statuses[currentPosition:] {
		displayTweet(i, s)
	}
}

func getMaxX() int {
	return termWidth - 2
}

func getMaxY() int {
	return termHeight - 2
}

func getTopLeft(i int) (int, int) {
	return 1, 1 + i*displayHeight
}

func getBottomRight(i int) (int, int) {
	return getMaxX(), 5 + i*displayHeight
}

func displayTweet(i int, s twitter.Tweet) {
	x0, y0 := getTopLeft(i)
	x1, y1 := getBottomRight(i)
	if y1 > getMaxY() {
		return
	}
	p := widgets.NewParagraph()
	p.Title = fmt.Sprintf("%s", s.User.Name)
	p.SetRect(x0, y0, x1, y1)
	p.Text = fmt.Sprintf("%s", s.FullText)

	ui.Render(p)

}
