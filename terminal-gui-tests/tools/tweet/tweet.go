package tweet

import (
	"fmt"
	"github.com/kurrik/oauth1a"
	"github.com/kurrik/twittergo"
	"net/http"
	"net/url"
	"os"
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
func FetchTweets() []twittergo.Tweet {
	var (
		err     error
		client  *twittergo.Client
		req     *http.Request
		resp    *twittergo.APIResponse
		results *twittergo.SearchResults
	)
	client, err = loadCredentials()
	if err != nil {
		fmt.Printf("Could not parse CREDENTIALS file: %v\n", err)
		os.Exit(1)
	}
	query := url.Values{}
	query.Set("q", "twitterapi")
	url := fmt.Sprintf("/1.1/search/tweets.json?%v", query.Encode())
	req, err = http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Printf("Could not parse request: %v\n", err)
		os.Exit(1)
	}
	resp, err = client.SendRequest(req)
	if err != nil {
		fmt.Printf("Could not send request: %v\n", err)
		os.Exit(1)
	}
	results = &twittergo.SearchResults{}
	err = resp.Parse(results)
	if err != nil {
		fmt.Printf("Problem parsing response: %v\n", err)
		os.Exit(1)
	}

	return results.Statuses()
}