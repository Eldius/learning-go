package tweet

import (
	"fmt"
	"net/http"
	"net/url"
	"os"

	"github.com/kurrik/oauth1a"
	"github.com/kurrik/twittergo"
)

const (
	COUNT       int = 10
	SCREEN_NAME     = "Eldius"
)

func loadCredentials() (client *twittergo.Client, err error) {

	consumerKey := os.Getenv("TW_CONSUMER_KEY")
	consumerSecret := os.Getenv("TW_CONSUMER_SECRET")
	accessToken := os.Getenv("TW_ACCESS_TOKEN")
	accessTokenSecret := os.Getenv("TW_ACCESS_TOKEN_SECRET")
	config := &oauth1a.ClientConfig{
		ConsumerKey:    consumerKey,
		ConsumerSecret: consumerSecret,
	}
	user := oauth1a.NewAuthorizedConfig(accessToken, accessTokenSecret)
	client = twittergo.NewClient(config, user)
	return
}

/*
FetchTweets loads fetch tweets from Twitter
*/
func FetchTweets(maxQtd int) twittergo.Timeline {
	var (
		err    error
		client *twittergo.Client
		req    *http.Request
		resp   *twittergo.APIResponse
		rle    twittergo.RateLimitError
		ok     bool
		query  url.Values
		endpt  string
	)
	client, err = loadCredentials()
	if err != nil {
		fmt.Printf("Could not parse CREDENTIALS file: %v\n", err)
		os.Exit(1)
	}
	query = url.Values{}
	query.Set("count", fmt.Sprintf("%v", maxQtd))
	if client.User == nil {
		// With a user token, the user_timeline.json method
		// returns the current user.  Without, you need to specify
		// an explicit ID.
		//query.Set("screen_name", SCREEN_NAME)
	}
	endpt = fmt.Sprintf("/1.1/statuses/user_timeline.json?%v", query.Encode())
	if req, err = http.NewRequest("GET", endpt, nil); err != nil {
		panic(err.Error())
	}
	if resp, err = client.SendRequest(req); err != nil {
		panic(err.Error())
	}
	t := &twittergo.Timeline{}
	if err = resp.Parse(t); err != nil {
		if rle, ok = err.(twittergo.RateLimitError); ok {
			err = fmt.Errorf("Rate limited. Reset at %v", rle.Reset)
		}
		panic(err.Error())
	}

	return *t
}
