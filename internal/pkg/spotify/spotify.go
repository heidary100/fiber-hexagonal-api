package spotify

import (
	"context"
	"github.com/zmb3/spotify"
	"golang.org/x/oauth2/clientcredentials"
	"log"
)

func Search(q string) (*spotify.SearchResult, error) {
	config := &clientcredentials.Config{
		ClientID:     "b6a1c5b4ced24b2ab8fd100ece083a22",
		ClientSecret: "3357375942a940ef96c004718958780c",
		TokenURL:     spotify.TokenURL,
	}
	token, err := config.Token(context.Background())
	if err != nil {
		log.Fatalf("couldn't get token: %v", err)
	}

	client := spotify.Authenticator{}.NewClient(token)
	// search for playlists and albums containing "holiday"
	return client.Search(q, spotify.SearchTypeTrack|spotify.SearchTypeAlbum|spotify.SearchTypePlaylist)
}
