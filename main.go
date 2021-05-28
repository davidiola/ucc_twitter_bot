package main

import (
	c "github.com/davidiola/ucc_twitter_bot/constants"
	sc "github.com/davidiola/ucc_twitter_bot/spotify_client"
	tc "github.com/davidiola/ucc_twitter_bot/twitter_client"
	u "github.com/davidiola/ucc_twitter_bot/utils"
	"log"
)

func main() {
	spotify_cl := sc.NewSpotifyCl()
	twitter_cl := tc.NewTwitterCl()

	spotifyEpisodeList := spotify_cl.RetrieveEpisodesForID(c.UCC_SHOW_ID)
	spotifyEpisodeLinks := make([]string, len(spotifyEpisodeList))
	for idx, ep := range spotifyEpisodeList {
		spotifyEpisodeLinks[len(spotifyEpisodeList)-idx-1] = sc.RetrieveLinkFromEpisode(ep)
	}
	uccBotTweets := twitter_cl.RetrieveUserTweets()
	uccBotLinks := make([]string, len(uccBotTweets))

	for idx, t := range uccBotTweets {
		uccBotLinks[idx] = tc.RetrieveExpandedLinkFromTweet(t)
	}

	log.Printf("Retrieved: %d tweets from ucc__bot", len(uccBotLinks))
	for _, epLink := range spotifyEpisodeLinks {
		if !u.Contains(uccBotLinks, epLink) {
			log.Printf("Episode link of: %s not present in tweets: %v. Publishing..", epLink, uccBotLinks)
			twitter_cl.Tweet(epLink)
		}
	}
}
