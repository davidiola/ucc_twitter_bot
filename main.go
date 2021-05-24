package main

import (
	"context"
	c "github.com/davidiola/ucc_twitter_bot/constants"
	"github.com/zmb3/spotify"
	"golang.org/x/oauth2/clientcredentials"
	"log"
	"os"
)

func main() {
	config := &clientcredentials.Config{
		ClientID:     os.Getenv(c.SPOTIFY_ID_ENV),
		ClientSecret: os.Getenv(c.SPOTIFY_SEC_ENV),
		TokenURL:     spotify.TokenURL,
	}
	token, err := config.Token(context.Background())
	if err != nil {
		log.Fatalf("Unable to retrieve authorization token: %v", err)
	}

	client := spotify.Authenticator{}.NewClient(token)

	country := "US"
	limit := 50
	opts := spotify.Options{
		Country: &country,
		Limit:   &limit,
	}

	episodeList := make([]spotify.EpisodePage, 0)
	uccEps, err := client.GetShowEpisodesOpt(&opts, c.UCC_SHOW_ID)

	noMorePages := false
	for !noMorePages {
		if err != nil {
			log.Fatalf("Unable to retrieve episode list: %v", err)
		} else {
			episodeList = append(episodeList, uccEps.Episodes...)
		}
		if len(uccEps.Next) == 0 {
			noMorePages = true
		} else {
			client.NextPage(uccEps)
		}
	}

	for idx, ep := range episodeList {
		log.Printf("idx: %d, ep name: %s", idx, ep.Name)
	}
}
