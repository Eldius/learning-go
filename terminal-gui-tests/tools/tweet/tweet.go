package tweet

import (
	"fmt"
	"os"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)


/*
MyTwitterClient is an abstraction for the client connection
*/
type MyTwitterClient struct {
	client *twitter.Client
	IsConnected bool
}

/*
Connect will connect your client
*/
func (c *MyTwitterClient)Connect() (err error) {

	consumerKey := os.Getenv("TW_CONSUMER_KEY")
	consumerSecret := os.Getenv("TW_CONSUMER_SECRET")
	accessToken := os.Getenv("TW_ACCESS_TOKEN")
	accessTokenSecret := os.Getenv("TW_ACCESS_TOKEN_SECRET")

	config := oauth1.NewConfig(consumerKey, consumerSecret)
	token := oauth1.NewToken(accessToken, accessTokenSecret)
	// OAuth1 http.Client will automatically authorize Requests
	httpClient := config.Client(oauth1.NoContext, token)

	// Twitter client
	c.client = twitter.NewClient(httpClient)
	
	// Verify Credentials
	verifyParams := &twitter.AccountVerifyParams{
		SkipStatus:   twitter.Bool(true),
		IncludeEmail: twitter.Bool(true),
	}
	user, _, _ := c.client.Accounts.VerifyCredentials(verifyParams)
	fmt.Printf("User's ACCOUNT:\n%+v\n", user)

	c.IsConnected = true
	return
}

/*
FetchTimeline loads fetch tweets from Twitter
*/
func (c *MyTwitterClient)FetchTimeline(maxQtd int) []twitter.Tweet {
	client := c.client
	// Home Timeline
	homeTimelineParams := &twitter.HomeTimelineParams{
		Count:     maxQtd,
		TweetMode: "extended",
	}
	tweets, _, _ := client.Timelines.HomeTimeline(homeTimelineParams)

	return tweets
}
