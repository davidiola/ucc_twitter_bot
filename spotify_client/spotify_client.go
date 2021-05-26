package spotify_client

import (
	"context"
	c "github.com/davidiola/ucc_twitter_bot/constants"
	"github.com/zmb3/spotify"
	"golang.org/x/oauth2/clientcredentials"
	"log"
	"os"
)

type SpotifyCl struct {
	c *spotify.Client
}

func NewSpotifyCl() *SpotifyCl {
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

	return &SpotifyCl{c: &client}
}

func (sc *SpotifyCl) RetrieveEpisodesForID(id string) []spotify.EpisodePage {
	country := "US"
	limit := 50
	opts := spotify.Options{
		Country: &country,
		Limit:   &limit,
	}

	episodeList := make([]spotify.EpisodePage, 0)
	uccEps, err := sc.c.GetShowEpisodesOpt(&opts, id)

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
			sc.c.NextPage(uccEps)
		}
	}

	return episodeList
}
