package tweet

import (
	"fmt"
	"os"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

func newClient() (client *twitter.Client, err error) {

	consumerKey := os.Getenv("TW_CONSUMER_KEY")
	consumerSecret := os.Getenv("TW_CONSUMER_SECRET")
	accessToken := os.Getenv("TW_ACCESS_TOKEN")
	accessTokenSecret := os.Getenv("TW_ACCESS_TOKEN_SECRET")

	config := oauth1.NewConfig(consumerKey, consumerSecret)
	token := oauth1.NewToken(accessToken, accessTokenSecret)
	// OAuth1 http.Client will automatically authorize Requests
	httpClient := config.Client(oauth1.NoContext, token)

	// Twitter client
	client = twitter.NewClient(httpClient)
	return
}

/*
FetchTweets loads fetch tweets from Twitter
*/
func FetchTweets(maxQtd int) []twitter.Tweet {
	client, err := newClient()
	if err != nil {
		fmt.Printf("Could not parse CREDENTIALS file: %v\n", err)
		os.Exit(1)
	}

	// Verify Credentials
	verifyParams := &twitter.AccountVerifyParams{
		SkipStatus:   twitter.Bool(true),
		IncludeEmail: twitter.Bool(true),
	}
	user, _, _ := client.Accounts.VerifyCredentials(verifyParams)
	fmt.Printf("User's ACCOUNT:\n%+v\n", user)

	// Home Timeline
	homeTimelineParams := &twitter.HomeTimelineParams{
		Count:     maxQtd,
		TweetMode: "extended",
	}
	tweets, _, _ := client.Timelines.HomeTimeline(homeTimelineParams)

	return tweets
}
