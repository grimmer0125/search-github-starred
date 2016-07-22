package github

import (
	"github.com/tmtk75/go-oauth2/oauth2"

	xoauth2 "golang.org/x/oauth2"
	xgithub "golang.org/x/oauth2/github"
)

const profileEndpoint = "https://api.github.com/user"

var defaultScopes = []string{}

type githubProvider struct {
	config *xoauth2.Config
}

// New returns oauth2.Provider for github
func New(c *xoauth2.Config) oauth2.Provider {
	c.Endpoint = xgithub.Endpoint
	return githubProvider{config: c}
}

func (g githubProvider) Name() string {
	return oauth2.GITHUB
}

func (g githubProvider) Config() *xoauth2.Config {
	return g.config
}

func (g githubProvider) Profile(token *xoauth2.Token) (oauth2.Profile, error) {
	data, err := oauth2.GetProfileData(profileEndpoint, token.AccessToken)
	if err != nil {
		return nil, err
	}
	return &githubUser{token: token, profile: data}, nil
}

type githubUser struct {
	token   *xoauth2.Token
	profile map[string]interface{}
}

func (u *githubUser) Token() *xoauth2.Token {
	return u.token
}

func (u *githubUser) Name() string {
	return u.profile["name"].(string)
}

func (u *githubUser) Nickname() string {
	return u.profile["login"].(string)
}

func (u *githubUser) AvatarURL() string {
	return u.profile["avatar_url"].(string)
}
