package sheetclient

import (
	"net/http"

	"golang.org/x/net/context"
	"golang.org/x/oauth2"
)

// Retrieve a token, saves the token, then returns the generated client.
func Client(config *oauth2.Config) *http.Client {
	// The file token.json stores the user's access and refresh tokens, and is
	// created automatically when the authorization flow completes for the first
	// time.
	tokFile := "token.json"
	tok, err := tokenFromFile(tokFile)
	if err != nil {
		tok = getTokenFromWeb(config)
		saveToken(tokFile, tok)
	}
	return config.Client(context.Background(), tok)
}
