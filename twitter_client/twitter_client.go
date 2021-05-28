package twitter_client

import (
	c "github.com/davidiola/ucc_twitter_bot/constants"
	u "github.com/davidiola/ucc_twitter_bot/utils"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"log"
	"math"
)

type TwitterCl struct {
	c  *twitter.Client
	id int64
}

func NewTwitterCl() *TwitterCl {
	apiKey := u.GetEnv(c.TWITTER_API_KEY)
	apiKeySec := u.GetEnv(c.TWITTER_API_KEY_SEC)
	accessToken := u.GetEnv(c.TWITTER_ACCESS_TOKEN)
	accessTokenSec := u.GetEnv(c.TWITTER_ACCESS_TOKEN_SEC)

	config := oauth1.NewConfig(apiKey, apiKeySec)
	token := oauth1.NewToken(accessToken, accessTokenSec)
	httpClient := config.Client(oauth1.NoContext, token)

	client := twitter.NewClient(httpClient)

	verifyParams := &twitter.AccountVerifyParams{
		SkipStatus:   twitter.Bool(true),
		IncludeEmail: twitter.Bool(true),
	}
	user, _, err := client.Accounts.VerifyCredentials(verifyParams)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Printf("User verified successfully. User's description: %s", user.Description)
	}

	return &TwitterCl{
		c:  client,
		id: user.ID,
	}
}

func (tc *TwitterCl) RetrieveUserTweets() []twitter.Tweet {
	maxID := int64(-1)
	var params *twitter.UserTimelineParams
	var tweetList []twitter.Tweet
	var currentTweets []twitter.Tweet
	var err error
	// Use max_id strategy for pagination as described in:
	// https://developer.twitter.com/en/docs/twitter-api/v1/tweets/timelines/guides/working-with-timelines
	for {
		params = &twitter.UserTimelineParams{
			UserID: tc.id,
			Count:  10,
		}
		if maxID != -1 {
			params.MaxID = maxID
		}
		currentTweets, _, err = tc.c.Timelines.UserTimeline(params)
		if err != nil {
			log.Fatalf("Failed to retrieve user's tweets: %s", err)
		} else {
			tweetList = append(tweetList, currentTweets...)
		}

		if len(currentTweets) == 0 {
			break
		}

		currentMaxID := int64(math.MaxInt64)
		for _, t := range currentTweets {
			if t.ID < currentMaxID {
				currentMaxID = t.ID
			}
		}
		maxID = currentMaxID - 1
	}

	return tweetList
}

func (tc *TwitterCl) Tweet(body string) *twitter.Tweet {
	t, _, err := tc.c.Statuses.Update(body, nil)
	if err != nil {
		log.Fatalf("Failed to tweet: %s", err)
	}
	return t
}

func (tc *TwitterCl) Destroy(id int64) *twitter.Tweet {
	t, _, err := tc.c.Statuses.Destroy(id, nil)
	if err != nil {
		log.Fatalf("Failed to destroy tweet: %s", err)
	}
	return t
}

func RetrieveExpandedLinkFromTweet(t twitter.Tweet) string {
	if t.Entities == nil || len(t.Entities.Urls) != 1 {
		log.Printf("Failed to retrieve expanded link from tweet: %v", t)
		return ""
	}
	return t.Entities.Urls[0].ExpandedURL
}
