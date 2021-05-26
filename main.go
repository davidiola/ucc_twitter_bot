package main

import (
	c "github.com/davidiola/ucc_twitter_bot/constants"
	sc "github.com/davidiola/ucc_twitter_bot/spotify_client"
	"log"
)

func main() {
	spotify_cl := sc.NewSpotifyCl()
	epList := spotify_cl.RetrieveEpisodesForID(c.UCC_SHOW_ID)

	for idx, ep := range epList {
		log.Printf("idx: %d, ep name: %s", idx, ep.Name)
	}
}
