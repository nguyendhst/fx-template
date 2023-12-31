package oauth2

import (
	"context"
	"net/http"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

// GoogleProvider represents the Google OAuth2 provider.
type GoogleProvider struct {
	config *oauth2.Config
}

// NewGoogleProvider creates a new GoogleProvider with the specified client ID and secret.
func NewGoogleProvider(clientID, clientSecret, redirectURL string) *GoogleProvider {
	config := &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  redirectURL,
		Scopes:       []string{"profile", "email"},
		Endpoint:     google.Endpoint,
	}

	return &GoogleProvider{config: config}
}

// AuthCodeURL returns the URL to redirect the user to for authorization.
func (p *GoogleProvider) AuthCodeURL(state string) string {
	return p.config.AuthCodeURL(state)
}

// Exchange converts an authorization code into a token.
func (p *GoogleProvider) Exchange(ctx context.Context, code string) (*oauth2.Token, error) {
	return p.config.Exchange(ctx, code)
}

// Client returns an HTTP client that adds the OAuth2 token to requests.
func (p *GoogleProvider) Client(ctx context.Context, token *oauth2.Token) *http.Client {
	return p.config.Client(ctx, token)
}
