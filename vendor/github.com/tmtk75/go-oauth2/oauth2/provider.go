package oauth2

import (
	"os"
	"strings"

	xoauth2 "golang.org/x/oauth2"
)

// Supported providers
const (
	GITHUB   = "github"
	FACEBOOK = "facebook"
	GOOGLE   = "google"
	SLACK    = "slack"
)

// Provider defines behaviors for oauth2 provider
type Provider interface {
	Name() string
	Config() *xoauth2.Config
	Profile(token *xoauth2.Token) (Profile, error)
}

var providers []Provider

// ProviderByName returns a provider given for nnme.
// This returns nil if there is no registered provider.
func ProviderByName(name string) Provider {
	for _, p := range providers {
		if p.Name() == name {
			return p
		}
	}
	return nil
}

// WithProviders registers available providers.
func WithProviders(p ...Provider) {
	providers = p
}

// Providers return registered providers.
func Providers() []Provider {
	return providers
}

// ProfileByCode returns a Profile from the given Provider.
func ProfileByCode(p Provider, code string) (Profile, error) {
	t, err := p.Config().Exchange(xoauth2.NoContext, code)
	if err != nil {
		return nil, err
	}

	u, err := p.Profile(t)
	if err != nil {
		return nil, err
	}

	return u, nil
}

// NewConfig is a helper function to make an instance of golang.org/x/oauth2.Config.
// This must need two environment variables having a given prefix.
// For example, if prefix is `github`, the variables are `GITHUB_CLIENT_ID` and `GITHUB_CLIENT_SECRET`.
// Prefix is normalized with strings.ToUpper. If either value is missing, it makes panic.
// url is a RedirectURL of oauth2.Config
func NewConfig(prefix, url string) *xoauth2.Config {
	m := strings.ToUpper(prefix) + "_CLIENT_ID"
	id, b := os.LookupEnv(m)
	if !b {
		panic(m + " is missing")
	}
	n := strings.ToUpper(prefix) + "_CLIENT_SECRET"
	secret, b := os.LookupEnv(n)
	if !b {
		panic(n + " is missing")
	}
	return &xoauth2.Config{
		ClientID:     id,
		ClientSecret: secret,
		RedirectURL:  url,
	}
}
